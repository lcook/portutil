#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
PROG=		portutil
VERSION=	0.1

.if exists(${.CURDIR}/.git)
HASH!=		git rev-parse --short HEAD
BRANCH!=	git symbolic-ref HEAD | sed 's,refs/heads/,,'
VERSION:=	${BRANCH}/${VERSION}-${HASH}
.endif

GO_MODULE=	github.com/lcook/portutil
GO_FLAGS=	-v -ldflags "-s -w -X '${GO_MODULE}/cmd.version=${VERSION}'"

LOCALBASE?=	/usr/local
SHARE=		${LOCALBASE}/share/${PROG}

default: build .PHONY

build: .PHONY
	go build ${GO_FLAGS} -o ${PROG}
clean: .PHONY
	go clean -x
install: build .PHONY
	mkdir -p ${SHARE}/Mk
	cp -vR Mk/ ${SHARE}/Mk
	cp -v ${PROG} ${LOCALBASE}/bin
deinstall: .PHONY
	rm -rfv ${SHARE}
	rm -fv ${LOCALBASE}/bin/${PROG}
format: .PHONY
	find . -type f -name '#.go' -exec gofmt -w {} +
lint: .PHONY
	golangci-lint run
