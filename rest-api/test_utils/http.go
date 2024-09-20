package test_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	StatusCode int
	Body       string
}

func GetHttpResponse(responseRecorder *httptest.ResponseRecorder) *HttpResponse {

	result := responseRecorder.Result()

	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)

	if err != nil {
		fmt.Sprintf("Unexpected error trying to parse response body: %v\n", err.Error())
		panic("invalid http response")
	}

	return &HttpResponse{
		StatusCode: responseRecorder.Code,
		Body:       string(data),
	}
}

func SetRequestBody(requestObject any, context *gin.Context) {

	bytes, _ := json.Marshal(requestObject)

	context.Request = httptest.NewRequest(http.MethodPost, "http://www.test.com", io.NopCloser(
		strings.NewReader(string(bytes))))

}
