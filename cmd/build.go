/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2022, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"github.com/lcook/portutil/internal/script"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a local port.",
	RunE:  script.RunFunc(BuildMk),
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
