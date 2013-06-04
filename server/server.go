package server

type Results interface {
	GetErr() string
	GetOutput() string
}
