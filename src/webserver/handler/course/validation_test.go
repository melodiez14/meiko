package course

import (
	"reflect"
	"testing"
)

func Test_createParams_validate(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		UCU         string
		Semester    string
		StartTime   string
		EndTime     string
		Class       string
		Day         string
		PlaceID     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    createArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name:    "Test Case 2",
			fields:  fields{},
			want:    createArgs{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := createParams{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				UCU:         tt.fields.UCU,
				Semester:    tt.fields.Semester,
				StartTime:   tt.fields.StartTime,
				EndTime:     tt.fields.EndTime,
				Class:       tt.fields.Class,
				Day:         tt.fields.Day,
				PlaceID:     tt.fields.PlaceID,
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
