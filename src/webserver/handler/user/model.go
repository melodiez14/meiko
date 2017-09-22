package user

type signUpParams struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type signUpArgs struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type emailVerificationParams struct {
	Email        string
	IsResendCode string
	Code         string
}

type emailVerificationArgs struct {
	Email        string
	IsResendCode bool
	Code         uint16
}

type signInParams struct {
	Email    string
	Password string
}

type signInArgs struct {
	Email    string
	Password string
}

type forgotRequestParams struct {
	Email string
}

type forgotRequestArgs struct {
	Email string
}

type forgotRequestResponse struct {
	Email          string `json:"email"`
	ExpireDuration string `json:"expire_duration"`
	MaxAttempt     uint8  `json:"max_attempt"`
}

type forgotConfirmationParams struct {
	Email    string
	Password string
	Code     string
}

type forgotConfirmationArgs struct {
	Email    string
	Password string
	Code     uint16
}

type getVerifiedParams struct {
	Page  string
	Total string
}

type getVerifiedArgs struct {
	Page  int64
	Total int64
}

type getVerifiedResponse struct {
	ID     int64  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type setUserAccoutParams struct {
	Name    string
	Gender  string
	Phone   string
	LineID  string
	College string
	Note    string
}

type setUserAccoutArgs struct {
	Name    string
	Gender  int8
	Phone   string
	LineID  string
	College string
	Note    string
}
type setChangePasswordParams struct {
	Password        string
	ConfirmPassword string
}
type setChangePasswordArgs struct {
	Password string
}

type activationParams struct {
	ID     string
	Status string
}

type activationArgs struct {
	ID     int64
	Status int8
}
