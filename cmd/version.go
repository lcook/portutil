/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2022, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"github.com/lcook/portutil/internal/fetch"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display local and latest package versions from a given ports maintainer.",
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs, err := fetch.Packages(cmd)
		if err != nil {
			return err
		}
		for _, pkg := range pkgs {
			pkg.Summary()
		}
		return nil
	},
	Aliases: []string{"v", "ver"},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
