#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
secdir="$workspace/src/github.com/PeteWorkspace"
if [ ! -L "$secdir/go-nex" ]; then
    mkdir -p "$secdir"
    cd "$secdir"
    ln -s ../../../../../. go-nex
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$secdir/go-nex"
PWD="$secdir/go-nex"

# Launch the arguments with the configured environment.
exec "$@"
