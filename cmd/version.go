/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"fmt"
	"strings"

	"github.com/lcook/portutil/internal/fetch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	format string
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
			if format != "" {
				base := viper.GetString("base")
				if !strings.HasSuffix(base, "/") {
					base += "/"
				}
				msg := format
				fmts := map[string]string{
					"%o": pkg.Origin,
					"%p": pkg.Prefix,
					"%v": pkg.Version,
					"%m": pkg.Maintainer,
					"%l": pkg.Latest,
					"%d": base + pkg.Origin,
				}

				for k, v := range fmts {
					msg = strings.Replace(msg, k, v, -1)
				}

				fmt.Println(msg)
			} else {
				pkg.Summary()
			}
		}
		return nil
	},
	Aliases: []string{"v", "ver"},
}

func init() {
	versionCmd.Flags().StringVarP(&format, "format", "f", "", "Format output string. (%o, %p, %v, %m, %l, %d)")
	rootCmd.AddCommand(versionCmd)
}
