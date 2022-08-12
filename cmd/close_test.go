/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

package cmd

import (
	"fmt"
	"net/http"
	"testing"
)

var closePathTemplate = "/api/workflows/%s/close/"

func TestClose(t *testing.T) {
	tests := map[string]TestCmdParams{
		"success": {
			serverPath:     fmt.Sprintf(closePathTemplate, "my_workflow"),
			serverResponse: "{}",
			statusCode:     http.StatusOK,
			args:           []string{"-w", "my_workflow"},
			expected: []string{
				"Interactive session for workflow my_workflow was successfully closed",
			},
		},
		"error": {
			serverPath:     fmt.Sprintf(closePathTemplate, "my_workflow"),
			serverResponse: `{"message": "Workflow - my_workflow has no open interactive session."}`,
			statusCode:     http.StatusNotFound,
			args:           []string{"-w", "my_workflow"},
			expected:       []string{"Workflow - my_workflow has no open interactive session."},
		},
	}

	for name, params := range tests {
		t.Run(name, func(t *testing.T) {
			params.cmd = "close"
			testCmdRun(t, params)
		})
	}
}