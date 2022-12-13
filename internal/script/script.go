package script

import (
	"os"
	"os/exec"
	"strings"

	"github.com/lcook/portutil/internal/fetch"
	"github.com/lcook/portutil/internal/port"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func execute(command, base string, pkgs port.Packages) {
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	for _, pkg := range pkgs {
		env := map[string]string{
			"PACKAGE_ORIGIN":     pkg.Origin,
			"PACKAGE_PREFIX":     pkg.Prefix,
			"PACKAGE_VERSION":    pkg.Version,
			"PACKAGE_MAINTAINER": pkg.Maintainer,
			"PACKAGE_LATEST":     pkg.Latest,
			"PACKAGE_TYPE":       pkg.TypeString(),
			"PACKAGE_DIR":        base + pkg.Origin,
		}

		for key, value := range env {
			//nolint
			os.Setenv(key, value)
		}

		args := strings.Fields(command)
		//nolint
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//nolint
		cmd.Run()
	}
}

func makeCmd(cmd *cobra.Command, script string) error {
	pkgs, err := fetch.Packages(cmd)
	if err != nil {
		return err
	}

	execute("make -f"+viper.GetString("scripts")+script, viper.GetString("base"), pkgs)

	return nil
}

func RunFunc(script string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return makeCmd(cmd, script)
	}
}
