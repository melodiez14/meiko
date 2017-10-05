package file

import "image"

type ByteSize uint64

const (
	B  ByteSize = 1
	KB          = B << 10
	MB          = KB << 10
	GB          = MB << 10

	MimeJPEG = "image/jpeg"
	ExtJPEG  = "jpg"
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
	Payload   string
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
