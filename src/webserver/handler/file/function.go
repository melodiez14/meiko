package file

import (
	"archive/zip"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	asg "github.com/melodiez14/meiko/src/module/assignment"
	cs "github.com/melodiez14/meiko/src/module/course"
	fl "github.com/melodiez14/meiko/src/module/file"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/alias"
	"github.com/melodiez14/meiko/src/util/helper"
)

func handleSingleWithMeta(payload, filename string, w http.ResponseWriter) error {

	var file *os.File

	// load file from disk
	path := fmt.Sprintf("%s/%s/%s", alias.Dir["data"], payload, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fn, _, err := helper.ExtractExtension(filename)
	if err != nil {
		return err
	}

	// load metadata
	fileInfo, err := fl.GetByIDExt(fn, fl.ColName, fl.ColExtension, fl.ColMime)
	if err != nil {
		return err
	}

	cntDisposition := fmt.Sprintf(`attachment; filename="%s.%s"`, fileInfo.Name, fileInfo.Extension)
	w.Header().Set("Content-Type", fileInfo.Mime)
	w.Header().Set("Content-Disposition", cntDisposition)
	w.Header().Set("Cache-Control", "private, max-age=2628000")

	stat, err := file.Stat()
	if err == nil {
		size := strconv.FormatInt(stat.Size(), 10)
		w.Header().Set("Content-Length", size)
	}

	// seek the read before copying to the response
	file.Seek(0, 0)
	io.Copy(w, file)

	return nil
}

func handleSingleWithoutMeta(payload, filename string, w http.ResponseWriter) error {

	// load file from disk
	path := fmt.Sprintf("%s/%s/%s", alias.Dir["data"], payload, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "private, max-age=2628000")

	stat, err := file.Stat()
	if err == nil {
		size := strconv.FormatInt(stat.Size(), 10)
		w.Header().Set("Content-Length", size)
	}

	// seek the read before copying to the response
	file.Seek(0, 0)
	io.Copy(w, file)

	return nil
}

func handleUserAssignment(userID, assignmentID int64, w http.ResponseWriter) error {

	assignment, err := asg.GetByID(assignmentID)
	if err != nil {
		return err
	}

	scheduleID, err := cs.GetScheduleIDByGP(assignment.GradeParameterID)
	if err != nil {
		return err
	}

	if !cs.IsAssistant(userID, scheduleID) {
		return fmt.Errorf("You are not authorized")
	}

	tableID := strconv.FormatInt(assignment.ID, 10)
	files, err := fl.SelectByRelation(fl.TypAssignmentUpload, []string{tableID}, nil)
	if err != nil {
		return err
	}

	studentsID := []int64{}
	for _, val := range files {
		studentsID = append(studentsID, val.UserID)
	}

	users, err := usr.SelectByID(studentsID, false, usr.ColID, usr.ColIdentityCode)
	if err != nil {
		return err
	}

	studentMap := map[int64]int64{}
	for _, val := range users {
		studentMap[val.ID] = val.IdentityCode
	}

	cntDisposition := fmt.Sprintf(`attachment; filename="%s_%s.zip"`, time.Now().Format("20060102150405"), assignment.Name)
	w.Header().Set("Pragma", "public")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", cntDisposition)
	w.Header().Set("Content-Transfer-Encoding", "binary")

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, val := range files {
		path := fmt.Sprintf("%s/assignment/%s.%s", alias.Dir["data"], val.ID, val.Extension)
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return err
		}

		hdr, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		hdr.Method = zip.Deflate
		hdr.Name = fmt.Sprintf("%d_%s.%s", studentMap[val.UserID], val.Name, val.Extension)
		writter, err := zw.CreateHeader(hdr)
		if err != nil {
			return err
		}

		_, err = io.Copy(writter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleUpload(file multipart.File, header *multipart.FileHeader, userID int64, typ string, payload string) (fileResponse, int, error) {

	var resp fileResponse

	// extract file extension
	defer file.Close()
	file.Seek(0, 0)

	fn, ext, err := helper.ExtractExtension(header.Filename)
	if err != nil {
		return resp, http.StatusBadRequest, err
	}

	params := metaParams{
		fileName:  fn,
		extension: ext,
		mime:      header.Header.Get("Content-Type"),
	}

	args, err := params.validate()
	if err != nil {
		return resp, http.StatusBadRequest, fmt.Errorf("Invalid Request")
	}

	// get filename
	t := time.Now().UnixNano()
	rand.Seed(t)
	fileID := fmt.Sprintf("%d.%06d", t, rand.Intn(999999))

	// save file
	go func() {
		path := fmt.Sprintf("%s/%s/%s.%s", alias.Dir["data"], payload, fileID, args.extension)
		f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()

		file.Seek(0, 0)
		io.Copy(f, file)
	}()

	err = fl.Insert(fileID, args.fileName, args.mime, args.extension, userID, typ, nil)
	if err != nil {
		return resp, http.StatusInternalServerError, fmt.Errorf("Internal Error")
	}

	return fileResponse{
		ID:           fileID,
		Name:         header.Filename,
		URLThumbnail: helper.MimeToThumbnail(params.mime),
	}, http.StatusOK, nil
}
