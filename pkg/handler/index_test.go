package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		scenario string
		url      string
		fn       func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			scenario: "Test when no query parameter is provided",
			url:      "/",
			fn:       testIndexNoQuery,
		},
		{
			scenario: "Test when favorite query parameter is given",
			url:      "/?favoriteTree=baobab",
			fn:       testIndexFavoriteTreeParam,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.url, nil)

			w := httptest.NewRecorder()
			Index(w, req)

			tc.fn(t, w)
		})
	}
}

func testIndexNoQuery(t *testing.T, rec *httptest.ResponseRecorder) {
	testCommonHTTPMeta(t, rec.Result())

	actual := extractBodyString(rec.Result().Body)
	expected := defaultBody

	if !strings.Contains(actual, expected) {
		t.Errorf("unexpected body content: got %v want %v",
			actual, expected)
	}
}

func testIndexFavoriteTreeParam(t *testing.T, rec *httptest.ResponseRecorder) {
	testCommonHTTPMeta(t, rec.Result())

	actual := extractBodyString(rec.Result().Body)
	expected := treePrefix + "baobab"

	if !strings.Contains(actual, expected) {
		t.Errorf("unexpected body content: got %v want %v",
			actual, expected)
	}
}

func testCommonHTTPMeta(t *testing.T, res *http.Response) {
	expectedCode := http.StatusOK
	if status := res.StatusCode; status != expectedCode {
		t.Errorf("wrong status code: got %v want %v",
			status, expectedCode)
	}

	expectedType := "text/html; charset=utf-8"
	if cType := res.Header.Get("Content-Type"); cType != expectedType {
		t.Errorf("wrong content-type: got %v want %v",
			cType, expectedType)
	}
}

func extractBodyString(rc io.ReadCloser) string {
	bodyBytes, _ := ioutil.ReadAll(rc)
	return string(bodyBytes)
}
