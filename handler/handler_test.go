package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAverageHandler(t *testing.T) {
	tests := []struct {
		body                 string
		query                string
		expectedStatus       int
		expectedBody         string
		expectedErrorMessage string
	}{
		{
			// Valid case
			body:                 "10 20 30 40",
			query:                "n=3",
			expectedStatus:       http.StatusOK,
			expectedBody:         "The average is 25.00\n",
			expectedErrorMessage: "",
		},
		{
			// Empty body
			body:                 "",
			query:                "n=3",
			expectedStatus:       http.StatusBadRequest,
			expectedBody:         "no numbers provided\n",
			expectedErrorMessage: "",
		},
		{
			// Invalid input
			body:                 "10 20 x 40" + "\n",
			query:                "n=3",
			expectedStatus:       http.StatusBadRequest,
			expectedBody:         "Invalid input: please provide space-separated digits only\n",
			expectedErrorMessage: "",
		},
		{
			// Even case for n=1
			body:                 "1 2 3 4 5",
			query:                "n=1",
			expectedStatus:       http.StatusOK,
			expectedBody:         "The average is 3.00\n",
			expectedErrorMessage: "",
		},

		{
			// Odd case for n=2
			body:                 "1 2 3 4 5",
			query:                "n=2",
			expectedStatus:       http.StatusOK,
			expectedBody:         "The average is 3.00\n",
			expectedErrorMessage: "",
		},
		{
			// Larger numbers
			body:                 "100 200 300",
			query:                "n=3",
			expectedStatus:       http.StatusOK,
			expectedBody:         "The average is 200.00\n",
			expectedErrorMessage: "",
		},
		{
			// Invalid n parameter
			body:                 "10 20 30",
			query:                "n=4",
			expectedStatus:       http.StatusBadRequest,
			expectedBody:         "Invalid choice\n",
			expectedErrorMessage: "",
		},
		{
			// No even numbers
			body:                 "1 3 5",
			query:                "n=1",
			expectedStatus:       http.StatusBadRequest,
			expectedBody:         "no even numbers provided",
			expectedErrorMessage: "",
		},
		{
			// No odd numbers
			body:                 "2 4 6",
			query:                "n=2",
			expectedStatus:       http.StatusBadRequest,
			expectedBody:         "no odd numbers provided",
			expectedErrorMessage: "",
		},
	}

	for _, testcase := range tests {
		t.Run("Query:"+testcase.query, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/average?"+testcase.query, strings.NewReader(testcase.body))
			w := httptest.NewRecorder()

			AverageHandler(w, req)

			resp := w.Result()
			respBody, _ := io.ReadAll(resp.Body)

			if testcase.expectedErrorMessage == "" {
				if string(respBody) != testcase.expectedBody {
					t.Errorf("expected response body %v; got %v", testcase.expectedBody, string(respBody))
				}
				if resp.StatusCode != testcase.expectedStatus {
					t.Errorf("expected status %v; got %v", testcase.expectedStatus, resp.StatusCode)
				}
			} else {
				if string(respBody) != testcase.expectedErrorMessage {
					t.Errorf("expected error message %v; got %v", testcase.expectedErrorMessage, string(respBody))
				}
			}
		})
	}
}

func TestAverageHandlerInvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/average", nil)
	w := httptest.NewRecorder()

	AverageHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status Method Not Allowed; got %v", resp.StatusCode)
	}
}

func BenchmarkAverageHandler(b *testing.B) {
	body := "10 20 30 40"

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/average?n=3", strings.NewReader(body))
		w := httptest.NewRecorder()

		AverageHandler(w, req)
		// Optionally, you can check the result if needed
		w.Result()
	}
}
