package gofarmhttp

import (
	"fmt"
	"io"
	"strconv"
)

type Response struct {
	Status     string // 200 OK
	StatusCode int    // 200
	Proto      string // HTTP/1.0
	ProtoMajor int    // 1
	ProtoMinor int    // 0

	// TODO(Jeremy): Implement headers
	// Header Header

	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	// TODO(Jeremy): Implement Trailers
	// Trailer Header
	//TODO(Jeremy): Implement Request
	// Request *Request
	// TLS *tls.ConnectionState

}

func (r *Response) Write(w io.Writer) error {
	requestMethod := "GET"
	// status_text := r.Status

	// Writes Http version and status code into the buffer
	if _, err := fmt.Fprintf(w, "HTTP/%d.%d %03d %s\r\n", r.ProtoMajor, r.ProtoMinor, r.StatusCode, statusText[r.StatusCode]); err != nil {
		return err
	}

	// tw, err := newTransferWriter(r1)

	// Write content length header into buffer
	if shouldSendContentLength(r.ContentLength, &requestMethod) {
		io.WriteString(w, "Content-Length: ")
		io.WriteString(w, strconv.FormatInt(r.ContentLength, 10)+"\r\n")
	}

	// Write end of header into buffer
	io.WriteString(w, "\r\n")

	// copy body into io buffer
	io.Copy(w, r.Body)

	// Write trailer into io buffer
	// io.WriteString(w, "\r\n")
	return nil
}

func shouldSendContentLength(contentLength int64, requestMethod *string) bool {
	if contentLength > 0 {
		return true
	}
	if contentLength < 0 {
		return false
	}

	return false
}
