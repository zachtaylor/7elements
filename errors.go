package vii

import "errors"

// ErrNotImplemented is a self explanatory
var ErrNotImplemented = errors.New("not implemented")

// ErrKarma means you don't have enough karma
var ErrKarma = errors.New("requires more karma")
