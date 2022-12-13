package port

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/muesli/termenv"
)

const (
	PackageDefault int = iota
	PackageCargo
	PackageGo
)

type Package struct {
	Origin     string
	Prefix     string
	Version    string
	Maintainer string
	Type       int
	Latest     string
}

func (p *Package) TypeString() string {
	switch p.Type {
	case PackageCargo:
		return "CARGO"
	case PackageGo:
		return "GO"
	}

	return "DEFAULT"
}

func (p *Package) IsLatest() bool {
	latest := strings.TrimPrefix(p.Latest, p.Prefix)
	pkgver := func(s1, s2 string) bool {
		out, _ := exec.Command("pkg", "version", "-t", s1, s2).Output()
		return strings.TrimRight(string(out), "\n") != "<"
	}

	return p.Version == latest || pkgver(p.Version, latest)
}

func (p *Package) Summary() {
	fmt.Printf("[%s] %20s: %s -> %s\n", p.TypeString()[0:1], p.Origin,
		termenv.String(p.Version).
			Foreground(termenv.ColorProfile().Color("#F3713D")),
		termenv.String(strings.TrimPrefix(p.Latest, p.Prefix)).
			Foreground(termenv.ColorProfile().Color("#6EB77F")),
	)
}

func NewPackage(base, origin, latest string) (Package, error) {
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	pkgPath := base + origin

	if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
		return Package{}, err
	}

	makeVar := func(val string) string {
		out, err := exec.Command("make", "-C", pkgPath, "-V", val).Output()
		if err != nil {
			return ""
		}

		return strings.TrimRight(string(out), "\n")
	}

	pkgType := PackageDefault
	if strings.Contains(makeVar("USES"), "cargo") {
		pkgType = PackageCargo
	}

	if strings.Contains(makeVar("USES"), "go:modules") &&
		makeVar("GO_MODULE") == "" {
		pkgType = PackageGo
	}

	return Package{
		Origin:     origin,
		Prefix:     makeVar("DISTVERSIONPREFIX"),
		Version:    makeVar("DISTVERSION"),
		Maintainer: makeVar("MAINTAINER"),
		Type:       pkgType,
		Latest:     latest,
	}, nil
}

type Packages []Package

func (p *Packages) Add(pkg Package) {
	*p = append(*p, pkg)
}
