package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_SiteChecker(t *testing.T) {
	mockRsp := &MockHttpClient{
		Response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("ok")),
		},
		err: nil,
	}
	config := siteConfig{
		Url:             "http:localhost:1975",
		AcceptableCodes: []int{200},
		Frequency:       time.Second * 3,
	}
	result := make(chan Result, 1)
	t.Run("Success", func(t *testing.T) {
		scheduleCheck(config, mockRsp, result)
		res := <-result
		if !res.Up {
			t.Error("Expected up")
		}
		if res.StatusCode != 200 {
			t.Error("Expected 200")
		}
	})
}

func Test_Checker(t *testing.T) {
	mockRsp := &MockHttpClient{
		Response: &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("ok")),
		},
		err: nil,
	}
	config := siteConfig{
		Url:             "http:localhost:1975",
		AcceptableCodes: []int{200},
		Frequency:       time.Second * 3,
	}
	result := make(chan Result, 1)
	t.Run("Success", func(t *testing.T) {
		checker(config, mockRsp, result)
		res := <-result
		if !res.Up {
			t.Error("Expected up")
		}
		if res.StatusCode != 200 {
			t.Error("Expected 200")
		}
	})
	t.Run("Error case", func(t *testing.T) {
		mockRsp = &MockHttpClient{
			err: fmt.Errorf("error"),
		}
		checker(config, mockRsp, result)
		res := <-result
		if res.Up {
			t.Error("Expected down")
		}
		if res.StatusCode == 200 {
			t.Error("Expected non-200")
		}
	})
}

func Test_HttpClient(t *testing.T) {
	client := &DefaultClient{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer server.Close()
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(bs) != "ok" {
		t.Error(string(bs))
	}
}

type MockHttpClient struct {
	Response *http.Response
	err      error
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	return m.Response, m.err
}
