#!/bin/bash -e

# make debian/debpackages/ directory if not exists
DEB_DESTINATION=debian/debpackages
if [ ! -d ${DEB_DESTINATION} ]; then
  mkdir ${DEB_DESTINATION}
fi

# go build creates srv-stockwatcherservice-go binary file
go build

# extract package name and current version from changelog
read -r HEADERLINE < debian/changelog
PACKAGE_NAME=`echo ${HEADERLINE// /} |  cut -d "(" -f1`
PACKAGE_VERSION_SUBSTR=`echo ${HEADERLINE// /} |  cut -d "(" -f2`
PACKAGE_VERSION=`echo ${PACKAGE_VERSION_SUBSTR} |  cut -d ")" -f1`

# pack binary as debian package with fpm
echo 'INFO: build debian package for' $PACKAGE_NAME 'version' $PACKAGE_VERSION
fpm -s dir -t deb -n ${DEB_DESTINATION}/${PACKAGE_NAME} -v ${PACKAGE_VERSION} ${PACKAGE_NAME}=/usr/bin/
