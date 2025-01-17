package domain

import "errors"

var ErrBadRequest = errors.New("bad request")
var ErrUserNotFound = errors.New("user not found")
var ErrIdentityNumberNotFound = errors.New("identity number not found")
var ErrDepartmentNotFound = errors.New("department not found")
var ErrEmployeeNotFound = errors.New("employee not found")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrInvalidActionItem = errors.New("action unknown")
var ErrInvalidUrl = errors.New("invalid url")
var ErrEmailExists = errors.New("email already exists")
var ErrNotFound = errors.New("entity not found")
var ErrEmployeeExists = errors.New("employee already exists")
var ErrDepartmentHasEmployees = errors.New("department still contain employee")
