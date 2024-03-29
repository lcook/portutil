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

	PackageOld string = "#F3713D"
	PackageNew string = "#6EB77F"
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
	pkgver := func(s1, s2 string) bool {
		out, _ := exec.Command("pkg", "version", "-t", s1, s2).Output()
		return strings.TrimRight(string(out), "\n") != "<"
	}

	return p.Version == p.Latest || pkgver(p.Version, p.Latest)
}

func (p *Package) Summary() {
	fmt.Printf("[%s] %20s: %s -> %s\n", p.TypeString()[0:1], p.Origin,
		termenv.String(p.Version).
			Foreground(termenv.EnvColorProfile().Color(PackageOld)),
		termenv.String(p.Latest).
			Foreground(termenv.EnvColorProfile().Color(PackageNew)),
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

	pkg := Package{
		Origin:     origin,
		Prefix:     makeVar("DISTVERSIONPREFIX"),
		Version:    makeVar("DISTVERSION"),
		Maintainer: makeVar("MAINTAINER"),
		Type:       pkgType,
		Latest:     latest,
	}

	pkg.Latest = strings.TrimPrefix(pkg.Latest, pkg.Prefix)

	return pkg, nil
}

type Packages []Package

func (p *Packages) Add(pkg Package) {
	*p = append(*p, pkg)
}
