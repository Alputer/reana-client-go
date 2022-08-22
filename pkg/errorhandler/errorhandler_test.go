/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

package errorhandler

import (
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/spf13/viper"
)

type testApiError struct {
	Payload struct{ Message string }
}

func (e *testApiError) Error() string { return e.Payload.Message }

func TestHandleApiError(t *testing.T) {
	serverURL := "https://localhost:8080"
	viper.Set("server-url", serverURL)
	t.Cleanup(func() {
		viper.Reset()
	})

	urlError := url.Error{}
	apiError := testApiError{Payload: struct{ Message string }{Message: "API Error"}}
	otherError := errors.New("other Error")

	tests := []struct {
		arg  error
		want string
	}{
		{
			arg: &urlError,
			want: fmt.Sprintf(
				"'%s' not found, please verify the provided server URL or check your internet connection",
				serverURL,
			),
		},
		{
			arg:  &apiError,
			want: apiError.Error(),
		},
		{
			arg:  otherError,
			want: otherError.Error(),
		},
	}
	for _, test := range tests {
		got := HandleApiError(test.arg)
		if got.Error() != test.want {
			t.Errorf("Expected %s, got %s", test.want, got)
		}
	}
}
