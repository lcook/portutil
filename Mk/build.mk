#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.include "guard.mk"
.include "make.mk"

default:
	@if [ -d "${PACKAGE_DIR}/work" ]; then \
		${_MAKE_CMD} clean; \
	fi
	${_MAKE_CMD} build
