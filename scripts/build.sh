#!/bin/sh
# Author: Stanislav Yakush <st.yakush@yandex.ru>

set -e

check_variable()
{
    local name="$1"
    local value=$2

    if [ -z $value ]; then
        echo "Variable \"$name\" is not set"
        exit 1
    fi
}

check_variable "Application name" $APP_NAME
check_variable "Application directory" $APP_DIR
check_variable "Build directory" $BUILD_DIR

cd $APP_DIR
go build -o $BUILD_DIR/$APP_NAME