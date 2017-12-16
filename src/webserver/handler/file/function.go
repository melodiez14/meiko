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

	fl "github.com/melodiez14/meiko/src/module/file"
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

func handleUserAssignment(w http.ResponseWriter) error {

	cntDisposition := fmt.Sprintf(`attachment; filename="%d.zip"`, time.Now().Unix())
	w.Header().Set("Pragma", "public")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", cntDisposition)
	w.Header().Set("Content-Transfer-Encoding", "binary")

	zw := zip.NewWriter(w)
	defer zw.Close()

	path := []string{
		"files/var/www/meiko/data/profile/1509451269985766000.058679.1.jpg",
		"files/var/www/meiko/data/profile/1509451499164314000.255057.1.jpg",
	}

	len := len(path)
	if len < 1 {
		return fmt.Errorf("No data")
	} else if len == 1 {
		return handleSingleWithMeta("assignment-upload", "filename.extension", w)
	}

	for _, val := range path {
		file, err := os.Open(val)
		if err != nil {
			return err
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
