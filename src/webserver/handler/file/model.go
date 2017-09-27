package file

type ByteSize uint64

const (
	B  ByteSize = 1
	KB          = B << 10
	MB          = KB << 10
	GB          = MB << 10
)

type uploadImageParams struct {
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
