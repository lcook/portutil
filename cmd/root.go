/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2022, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	BuildMk  string = "build.mk"
	CommitMk string = "commit.mk"
	InfoMk   string = "info.mk"
	TestMk   string = "test.mk"
	UpdateMk string = "update.mk"

	DefaultPortBase   string = "/usr/ports/"
	DefaultScriptBase string = "/usr/local/share/portutil/Mk/"

	Portscout    string = "https://portscout.freebsd.org/"
	PortscoutRSS string = Portscout + "rss/rss.cgi?m="
)

var (
	version string = "dev"

	base       string
	config     string
	maintainer string
	origins    []string
	scripts    string

	rootCmd = &cobra.Command{
		SilenceUsage: true,
		Use:          "portutil",
		Version:      version,
		Short:        "Utility to manage the process of updating FreeBSD ports.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVarP(&base, "base", "b", DefaultPortBase, "Port base directory.")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "Configuration file.")
	rootCmd.PersistentFlags().StringVarP(&maintainer, "maintainer", "m", "", "Package maintainer email.")
	rootCmd.PersistentFlags().StringSliceVarP(&origins, "origins", "o", []string{}, "List of package origins.")
	rootCmd.PersistentFlags().StringVarP(&scripts, "scripts", "s", DefaultScriptBase, "Script directory.")
}

func initConfig() {
	viper.SetDefault("portscout", PortscoutRSS)

	//nolint
	viper.BindPFlag("base", rootCmd.PersistentFlags().Lookup("base"))
	//nolint
	viper.BindPFlag("maintainer", rootCmd.PersistentFlags().Lookup("maintainer"))
	//nolint
	viper.BindPFlag("scripts", rootCmd.PersistentFlags().Lookup("scripts"))

	if config != "" {
		viper.SetConfigName(config)
	} else {
		viper.SetConfigName(".portutil")
		viper.SetConfigType("toml")
		viper.AddConfigPath("$HOME")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config file: ", err.Error())
	}
}
