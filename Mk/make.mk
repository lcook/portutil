#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
_MAKE?=		/usr/bin/make
_MAKE_ARGS=	-C ${PACKAGE_DIR} \
		-DNO_DIALOG
_MAKE_CMD=	${_MAKE} ${_MAKE_ARGS}
