package main

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"time"
	"context"
	"errors"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

type SpyResponseWriter struct {
	written bool
}


func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}




func TestServer(t *testing.T) {

	t.Run("Happy path", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t} 
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)


		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)


		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}

		store.assertWasCancelled()

	})
	
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

