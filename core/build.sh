#!/bin/sh

mkdir -p build/android
mkdir -p build/ios

while :;do printf .;sleep 1;done &
#Die with parent if we die prematurely
trap "kill $!" EXIT

WORKING_DIR="../$(basename $(pwd))"

stringer -type=ActionType || echo "Please run 'go get golang.org/x/tools/cmd/stringer'"

echo "\nRunning tests ✅"
go test || exit 1

echo "\nBuilding Android .aar 🤖"
gomobile bind -target=android/arm,android/386 -o build/android/rey-core.aar -ldflags="-s -w" -javapkg=com.ronocod.rey $WORKING_DIR # || exit 1

if [ "$(uname)" = "Linux" ]; then
    echo "\nNot building iOS framework as we're on Linux 🐧"
else
    echo "\nBuilding iOS framework "
    gomobile bind -target ios -o build/ios/Core.framework $WORKING_DIR || exit 1
fi

echo "\nDone!"

#Kill the loop and unset the trap or else the pid might get reassigned and we might end up killing a completely different process
kill $! && trap " " EXIT
