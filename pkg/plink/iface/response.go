package iface

import "io"

type Response struct {
	StatusCode    int
	Body          io.ReadCloser
	ContentLength int64
	Request       *Request
}

func (r Response) Write([]byte) (int, error) {
	return 1, nil
}
func (r Response) WriteHeader(statusCode int) {
	println("123")
}

type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

func Error(w ResponseWriter, r *Request, Status int) {

}
