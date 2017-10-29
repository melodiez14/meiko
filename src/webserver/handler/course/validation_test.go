package course

import (
	"database/sql"
	"reflect"
	"testing"
)

func Test_createParams_validate(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Description string
		UCU         string
		Semester    string
		Year        string
		StartTime   string
		EndTime     string
		Class       string
		Day         string
		PlaceID     string
		IsUpdate    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    createArgs
		wantErr bool
	}{
		{
			name:    "All empty",
			fields:  fields{},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Overlength ID",
			fields: fields{
				ID: "D10K-7D02-D10K-7D02-D10K-7D02-D10K-7D02-D10K-7D02",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct ID and Empty Name",
			fields: fields{
				ID: "D10K-7D02",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Name Overlength",
			fields: fields{
				ID:   "D10K-7D02",
				Name: "uvuvwevwevwe onyetenyevwe ughemubwem ughemubwem ossasossas",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Name Symbol",
			fields: fields{
				ID:   "D10K-7D02",
				Name: "Risal !@$ Falah",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Name and Empty UCU",
			fields: fields{
				ID:   "D10K-7D02",
				Name: " Risal Falah  ",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Alphabetical UCU",
			fields: fields{
				ID:   "D10K-7D02",
				Name: " Risal Falah  ",
				UCU:  "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Less UCU",
			fields: fields{
				ID:   "D10K-7D02",
				Name: " Risal Falah  ",
				UCU:  "0",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over UCU",
			fields: fields{
				ID:   "D10K-7D02",
				Name: " Risal Falah  ",
				UCU:  "6",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct UCU and Empty Semester",
			fields: fields{
				ID:   "D10K-7D02",
				Name: " Risal Falah  ",
				UCU:  "3",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Alphabetical Semester",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Less Semester",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "0",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over Semester",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "8",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Semester and Empty Year",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "3",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over year",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "3",
				Year:     "9000000000000",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Less year",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "3",
				Year:     "2016",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over year",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "3",
				Year:     "2020",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Year and Empty Start Time",
			fields: fields{
				ID:       "D10K-7D02",
				Name:     " Risal Falah  ",
				UCU:      "3",
				Semester: "3",
				Year:     "2017",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Alphabetical Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "900000000000",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Less Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "-1",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "1441",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Alphabetical End Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over End Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "900000000000",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Less End Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "-1",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Over End Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "1441",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "End Time Less Than Start Time",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "200",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct End Time and Empty Class",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Overlength Class",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Overlength Class",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "ABC",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Numeric Class",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "1",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Class and Empty Day",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Invalid day",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
				Day:       "mondayy",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Day and Empty Place ID",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
				Day:       "monday",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Overlength Place ID",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
				Day:       "monday",
				PlaceID:   "UBJT-0209-UBJT-0209-UBJT-0209-UBJT-0209",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Correct Place ID",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
				Day:       "monday",
				PlaceID:   "UBJT-0209",
			},
			want: createArgs{
				ID:        "D10K-7D02",
				Name:      "Risal Falah",
				UCU:       3,
				Semester:  3,
				Year:      2017,
				StartTime: 600,
				EndTime:   800,
				Class:     "A",
				Day:       1,
				PlaceID:   "UBJT-0209",
			},
			wantErr: false,
		},
		{
			name: "True IsUpdate",
			fields: fields{
				ID:        "D10K-7D02",
				Name:      " Risal Falah  ",
				UCU:       "3",
				Semester:  "3",
				Year:      "2017",
				StartTime: "600",
				EndTime:   "800",
				Class:     "A",
				Day:       "monday",
				PlaceID:   "UBJT-0209",
				IsUpdate:  "true",
			},
			want: createArgs{
				ID:        "D10K-7D02",
				Name:      "Risal Falah",
				UCU:       3,
				Semester:  3,
				Year:      2017,
				StartTime: 600,
				EndTime:   800,
				Class:     "A",
				Day:       1,
				PlaceID:   "UBJT-0209",
				IsUpdate:  true,
			},
			wantErr: false,
		},
		{
			name: "Add Description",
			fields: fields{
				ID:          "D10K-7D02",
				Name:        " Risal Falah  ",
				Description: "This course teach about something",
				UCU:         "3",
				Semester:    "3",
				Year:        "2017",
				StartTime:   "600",
				EndTime:     "800",
				Class:       "A",
				Day:         "monday",
				PlaceID:     "UBJT-0209",
			},
			want: createArgs{
				ID:          "D10K-7D02",
				Name:        "Risal Falah",
				Description: sql.NullString{Valid: true, String: "This course teach about something"},
				UCU:         3,
				Semester:    3,
				Year:        2017,
				StartTime:   600,
				EndTime:     800,
				Class:       "A",
				Day:         1,
				PlaceID:     "UBJT-0209",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := createParams{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				UCU:         tt.fields.UCU,
				Semester:    tt.fields.Semester,
				Year:        tt.fields.Year,
				StartTime:   tt.fields.StartTime,
				EndTime:     tt.fields.EndTime,
				Class:       tt.fields.Class,
				Day:         tt.fields.Day,
				PlaceID:     tt.fields.PlaceID,
				IsUpdate:    tt.fields.IsUpdate,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("createParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
