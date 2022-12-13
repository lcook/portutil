#
# SPDX-License-Identifier: BSD-2-Clause
#
# Copyright (c) 2021, Lewis Cook <lcook@FreeBSD.org>
# All rights reserved.
#
.include "guard.mk" 
.include "make.mk"

_TYPES=				CARGO:cargo-crates-merge GO:gomod-vendor
.  for tuple in ${_TYPES}
_MAKE_${tuple:C/(.#):(.*)/\1/}=	${_MAKE_CMD} ${tuple:C/(.*):(.*)/\2/}
.  endfor

_PORTEDIT=	/usr/local/bin/portedit
default:
	@if [ -d "${PACKAGE_DIR}/work" ]; then \
		${_MAKE_CMD} clean; \
	fi
	${_PORTEDIT} set-version -i ${PACKAGE_LATEST} ${PACKAGE_DIR}/Makefile
	${_MAKE_CMD} makesum
.  if ${PACKAGE_TYPE} != ""
.    for tuple in ${_TYPES:C/(.#):(.*)/\1/}
.      if ${PACKAGE_TYPE} == ${tuple}
	${_MAKE_${tuple}}
.        if ${tuple} == GO
		${_MAKE_${tuple}} | ${_PORTEDIT} merge -i ${PACKAGE_DIR}/Makefile
.        endif
	${_MAKE_CMD} makesum
.      endif
.    endfor
.  endif
