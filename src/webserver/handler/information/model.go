package information

const (
	TableNameInformation = "informations"
)

type getParams struct {
	page  string
	total string
}

type getArgs struct {
	page  int64
	total int64
}

type getResponse struct {
	ID             int64  `json:"id"`
	Title          string `json:"title"`
	Date           string `json:"date"`
	Description    string `json:"description"`
	Image          string `json:"image"`
	ImageThumbnail string `json:"image_thumbnail"`
}

type detailInfromationParams struct {
	ID string
}
type detailInfromationArgs struct {
	ID int64
}

type createParams struct {
	Title       string
	Description string
	ScheduleID  string
	FilesID     string
}
type createArgs struct {
	Title       string
	Description string
	ScheduleID  int64
	FilesID     []string
}
type updateParams struct {
	ID          string
	Title       string
	Description string
	ScheduleID  string
	FilesID     string
}

type upadateArgs struct {
	ID          int64
	Title       string
	Description string
	ScheduleID  int64
	FilesID     []string
}

type deleteParams struct {
	ID string
}

type deleteArgs struct {
	ID int64
}

type readListParams struct {
	total string
	page  string
}

type readListArgs struct {
	total int64
	page  int64
}
