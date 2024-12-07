package custom_err

import "errors"

var ErrUnauthorized = errors.New("you are not authorized to access this resource")
var ErrNotFound = errors.New("resource not found")
var ErrUserNameOrEmailAlreadyExist = errors.New("username or email already exist")
var ErrCredentialInvalid = errors.New("email or password is invalid")
