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

type getVerifiedUserResponse struct {
	ID    int64
	Name  string
	Email string
}
type setStatusUserParams struct {
	Email string
	Code  string
}

type setStatusUserArgs struct {
	Email string
	Code  uint16
}

type activationParams struct {
	ID     string
	Status string
}

type activationArgs struct {
	ID     int64
	Status int8
}
