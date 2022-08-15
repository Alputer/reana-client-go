/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reanahub/reana-client-go/client/operations"
	"reanahub/reana-client-go/utils"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

var logsPathTemplate = "/api/workflows/%s/logs"

func TestLogs(t *testing.T) {
	workflowName := "my_workflow"
	logsTemplate := `{
			"workflow_logs": "workflow logs",
			"job_logs": {
				"1": {
					"workflow_uuid": "%s",
					"job_name": "%s",
					"compute_backend": "Kubernetes",
					"backend_job_id": "backend1",
					"docker_img": "docker1",
					"cmd": "ls",
					"status": "finished",
					"logs": "%s",
					"started_at": "2022-07-20T12:09:09",
					"finished_at": "2022-07-20T19:09:09"
				},
				"2": {
					"workflow_uuid": "workflow_2",
					"job_name": "job2",
					"compute_backend": "Slurm",
					"backend_job_id": "backend2",
					"docker_img": "docker2",
					"cmd": "cd folder",
					"status": "running",
					"logs": "workflow 2 logs",
					"started_at": "2022-07-21T12:09:09",
					"finished_at": "2022-07-21T19:09:09"
				}
			},
			"engine_specific": "engine logs"
		}`
	successResponseRaw, _ := json.Marshal(operations.GetWorkflowLogsOKBody{
		Logs:         fmt.Sprintf(logsTemplate, "workflow_1", "job1", "workflow 1 logs"),
		User:         "user",
		WorkflowID:   "my_workflow_id",
		WorkflowName: "my_workflow",
	})
	successResponse := string(successResponseRaw)

	incompleteResponseRaw, _ := json.Marshal(operations.GetWorkflowLogsOKBody{
		Logs:         fmt.Sprintf(logsTemplate, "", "", ""),
		User:         "user",
		WorkflowID:   "my_workflow_id",
		WorkflowName: "my_workflow",
	})
	incompleteResponse := string(incompleteResponseRaw)

	tests := map[string]TestCmdParams{
		"default": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: successResponse,
			statusCode:     http.StatusOK,
			args:           []string{"-w", workflowName},
			expected: []string{
				"Workflow engine logs", "workflow logs",
				"Engine internal logs", "engine logs",
				"Job logs", "Step:", "job1", "Workflow ID:", "workflow_1",
				"Compute backend:", "Kubernetes", "Job ID:", "backend1",
				"Docker image:", "docker1", "Command:", "ls", "Status:", "finished",
				"Started:", "2022-07-20T12:09:09", "Finished:", "2022-07-20T19:09:09",
				"Logs:", "workflow 1 logs", "Step:", "job2",
			},
		},
		"without log information": {
			serverPath: fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: `{
				"logs": "{}",
				"user": "user",
				"workflow_id": "my_workflow_id",
				"workflow_name": "my_workflow"
			}`,
			statusCode: http.StatusOK,
			args:       []string{"-w", workflowName},
			unwanted: []string{
				"Workflow engine logs", "Engine internal logs",
				"Job logs", "Step:", "job1",
			},
		},
		"json": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: successResponse,
			statusCode:     http.StatusOK,
			args:           []string{"-w", workflowName, "--json"},
			expected: []string{
				"\"workflow_logs\": \"workflow logs\"",
				"\"job_logs\": {", "\"1\": {",
				"\"workflow_uuid\": \"workflow_1\"",
				"\"logs\": \"workflow 1 logs\"",
				"\"engine_specific\": \"engine logs\"",
			},
		},
		"with filters": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: successResponse,
			statusCode:     http.StatusOK,
			args:           []string{"-w", workflowName, "--filter", "compute_backend=kubernetes"},
			expected:       []string{"Step: job1"},
			unwanted:       []string{"Step: job2"},
		},
		"missing step names": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: successResponse,
			statusCode:     http.StatusOK,
			args:           []string{"-w", workflowName, "--filter", "step=3"},
			expected: []string{
				"ERROR:", "The logs of step(s) 3 were not found, check for spelling mistakes in the step names",
			},
		},
		"missing fields": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: incompleteResponse,
			statusCode:     http.StatusOK,
			args:           []string{"-w", workflowName, "--filter", "compute_backend=kubernetes"},
			expected:       []string{"Step: 1", "Step 1 emitted no logs."},
			unwanted: []string{
				"job1",
				"Workflow ID:", "workflow_1",
				"Logs:", "workflow 1 logs",
			},
		},
		"malformed filters": {
			serverPath: fmt.Sprintf(logsPathTemplate, workflowName),
			args:       []string{"-w", workflowName, "--filter", "name"},
			expected: []string{
				"wrong input format. Please use --filter filter_name=filter_value",
			},
			wantError: true,
		},
		"unexisting workflow": {
			serverPath:     fmt.Sprintf(logsPathTemplate, "invalid"),
			serverResponse: `{"message": "REANA_WORKON is set to invalid, but that workflow does not exist."}`,
			statusCode:     http.StatusNotFound,
			args:           []string{"-w", "invalid"},
			expected: []string{
				"REANA_WORKON is set to invalid, but that workflow does not exist.",
			},
			wantError: true,
		},
		"invalid page": {
			serverPath:     fmt.Sprintf(logsPathTemplate, workflowName),
			serverResponse: `{"message": "Field 'page': Must be at least 1."}`,
			statusCode:     http.StatusBadRequest,
			args:           []string{"-w", workflowName, "--page", "0"},
			expected:       []string{"Field 'page': Must be at least 1."},
			wantError:      true,
		},
	}

	for name, params := range tests {
		t.Run(name, func(t *testing.T) {
			params.cmd = "logs"
			testCmdRun(t, params)
		})
	}
}

