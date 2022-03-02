#!/bin/bash

case $1 in 
	"release")
		if [ -z "$2" ]; then
			echo -n "Enter release version: "
			read VERSION
		fi
		VERSION=${2:-$VERSION}
		echo "Releasing version $VERSION ..."
		git tag -as $VERSION -m "Release $VERSION"
		git push origin --tags
		;;
esac
