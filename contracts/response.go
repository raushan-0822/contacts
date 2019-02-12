package contract

import (
	"github.com/raushan-0822/contacts/error"
)

//BaseResponse defines the common structure of standard response
//This contains all but one field of final response
//All the request response will embed this structure
//and should have "response" key
type BaseResponse struct {
	RequestID string          `json:"request_id"`
	Method    string          `json:"method"`
	HTTPCode  int             `json:"http_code"`
	Metadata  *Metadata       `json:"metadata,omitempty"`
	Error     *errorpkg.Error `json:"error"`
}

//Metadata defines the structure for metadata
type Metadata struct {
	PageSize     *int    `json:"page_size,omitempty"`
	Page         *int    `json:"page,omitempty"`
	Total        *int    `json:"total,omitempty"`
	PreviousPage *string `json:"previous_page,omitempty"`
	NextPage     *string `json:"next_page,omitempty"`
}

//SetMetadata setter for Metadata in Response class
func (res *BaseResponse) SetMetadata(metadata *Metadata) *BaseResponse {
	if metadata == nil {
		return res
	}
	res.Metadata = metadata
	return res
}

//SetRequestID sets requst id for response
func (res *BaseResponse) SetRequestID(requestID string) *BaseResponse {
	res.RequestID = requestID
	return res
}

//SetHTTPCode stets the http code
func (res *BaseResponse) SetHTTPCode(code int) *BaseResponse {
	res.HTTPCode = code
	return res
}

//SetMethod  sets http method
func (res *BaseResponse) SetMethod(method string) *BaseResponse {
	res.Method = method
	return res
}

// SetErrorData error data
func (res *BaseResponse) SetErrorData(err *errorpkg.Error) {
	if err != nil {
		res.Error = err
		res.HTTPCode = err.HTTPCode
	}
}
