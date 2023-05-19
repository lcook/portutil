/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"github.com/lcook/portutil/internal/script"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run local port test suite.",
	RunE:  script.RunFunc(TestMk),
}

func init() {
	rootCmd.AddCommand(testCmd)
}
