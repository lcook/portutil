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

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update local port to match latest remote package version.",
	RunE:    script.RunFunc(UpdateMk),
	Aliases: []string{"u", "up", "upgrade", "update"},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
