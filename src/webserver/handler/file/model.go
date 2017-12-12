package file

import "image"

type ByteSize uint64

const (
	B  = 1
	KB = B << 10
	MB = KB << 10
	GB = MB << 10

	MimeJPEG = "image/jpeg"
	ExtJPEG  = "jpg"
	MimeZIP  = "zip"
	MimeRAR  = ""
)

type uploadImageMapper struct {
	fn        string
	multiple  bool
	mImg      *image.NRGBA
	mPath     string
	mType     string
	tImg      *image.NRGBA
	tPath     string
	tType     string
	tableID   int64
	tableName string
}

type uploadImageParams struct {
	Height    int
	Width     int
	FileName  string
	Extension string
	Mime      string
}

type uploadImageArgs struct {
	Payload   string
	FileName  string
	Extension string
	Mime      string
}

type getProfileParams struct {
	Payload string
}

type getProfileArgs struct {
	Payload string
}

type metaParams struct {
	fileName  string
	mime      string
	extension string
}

type metaArgs struct {
	fileName  string
	mime      string
	extension string
}

type getFileParams struct {
	payload  string
	filename string
}

type getFileArgs struct {
	payload  string
	filename string
}

type routerParams struct {
	payload string
	role    string
	id      string
}

type routerArgs struct {
	payload string
	role    string
	id      int64
}

type uploadFileParams struct {
	id      string
	payload string
	role    string
}

type uploadFileArgs struct {
	id      int64
	payload string
	role    string
}
