package errs

import "errors"

var (
	ErrRecordNotFound = errors.New("Record not found")
	ReturnErrsForApi = errors.New("Invalidate ID")
	ErrInternalServer = errors.New("Internal server error, try again later!")
	ErrCreateUserFailed = errors.New("Create user failed")
	ErrLastNameValidate = errors.New("Error validation LastName")
	ErrFirstNameValidate = errors.New("Error validation FirstName")
	ErrEmailValidate = errors.New("Error validation Email")
	ErrStatusNotFound = errors.New("Error validation Status")
)