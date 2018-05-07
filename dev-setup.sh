#!/usr/bin/env bash

cp githooks/* ./.git/hooks/
echo "Create a new branch before making commits, or the push will fail!"
