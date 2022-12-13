#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.include "guard.mk"

_GIT=		/usr/local/bin/git
_GIT_CMD=	${_GIT} -C ${PACKAGE_DIR}
default:
	${_GIT_CMD} add .
	${_GIT_CMD} commit -m "${PACKAGE_ORIGIN}: Update to ${PACKAGE_LATEST}" 
