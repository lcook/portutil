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

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Commit local port changes.",
	RunE:    script.RunFunc(CommitMk),
	Aliases: []string{"c", "com", "cm"},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
