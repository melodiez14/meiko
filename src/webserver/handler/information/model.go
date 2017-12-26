package information

import (
	"time"
)

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
type respListInformation struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	CreatedDate string `json:"created_at"`
	Description string `json:"description"`
	UpdatedDate string `json:"updated_at"`
	CourseName  string `json:"course_name"`
}
type respDetailInformation struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedDate string    `json:"created_at"`
	UpdatedDate string    `json:"updated_at"`
	Date        time.Time `json:"date"`
	ScheduleID  int64     `json:"schedule_id"`
	CourseName  string    `json:"course_name"`
}

type getDetailParams struct {
	id string
}

type getDetailArgs struct {
	id int64
}

type getDetailResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
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
