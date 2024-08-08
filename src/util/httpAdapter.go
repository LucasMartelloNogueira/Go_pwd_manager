package util

import (
	"net/http"
	"fmt"
	"encoding/json"
)

const (
	HttpInternalServerErrorStatusMsg = "Internal Server Error"
	HttpNotFoundStatusMsg = "Not Found"
	HttpUnauthorizedStatusMsg = "Unauthorized"

	HttpStatusSuccess = "Success"
	HttpStatusError = "Error"
)

var errorResponseHandler map[string]http.HandlerFunc = map[string]http.HandlerFunc {
	HttpInternalServerErrorStatusMsg: httpInternalServerErrorResponse,
	HttpNotFoundStatusMsg: httpNotFoundResponse,
	HttpUnauthorizedStatusMsg: httpUnauthorizedResponse,
}

func GetErrorResponseHandler(errorString string) func(http.ResponseWriter, *http.Request){
	errHandler, ok := errorResponseHandler[errorString]
	if ok {
		return errHandler
	}
	return httpInternalServerErrorResponse
}

func GetHttpResponse(w http.ResponseWriter, r *http.Request, data any, err error) {
	if err != nil {
		handler := GetErrorResponseHandler(err.Error())
		handler(w, r)
		return
	}
	httpSuccessResponse(w, r, data)
}


func httpDefaultResponse(w http.ResponseWriter, r *http.Request, payload any) {
	dataInBytes, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Fprint(w, string(dataInBytes))
}

func httpSuccessResponse(w http.ResponseWriter, r *http.Request, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	payload := map[string]any {
		"data": data,
		"status": HttpStatusSuccess,
		"statusCode": http.StatusOK,
	}
	httpDefaultResponse(w, r, payload)
}


func httpDefaultErrorResponse(w http.ResponseWriter, r *http.Request, httpStatusCode int, httpStatusMessage string) {
	payload := map[string]any {
		"error": map[string]any{
			"code": httpStatusCode,
			"mesage": httpStatusMessage,
		},
		"status": "error",
		"statusCode": httpStatusCode,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	httpDefaultResponse(w, r, payload)
}


func httpInternalServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	httpDefaultErrorResponse(w, r, http.StatusInternalServerError, HttpInternalServerErrorStatusMsg)
}

func httpNotFoundResponse(w http.ResponseWriter, r *http.Request) {
	httpDefaultErrorResponse(w, r, http.StatusNotFound, HttpNotFoundStatusMsg)
}

func httpUnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	httpDefaultErrorResponse(w, r, http.StatusUnauthorized, HttpUnauthorizedStatusMsg)
}