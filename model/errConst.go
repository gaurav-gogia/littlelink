package model

import "errors"

// ErrMethodNotAllowed defines an error where user calls an API with wrong method
// Severity: Low
var ErrMethodNotAllowed = errors.New("Method NOT allowed")

// ErrQueryMethodNotAllowed defines an error where user calls an API by sending params via wrong format
// Severity: Low
var ErrQueryMethodNotAllowed = errors.New("Query Params came through wrong source")

// ErrLinkFail defines an error where server is unable to verify if the long url is working or not
// Severity: Medium
var ErrLinkFail = errors.New("Could NOT verify long URL")

// ErrSomethingWentWrong tells user that something went wrong with server but deosn't tell what went wrong
var ErrSomethingWentWrong = errors.New("Something went wrong")
