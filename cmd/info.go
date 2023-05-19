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

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Show brief information about a local and remote port.",
	RunE:    script.RunFunc(InfoMk),
	Aliases: []string{"i"},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
