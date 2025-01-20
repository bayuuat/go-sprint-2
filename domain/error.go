package domain

import "errors"

var ErrBadRequest = errors.New("bad request")
var ErrActivityNotFound = errors.New("activity not found")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrInvalidActionItem = errors.New("action unknown")
var ErrInvalidUrl = errors.New("invalid url")
var ErrEmailExists = errors.New("email already exists")
var ErrNotFound = errors.New("entity not found")
