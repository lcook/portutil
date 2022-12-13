- [Requirements](#requirements)
- [Installation](#installation)
- [Commands](#commands)
  - [build](#build)
  - [commit](#commit)
  - [info](#info)
  - [test](#test)
  - [update](#update)
  - [version](#version)
- [License](#lICENSE)

Small command-line utility to ease the process in updating and committing
FreeBSD packages, but namely an ad-hoc project for FreeBSD port maintainers
like myself.

## Requirements

Running FreeBSD system with the following packages installed: `pkg install -f modules2tuple porttools`

## Installation

From the terminal, run the following:

```console
# make install
```

This will build and install the resulting Go binary, as-well-as accompanying
Makefiles residing under [Mk](Mk). There is an [example configuration file](.portutil.example),
that when placed in your home directory (named as `.portutil`), is used for
default values including port maintainer and base ports directory.

## Commands

Below are brief descriptions of the commands available. 

### build

Build a local port, equivalent to running `make -C /path/to/port`.

### commit

Commit local changes made to a port (typically ran after the `update` 
process), whereby the commit header becomes the according port-name and
version in which it was updated to, for example: `devel/gh: Update to 1.3.4`.  

### info

Print information about a local port and it's remote counterpart (fetched
from [portscout](https://portscout.freebsd.org/)), for example:
```console
------------------------------------
Origin: misc/broot
Local: 1.14.3
Latest: v1.17.0
Maintainer: lcook@FreeBSD.org
Type: CARGO
Directory: /home/lcook/dev/git/freebsd/ports/misc/broot
------------------------------------
```

### test

Run a local ports test-suite, equivalent to running `make test -C /path/to/port`.

### update

Update a local port to it's latest version gathered via [portscout](https://portscout.freebsd.org/).

Currently, ports using either the `USES=go:modules` or `USES=cargo`
directive are "special cases" handled within [update.mk](Mk/update.mk),
else simply `DISTVERSION` is amended to the latest version, followed by
`make makesum -C /path/to/port`, regenerating `distinfo`. Note that
`pkg-plist` is to be dealt with manually in most cases, similarly anything
that routinely requires human intervention (such as static variables used
for build-version output), outside the scope of this utility.

### version 

Display latest versions of local ports, for example:
```console
[D]         x11-wm/berry: 0.1.9 -> 0.1.12
[D]        devel/git-bug: 0.7.2 -> 0.8.0
[D]     sysutils/hostctl: 1.1.2 -> 1.1.3
[D]          textproc/ov: 0.11.2 -> 0.12.0
[D]              www/lux: 0.15.0 -> 0.16.0
[D]        sysutils/smug: 0.3.2 -> 0.3.3
[D]               net/k6: 0.37.0 -> 0.41.0
[C]     sysutils/fselect: 0.8.0 -> 0.8.1
```

Ports marked "[D]" are "default", meaning no special-case handling, while
"[C]" and "[G]" equate to [C]argo and [G]o-based ports respectively.

## License

[BSD 2-Clause](LICENSE)
