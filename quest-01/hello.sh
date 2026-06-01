#!/bin/bash

GIT_USERNAME=$(git config user.name)

GIT_USERNAME=${GIT_USERNAME:-"im-aderm"}

echo "hello $GIT_USERNAME!"
