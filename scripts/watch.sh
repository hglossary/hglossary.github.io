#!/usr/bin/env bash

if ! grep nganhdulich <package.json >/dev/null 2>&1; then
    printf '\nERROR: must run this script in the root directory\n'
    exit 1
fi

yarn snowpack dev
