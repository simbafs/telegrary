#!/bin/bash

if [ -z "$1" ]; then
	echo -n "Enter release version: "
	read VERSION
fi
VERSION=${1:-$VERSION}
git push
echo "Releasing version $VERSION ..."
git tag -as $VERSION
git push origin --tags
go list -m github.com/simba-fs/telegrary@$VERSION
