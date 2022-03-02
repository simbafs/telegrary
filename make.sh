#!/bin/bash

case $1 in 
	"release"|*)
		if [ -z "$2" ]; then
			echo -n "Enter release version: "
			read VERSION
		fi
		VERSION=${2:-$VERSION}
		echo "Releasing version $VERSION ..."
		git tag -as $VERSION 
		git push origin --tags
		go list -m github.com/simba-fs/telegrary@$VERSION
		;;
esac
