/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) 2022, Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package fetch

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lcook/portutil/internal/port"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func feed(url string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)

	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

func Packages(cmd *cobra.Command) (port.Packages, error) {
	maintainer := viper.GetString("maintainer")
	if maintainer == "" {
		return nil, errors.New("No port maintainer provided or found in the configuration file")
	}

	data, err := feed(viper.GetString("portscout") + maintainer)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, fmt.Errorf("No package updates found for maintainer '%s'", maintainer)
	}

	base := viper.GetString("base")

	if strings.HasPrefix(base, "~/") {
		dir, _ := os.UserHomeDir()
		base = filepath.Join(dir, base[2:])
	}

	if _, err := os.Stat(base); os.IsNotExist(err) {
		return nil, fmt.Errorf("Ports base directory '%s' does not exist", base)
	}

	var (
		packages port.Packages
		wgroup   sync.WaitGroup
	)

	origins, _ := cmd.Flags().GetStringSlice("origins")

	for _, item := range data {
		wgroup.Add(1)

		go func(i *gofeed.Item) {
			itemExt := func(val string) string { return i.Extensions["port"][val][0].Value }
			origin := fmt.Sprintf("%s/%s", itemExt("portcat"),
				itemExt("portname"))

			if len(origins) > 0 {
				for _, o := range origins {
					pkg, err := port.NewPackage(base, origin, itemExt("newversion"))
					if err != nil || origin != o {
						continue
					}

					packages.Add(pkg)
				}
			} else {
				pkg, err := port.NewPackage(base, origin, itemExt("newversion"))
				if err != nil {
					return
				}
				if !pkg.IsLatest() {
					packages.Add(pkg)
				}
			}

			wgroup.Done()
		}(item)
	}

	wgroup.Wait()

	if len(packages) < 1 {
		return nil, errors.New("No package updates to display with applied origin filter")
	}

	return packages, nil
}
