#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.include "guard.mk"

default:
	@echo
	@echo ------------------------------------
	@echo Origin: ${PACKAGE_ORIGIN}
	@echo Local: ${PACKAGE_VERSION}
	@echo Latest: ${PACKAGE_LATEST}
	@echo Maintainer: ${PACKAGE_MAINTAINER}
	@echo Type: ${PACKAGE_TYPE}
	@echo Directory: ${PACKAGE_DIR}
	@echo ------------------------------------
	@echo
