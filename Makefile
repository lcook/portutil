#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.POSIX:
PROG=		portutil
VERSION=	0.1

LOCALBASE?=	/usr/local
BINDIR=		${LOCALBASE}/bin
SHAREDIR=	${LOCALBASE}/share/${PROG}

HASH!=		git rev-parse --short HEAD
BRANCH!=	git symbolic-ref HEAD | sed 's,refs/heads/,,'
VERSION:=	${BRANCH}/${VERSION}-${HASH}

GO_MODULE=	github.com/lcook/${PROG}
GO_FLAGS=	-v -ldflags "-s -w -X '${GO_MODULE}/cmd.version=${VERSION}'"

all: build
build:
	go build ${GO_FLAGS} -o ${PROG}
format:
	find . -type f -name '#.go' -exec gofmt -w {} +
lint:
	golangci-lint run
clean:
	go clean -x
install:
	mkdir -p ${SHAREDIR}/Mk
	cp -vR Mk/ ${SHAREDIR}/Mk
	cp -v ${PROG} ${BINDIR}
deinstall:
	rm -rfv ${SHAREDIR}
	rm -fv ${LOCALBASE}/bin/${PROG}
.PHONY: all build format lint clean install deinstall
