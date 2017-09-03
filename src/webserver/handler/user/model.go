package user

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
	Email          string
	ExpireDuration string
	MaxAttempt     uint8
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
