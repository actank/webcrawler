package base

import (
	"net/http"
)

type Request struct {
	httpReq *http.Request
	depth uint32
}

func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

func (req *Request) HttpReq *http.Request {
	return req.httpReq
}

func (req *Request) Depth() uint32 {
	return req.depth
}

type Response struct {
	httpResp *http.Response
	depth uint32
}

func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

func (resp *Response) HttpResp() *http.Response {
	return resp.httpResp
}
func (resp *Response) Depth() uint32 {
	return resp.depth
}

type Data interface {
	Valid() bool
}

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item != nil
}


type CrawlerError interface {
	Type() ErrorType
	Error() string
}

type myCrawlerError struct {
	errType ErrorType
	errMsg string
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}