func TestParseLogsFilters(t *testing.T) {
	tests := map[string]struct {
		filterInput []string
		wantError   bool
	}{
		"valid filters": {
			filterInput: []string{
				"compute_backend=kubernetes",
				"status=running",
				"docker_img=docker",
			},
		},
		"invalid filter key": {
			filterInput: []string{"invalid=kubernetes"},
			wantError:   true,
		},
		"invalid status filter": {
			filterInput: []string{"status=invalid"},
			wantError:   true,
		},
		"invalid compute backend filter": {
			filterInput: []string{"compute_backend=invalid"},
			wantError:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := parseLogsFilters(test.filterInput)
			if test.wantError && err == nil {
				t.Fatalf(
					"expected parseLogsFilters(%#v) to return an error but didn't",
					test.filterInput,
				)
			}
			if !test.wantError && err != nil {
				t.Fatalf(
					"parseLogsFilters(%#v) returned an unexpected error: %s",
					test.filterInput,
					err.Error(),
				)
			}
		})
	}

	t.Run("expected filter keys", func(t *testing.T) {
		filters, err := parseLogsFilters([]string{})
		if err != nil {
			t.Fatalf(
				"parseLogsFilters(%#v) returned an unexpected error: %s",
				[]string{},
				err.Error(),
			)
		}
		if !slices.Equal(filters.SingleFilterKeys, logsSingleFilters) {
			t.Fatalf(
				"expected single filter keys to be %#v but got %#v",
				logsSingleFilters,
				filters.SingleFilterKeys,
			)
		}
		if !slices.Equal(filters.MultiFilterKeys, logsMultiFilters) {
			t.Fatalf(
				"expected multi filter keys to be %#v but got %#v",
				logsMultiFilters,
				filters.MultiFilterKeys,
			)
		}
	})
}

func TestFilterJobLogs(t *testing.T) {
	tests := map[string]struct {
		filterInput []string
		wantLogs    map[string]jobLogItem
	}{
		"no filters": {
			filterInput: []string{},
			wantLogs: map[string]jobLogItem{
				"1": {ComputeBackend: "Kubernetes", Status: "running", DockerImg: "docker"},
				"2": {ComputeBackend: "Slurm", Status: "created", DockerImg: "docker2"},
				"3": {ComputeBackend: "HTCondor", Status: "created", DockerImg: "docker3"},
			},
		},
		"single filter": {
			filterInput: []string{"status=created"},
			wantLogs: map[string]jobLogItem{
				"2": {ComputeBackend: "Slurm", Status: "created", DockerImg: "docker2"},
				"3": {ComputeBackend: "HTCondor", Status: "created", DockerImg: "docker3"},
			},
		},
		"multiple filters": {
			filterInput: []string{"status=created", "compute_backend=slurm", "docker_img=docker2"},
			wantLogs: map[string]jobLogItem{
				"2": {ComputeBackend: "Slurm", Status: "created", DockerImg: "docker2"},
			},
		},
		"uppercase compute_backend": {
			filterInput: []string{"compute_backend=KUBERNETES"},
			wantLogs: map[string]jobLogItem{
				"1": {ComputeBackend: "Kubernetes", Status: "running", DockerImg: "docker"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filters, err := utils.NewFilters(logsSingleFilters, logsMultiFilters, test.filterInput)
			if err != nil {
				t.Fatalf("utils.NewFilters returned an unexpected error: %s", err.Error())
			}

			jobLogs := map[string]jobLogItem{
				"1": {ComputeBackend: "Kubernetes", Status: "running", DockerImg: "docker"},
				"2": {ComputeBackend: "Slurm", Status: "created", DockerImg: "docker2"},
				"3": {ComputeBackend: "HTCondor", Status: "created", DockerImg: "docker3"},
			}
			err = filterJobLogs(&jobLogs, filters)
			if err != nil {
				t.Fatalf("filterJobLogs returned an unexpected error: %s", err.Error())
			}
			if !reflect.DeepEqual(jobLogs, test.wantLogs) {
				t.Errorf("expected %#v, got %#v", test.wantLogs, jobLogs)
			}
		})
	}
}
