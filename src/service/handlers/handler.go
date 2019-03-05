package handler

import (
	"net/http"
	"strconv"

	contract "github.com/raushan-0822/contacts/contracts"
	"github.com/raushan-0822/contacts/error"
	"github.com/raushan-0822/contacts/model"

	"github.com/labstack/echo"
)

// ContactHandler returns response
func ContactHandler(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:
		return ContactCreateHandler(c)
	case http.MethodDelete:
		return ContactDeleteHandler(c)
	case http.MethodGet:
		return ContactGetHandler(c)
	case http.MethodPut:
		return ContactUpdateHandler(c)
	}
	return RawResponse(c, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

}

// ContactCreateHandler handle post request for contact api
func ContactCreateHandler(c echo.Context) error {
	var response contract.ContactsResponse
	requestID := c.Get("RequestID").(string)
	method := c.Request().Method
	req := new(contract.ContactPostRequest)
	if err := c.Bind(req); err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidBody(err))
	}
	if err := c.Validate(req); err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
	}
	var contact model.Contact
	contact.AccountID = c.Param("accountSid")
	contact.Number = req.Number
	contact.FirstName = req.FirstName
	contact.LastName = req.LastName
	contact.CompanyName = req.CompanyName
	contact.Email = req.Email
	ResContact, err := model.AddContact(&contact)
	if err != nil {
		return Response(c, &response, errorpkg.ErrInternalServerError(err))
	}
	response.RequestID = requestID
	response.Method = method
	response.HTTPCode = 200
	newContact := ResContact
	response.ResourceData = newContact
	return Response(c, &response, nil)
}

// ContactUpdateHandler handles put request
func ContactUpdateHandler(c echo.Context) error {
	var response contract.ContactsResponse
	requestID := c.Get("RequestID").(string)
	method := c.Request().Method
	req := new(contract.ContactUpdateRequest)
	if err := c.Bind(req); err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidBody(err))
	}
	if err := c.Validate(req); err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
	}
	var contact model.Contact
	contact.AccountID = c.Param("accountSid")
	if u64, err := strconv.ParseUint(c.Param("id"), 10, 64); err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
	} else {
		contact.ID = uint(u64)
	}
	//we can use phonenumber validation here
	if req.Number != "" {
		contact.Number = req.Number
	}
	if req.FirstName != "" {
		contact.FirstName = req.FirstName
	}
	if req.LastName != "" {
		contact.LastName = req.LastName
	}
	if req.CompanyName != "" {
		contact.CompanyName = req.CompanyName
	}
	if req.Email != "" {
		contact.Email = req.Email
	}
	ResContact, err := model.UpdateContact(&contact)
	if err != nil {
		return Response(c, &response, errorpkg.ErrInternalServerError(err))
	}
	response.RequestID = requestID
	response.Method = method
	response.HTTPCode = 200
	newContact := ResContact
	response.ResourceData = newContact
	return Response(c, &response, nil)

}

// ContactGetHandler handles get request
func ContactGetHandler(c echo.Context) error {
	var response contract.ContactsResponse
	requestID := c.Get("RequestID").(string)
	method := c.Request().Method
	accountID := c.Param("accountSid")
	u64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
	}
	ResContact, err := model.GetContactByID(uint(u64))
	if err != nil {
		return Response(c, &response, errorpkg.ErrInternalServerError(err))
	}
	if ResContact.AccountID != accountID {
		return Response(c, &response, errorpkg.ErrForbidden(err))
	}
	response.RequestID = requestID
	response.Method = method
	response.HTTPCode = 200
	newContact := ResContact
	response.ResourceData = newContact
	return Response(c, &response, nil)

}

// ContactSearchHandler handles search request
func ContactSearchHandler(c echo.Context) error {
	var response contract.ContactsSearchResponse
	requestID := c.Get("RequestID").(string)
	method := c.Request().Method
	accountID := c.Param("accountSid")
	name := c.Param("name")
	page := c.Param("page")
	var pageU64 uint64
	var err error
	if page != "" {
		pageU64, err = strconv.ParseUint(page, 10, 64)
		if err != nil {
			return BulkResponse(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
		}
	}
	ResContact, err := model.GetContactByName(name, accountID, pageU64)
	if err != nil {
		return BulkResponse(c, &response, errorpkg.ErrInternalServerError(err))
	}
	response.RequestID = requestID
	response.Method = method
	response.HTTPCode = 200
	response.ResourceData = ResContact
	var metaData contract.Metadata
	length := len(*ResContact)
	metaData.PageSize = length
	metaData.Page = int(pageU64)
	response.SetMetadata(&metaData)
	return BulkResponse(c, &response, nil)

}

// ContactDeleteHandler remove a given contact
func ContactDeleteHandler(c echo.Context) error {
	var response contract.ContactsResponse
	requestID := c.Get("RequestID").(string)
	method := c.Request().Method
	accountID := c.Param("accountSid")
	u64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return Response(c, &response, errorpkg.ErrBadRequestInvalidParameter(err))
	}
	ResContact, err := model.GetContactByID(uint(u64))
	if err != nil {
		return Response(c, &response, errorpkg.ErrInternalServerError(err))
	}
	if ResContact.AccountID != accountID {
		return Response(c, &response, errorpkg.ErrForbidden(err))
	}
	err = model.DeleteContact(uint(u64))
	if err != nil {
		return Response(c, &response, errorpkg.ErrInternalServerError(err))
	}
	response.RequestID = requestID
	response.Method = method
	response.HTTPCode = 200
	return Response(c, &response, nil)
}

//RawResponse creates response and responds it
//it encode the respone in the format that requester `accpets`
//checks the accept header for the same
func RawResponse(c echo.Context, response interface{}, httpCode int) error {
	var responseFunc func(int, interface{}) error
	switch c.Request().Header.Get("accept") {
	case "application/json", "text/json", "json":
		responseFunc = c.JSON
	case "text/xml", "application/xml", "xml":
		responseFunc = c.XML
	default:
		responseFunc = c.JSON
	}
	return responseFunc(httpCode, response)
}

// Response returns Response
func Response(c echo.Context, response *contract.ContactsResponse, err *errorpkg.Error) error {
	response.SetErrorData(err)
	return RawResponse(c, response, response.BaseResponse.HTTPCode)
}

// BulkResponse returns return pagination Response
func BulkResponse(c echo.Context, response *contract.ContactsSearchResponse, err *errorpkg.Error) error {
	response.SetErrorData(err)
	return RawResponse(c, response, response.BaseResponse.HTTPCode)
}
