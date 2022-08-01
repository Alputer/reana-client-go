/*
This file is part of REANA.
Copyright (C) 2022 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

/*
Package validation provides functions that validate given configurations or command flags.

In case of a failed validation, every function in this package returns an error explaining why it failed.
Otherwise, they return nil, meaning that the validation was successful.
*/
package validation

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

// ValidateAccessToken verifies if the access token has been set, ignoring any white spaces.
func ValidateAccessToken(token string) error {
	if strings.TrimSpace(token) == "" {
		return errors.New(
			"please provide your access token by using the -t/--access-token flag, or by setting the REANA_ACCESS_TOKEN environment variable",
		)
	}
	return nil
}

// ValidateServerURL verifies if REANA's server URL has been set, ignoring any white spaces.
func ValidateServerURL(serverURL string) error {
	if strings.TrimSpace(serverURL) == "" {
		return errors.New("please set REANA_SERVER_URL environment variable")
	}
	return nil
}

// ValidateWorkflow verifies if the workflow's name has been set, ignoring any white spaces.
func ValidateWorkflow(workflow string) error {
	if strings.TrimSpace(workflow) == "" {
		return errors.New(
			"workflow name must be provided either with `--workflow` option or with REANA_WORKON environment variable",
		)
	}
	return nil
}

// ValidateChoice verifies if the given argument (arg) is part of the slice of available choices.
// The third parameter, name, is the name of the argument/flag that should be displayed if the validation fails.
func ValidateChoice(arg string, choices []string, name string) error {
	if !slices.Contains(choices, arg) {
		return fmt.Errorf(
			"invalid value for '%s': '%s' is not part of '%s'",
			name,
			arg,
			strings.Join(choices, "', '"),
		)
	}
	return nil
}
