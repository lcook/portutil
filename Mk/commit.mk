#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.include "guard.mk"

_GIT=		/usr/local/bin/git
_GIT_CMD=	${_GIT}
default:
	@(cd ${PACKAGE_DIR}; ${_GIT_CMD} add .; \
		${_GIT_CMD} commit -m "${PACKAGE_ORIGIN}: Update to ${PACKAGE_LATEST}")
