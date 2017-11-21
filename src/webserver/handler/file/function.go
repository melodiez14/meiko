package file

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	fl "github.com/melodiez14/meiko/src/module/file"
	"github.com/melodiez14/meiko/src/util/helper"
)

func handleSingleWithMeta(payload, filename string, w http.ResponseWriter) error {

	var file *os.File

	// load file from disk
	path := fmt.Sprintf("files/var/www/meiko/data/%s/%s", payload, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fn, ext, err := helper.ExtractExtension(filename)
	if err != nil {
		return err
	}

	// load metadata
	fileInfo, err := fl.GetByIDExt(fn, ext, fl.ColName, fl.ColExtension, fl.ColMime)
	if err != nil {
		return err
	}

	cntDisposition := fmt.Sprintf(`attachment; filename="%s.%s"`, fileInfo.Name, fileInfo.Extension)
	w.Header().Set("Content-Type", fileInfo.Mime)
	w.Header().Set("Content-Disposition", cntDisposition)
	w.Header().Set("Cache-Control", "no-transform, max-age=2628000")

	io.Copy(w, file)

	return nil
}

func handleJPEGWithoutMeta(payload, filename string, w http.ResponseWriter) error {

	// load file from disk
	path := fmt.Sprintf("files/var/www/meiko/data/%s/%s", payload, filename)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-Control", "no-transform, max-age=2628000")

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
