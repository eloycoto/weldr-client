// Copyright 2021 by Red Hat, Inc. All rights reserved.
// Use of this source is goverend by the Apache License
// that can be found in the LICENSE file.

package compose

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/osbuild/weldr-client/v2/cmd/composer-cli/root"
)

var (
	resultsCmd = &cobra.Command{
		Use:   "results UUID",
		Short: "Get a tar of the the results for the compose",
		Long:  "Get a tar of the the results for the compose",
		RunE:  getResults,
		Args:  cobra.ExactArgs(1),
	}
	savePath string
)

func init() {
	resultsCmd.Flags().StringVarP(&savePath, "filename", "", "", "Optional path and filename to save tar into")
	composeCmd.AddCommand(resultsCmd)
}

func getResults(cmd *cobra.Command, args []string) error {
	fn, resp, err := root.Client.ComposeResultsPath(args[0], savePath)
	if err != nil {
		return root.ExecutionError(cmd, "Results error: %s", err)
	}
	if resp != nil && !resp.Status {
		return root.ExecutionErrors(cmd, resp.Errors)
	}

	fmt.Println(fn)

	return nil
}
