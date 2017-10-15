package user

import (
	"database/sql"
	"html"
	"reflect"
	"testing"

	"github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/helper"
)

func Test_signUpParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
		Name         string
		Email        string
		Password     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    signUpArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: "12345678901",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: "1234567890123",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "123456789012",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: " 12345 6789012 ",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah Asep Nur Muhammad Iskandar Yusuf Rifki Muhammad       ",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah !",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "risal falah",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "ris.",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 12",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
				Password:     "1234",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 13",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
				Password:     "123456",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
				Password:     "<script>alert('mantap')</script>",
			},
			want:    signUpArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 15",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         " Risal Falah ",
				Email:        "  risal@live.com  ",
				Password:     "<script>alert('Mantap123')</script>",
			},
			want: signUpArgs{
				IdentityCode: 123456789012,
				Name:         "Risal Falah",
				Email:        "risal@live.com",
				Password:     helper.StringToMD5(html.EscapeString(("<script>alert('Mantap123')</script>"))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := signUpParams{
				IdentityCode: tt.fields.IdentityCode,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("signUpParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signUpParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emailVerificationParams_validate(t *testing.T) {
	type fields struct {
		Email        string
		IsResendCode string
		Code         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    emailVerificationArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Email: "risal",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				Email: "  ris.",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				Email: "  risal@ live.com   ",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				Email:        "  risal@ live.com   ",
				IsResendCode: "false",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				Email:        "  risal@live.com   ",
				IsResendCode: "true",
			},
			want: emailVerificationArgs{
				Email:        "risal@live.com",
				IsResendCode: true,
				Code:         0,
			},
			wantErr: false,
		},
		{
			name: "Test Case 7",
			fields: fields{
				Email: "  risal@live.com   ",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				Email: "  risal@live.com   ",
				Code:  "1",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				Email: "  risal@live.com   ",
				Code:  "12345",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				Email: "  risal@live.com   ",
				Code:  "halo",
			},
			want:    emailVerificationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				Email: "  risal@live.com   ",
				Code:  "1234",
			},
			want: emailVerificationArgs{
				Email:        "risal@live.com",
				IsResendCode: false,
				Code:         1234,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := emailVerificationParams{
				Email:        tt.fields.Email,
				IsResendCode: tt.fields.IsResendCode,
				Code:         tt.fields.Code,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("emailVerificationParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("emailVerificationParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVerifiedParams_validate(t *testing.T) {
	type fields struct {
		Page  string
		Total string
	}
	tests := []struct {
		name    string
		fields  fields
		want    getVerifiedArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Total: "0",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				Page: "0",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				Page:  "hello",
				Total: "hello",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				Page:  "123",
				Total: "hello",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				Page:  "-1",
				Total: "-1",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				Page:  "1",
				Total: "-1",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				Page:  "-1",
				Total: "1",
			},
			want:    getVerifiedArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				Page:  "100",
				Total: "1000",
			},
			want: getVerifiedArgs{
				Page:  100,
				Total: 1000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := getVerifiedParams{
				Page:  tt.fields.Page,
				Total: tt.fields.Total,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("getVerifiedParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVerifiedParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_activationParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
		Status       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    activationArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    activationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: "1",
			},
			want:    activationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				Status: "1",
			},
			want:    activationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "helo",
				Status:       "1",
			},
			want:    activationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: "140810140016",
				Status:       "1",
			},
			want:    activationArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
				Status:       "active",
			},
			want: activationArgs{
				IdentityCode: 140810140016,
				Status:       user.StatusActivated,
			},
			wantErr: false,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "140810140016",
				Status:       "inactive",
			},
			want: activationArgs{
				IdentityCode: 140810140016,
				Status:       user.StatusVerified,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := activationParams{
				IdentityCode: tt.fields.IdentityCode,
				Status:       tt.fields.Status,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("activationParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("activationParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signInParams_validate(t *testing.T) {
	type fields struct {
		Email    string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		want    signInArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Email: "   risal@ live.com   ",
			},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				Email: "   risal,   ",
			},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				Email: "   risal@live.com   ",
			},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				Email:    "   risal@live.com   ",
				Password: "1234",
			},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				Email:    "   risal@live.com   ",
				Password: "<script>alert('mantap')</script>",
			},
			want:    signInArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				Email:    "   risal@live.com   ",
				Password: "<script>alert('Mantap123')</script>",
			},
			want: signInArgs{
				Email:    "risal@live.com",
				Password: helper.StringToMD5(html.EscapeString(("<script>alert('Mantap123')</script>"))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := signInParams{
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("signInParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signInParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateProfileParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
		Name         string
		Email        string
		Gender       string
		Phone        string
		LineID       string
		Note         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    updateProfileArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: " 123456789 ",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: " 1234567890123456789 ",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "hellohellohello",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: "140810140016",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal.",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@ live.com  ",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@live.com  ",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "Risal !",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
			},
			want: updateProfileArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
			},
			wantErr: false,
		},
		{
			name: "Test Case 12",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
			},
			want: updateProfileArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderMale,
			},
			wantErr: false,
		},
		{
			name: "Test Case 13",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "female",
			},
			want: updateProfileArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderFemale,
			},
			wantErr: false,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 15",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
				Phone:        "085860141146",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 16",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
			},
			want: updateProfileArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderMale,
				Phone:        sql.NullString{Valid: true, String: "85860141146"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 17",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfa",
			},
			want: updateProfileArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderMale,
				Phone:        sql.NullString{Valid: true, String: "85860141146"},
				LineID:       sql.NullString{Valid: true, String: "risalfa"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 18",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Note:         "Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfa",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 19",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "085860141146",
				LineID:       "risalfa",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 20",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfarisalfarisalfarisalfarisalfarisalfarisalfa",
			},
			want:    updateProfileArgs{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := updateProfileParams{
				IdentityCode: tt.fields.IdentityCode,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
				Gender:       tt.fields.Gender,
				Phone:        tt.fields.Phone,
				LineID:       tt.fields.LineID,
				Note:         tt.fields.Note,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("updateProfileParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateProfileParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_changePasswordParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode    string
		Email           string
		OldPassword     string
		Password        string
		ConfirmPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		want    changePasswordArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: " 123456789 ",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: " 1234567890123456789 ",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "hellohellohello",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: "140810140016",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal.",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@ live.com  ",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@live.com  ",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "1234",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 12",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "<script>alert('mantap')</script>",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 13",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "<script>alert('Mantap123')</script>",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "<script>alert('Mantap123')</script>",
				Password:     "1234",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 15",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "<script>alert('Mantap123')</script>",
				Password:     "<script>alert('mantap')</script>",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 15",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				OldPassword:  "<script>alert('Mantap123')</script>",
				Password:     "<script>alert('Mantap123')</script>",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 16",
			fields: fields{
				IdentityCode:    "140810140016",
				Email:           "risal@live.com",
				OldPassword:     "<script>alert('Mantap123')</script>",
				Password:        "<script>alert('Mantap123')</script>",
				ConfirmPassword: "<script>alert('mantap')</script>",
			},
			want:    changePasswordArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 17",
			fields: fields{
				IdentityCode:    "140810140016",
				Email:           "risal@live.com",
				OldPassword:     "<script>alert('Mantap123')</script>",
				Password:        "<script>alert('Mantap123')</script>",
				ConfirmPassword: "<script>alert('Mantap123')</script>",
			},
			want: changePasswordArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				OldPassword:  helper.StringToMD5(html.EscapeString(("<script>alert('Mantap123')</script>"))),
				Password:     helper.StringToMD5(html.EscapeString(("<script>alert('Mantap123')</script>"))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := changePasswordParams{
				IdentityCode:    tt.fields.IdentityCode,
				Email:           tt.fields.Email,
				OldPassword:     tt.fields.OldPassword,
				Password:        tt.fields.Password,
				ConfirmPassword: tt.fields.ConfirmPassword,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("changePasswordParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changePasswordParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forgotParams_validate(t *testing.T) {
	type fields struct {
		Email      string
		IsSendCode string
		Password   string
		Code       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    forgotArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Email: "risal",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				Email: "ris.",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				Email: "risal@live.com",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				Email:      "risal@live.com",
				IsSendCode: "haha",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				Email:      "risal@live.com",
				IsSendCode: "true",
			},
			want: forgotArgs{
				Email:      "risal@live.com",
				IsSendCode: true,
			},
			wantErr: false,
		},
		{
			name: "Test Case 7",
			fields: fields{
				Email:      "risal@live.com",
				IsSendCode: "true",
			},
			want: forgotArgs{
				Email:      "risal@live.com",
				IsSendCode: true,
			},
			wantErr: false,
		},
		{
			name: "Test Case 8",
			fields: fields{
				Email: "risal@live.com",
				Code:  "1",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				Email: "risal@live.com",
				Code:  "haha",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				Email: "risal@live.com",
				Code:  "1234",
			},
			want: forgotArgs{
				Email:      "risal@live.com",
				IsSendCode: false,
				Code:       1234,
				Password:   "",
			},
			wantErr: false,
		},
		{
			name: "Test Case 11",
			fields: fields{
				Email: "risal@live.com",
				Code:  "1234",
			},
			want: forgotArgs{
				Email:      "risal@live.com",
				IsSendCode: false,
				Code:       1234,
			},
			wantErr: false,
		},
		{
			name: "Test Case 12",
			fields: fields{
				Email:    "risal@live.com",
				Code:     "1234",
				Password: "12345",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 13",
			fields: fields{
				Email:    "risal@live.com",
				Code:     "1234",
				Password: "<script>alert('mantap')</script>",
			},
			want:    forgotArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				Email:    "risal@live.com",
				Code:     "1234",
				Password: "<script>alert('Mantap123')</script>",
			},
			want: forgotArgs{
				Email:      "risal@live.com",
				IsSendCode: false,
				Code:       1234,
				Password:   helper.StringToMD5(html.EscapeString(("<script>alert('Mantap123')</script>"))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := forgotParams{
				Email:      tt.fields.Email,
				IsSendCode: tt.fields.IsSendCode,
				Password:   tt.fields.Password,
				Code:       tt.fields.Code,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("forgotParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("forgotParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detailParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
	}
	tests := []struct {
		name    string
		fields  fields
		want    detailArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    detailArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: "Hello Moto",
			},
			want:    detailArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: "123456789",
			},
			want:    detailArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "1234567890123456789",
			},
			want:    detailArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: " 140810140016 ",
			},
			want:    detailArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
			},
			want: detailArgs{
				IdentityCode: 140810140016,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := detailParams{
				IdentityCode: tt.fields.IdentityCode,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("detailParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detailParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
		Name         string
		Email        string
		Gender       string
		Phone        string
		LineID       string
		Note         string
		Status       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    updateArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: " 123456789 ",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: " 1234567890123456789 ",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "hellohellohello",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: "140810140016",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "  risal.",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@ live.com  ",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "   risal@live.com  ",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "Risal !",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 12",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 13",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "female",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 14",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 15",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "haha",
				Phone:        "085860141146",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 16",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 17",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfa",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 18",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Note:         "Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello Hello ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfa",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 19",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "085860141146",
				LineID:       "risalfa",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 20",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Phone:        "85860141146",
				LineID:       "risalfarisalfarisalfarisalfarisalfarisalfarisalfa",
			},
			want:    updateArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 21",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Status:       "active",
			},
			want: updateArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Status:       user.StatusActivated,
			},
			wantErr: false,
		},
		{
			name: "Test Case 22",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "male",
				Status:       "active",
			},
			want: updateArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderMale,
				Status:       user.StatusActivated,
			},
			wantErr: false,
		},
		{
			name: "Test Case 23",
			fields: fields{
				IdentityCode: "140810140016",
				Email:        "risal@live.com",
				Name:         "  Risal  Falah     ",
				Gender:       "female",
				Status:       "inactive",
			},
			want: updateArgs{
				IdentityCode: 140810140016,
				Email:        "risal@live.com",
				Name:         "Risal Falah",
				Gender:       user.GenderFemale,
				Status:       user.StatusVerified,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := updateParams{
				IdentityCode: tt.fields.IdentityCode,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
				Gender:       tt.fields.Gender,
				Phone:        tt.fields.Phone,
				LineID:       tt.fields.LineID,
				Note:         tt.fields.Note,
				Status:       tt.fields.Status,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("updateParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
	}
	tests := []struct {
		name    string
		fields  fields
		want    deleteArgs
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			fields:  fields{},
			want:    deleteArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				IdentityCode: "Hello Moto",
			},
			want:    deleteArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: "123456789",
			},
			want:    deleteArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "1234567890123456789",
			},
			want:    deleteArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: " 140810140016 ",
			},
			want:    deleteArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "140810140016",
			},
			want: deleteArgs{
				IdentityCode: 140810140016,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := deleteParams{
				IdentityCode: tt.fields.IdentityCode,
			}
			got, err := params.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("deleteParams.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteParams.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createParams_validate(t *testing.T) {
	type fields struct {
		IdentityCode string
		Name         string
		Email        string
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
			name: "Test Case 2",
			fields: fields{
				IdentityCode: "123456789",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				IdentityCode: "1234567890123456789",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			fields: fields{
				IdentityCode: "123456789012",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			fields: fields{
				IdentityCode: " 12345 6789012 ",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 6",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah Asep Nur Muhammad Iskandar Yusuf Rifki Muhammad       ",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 7",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah !",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 8",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 9",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "risal falah",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 10",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "ris.",
			},
			want:    createArgs{},
			wantErr: true,
		},
		{
			name: "Test Case 11",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
			},
			want: createArgs{
				IdentityCode: 123456789012,
				Name:         "Risal Falah",
				Email:        "risal@live.com",
			},
			wantErr: false,
		},
		{
			name: "Test Case 12",
			fields: fields{
				IdentityCode: "123456789012",
				Name:         "Risal Falah",
				Email:        "  risal@live.com  ",
			},
			want: createArgs{
				IdentityCode: 123456789012,
				Name:         "Risal Falah",
				Email:        "risal@live.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := createParams{
				IdentityCode: tt.fields.IdentityCode,
				Name:         tt.fields.Name,
				Email:        tt.fields.Email,
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
