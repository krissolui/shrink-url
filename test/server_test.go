package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"shrink-url/internal/server"
	"testing"
)

func TestServer(t *testing.T) {
	server := server.Handler()

	tests := []struct {
		name               string
		path               string
		url                string
		expectedHTTPStatus int
		expectedReturn     string
	}{
		{
			name:               "test path not found",
			path:               "/somepath",
			url:                "",
			expectedHTTPStatus: http.StatusNotFound,
			expectedReturn:     "404 page not found\n",
		},
		{
			name:               "test url not found",
			path:               "/",
			url:                "",
			expectedHTTPStatus: http.StatusBadRequest,
			expectedReturn:     "url is required\n",
		},
		{
			name:               "test invalid url format",
			path:               "/",
			url:                "htt/invalid.com",
			expectedHTTPStatus: http.StatusBadRequest,
			expectedReturn:     "url format invalid\n",
		},
		{
			name:               "test success",
			path:               "/",
			url:                "https://www.quantbe.com/welcome/canada/logs/validate",
			expectedHTTPStatus: http.StatusAccepted,
			expectedReturn:     "https://example.co/ysrAXm",
		},
	}

	for _, testcase := range tests {

		t.Run(testcase.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?url=%s", testcase.path, testcase.url), nil)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, testcase.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), testcase.expectedReturn)
		})
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
