#!/usr/bin/env bash
set -eo pipefail

if ! grep nganhdulich <package.json >/dev/null 2>&1; then
    printf '\nERROR: must run this script in the root directory\n'
    exit 1
fi

mkdir -p src/_tmp
go run ./scripts/convert-data
yarn snowpack build

rm -rf docs || true
mkdir -p docs

# copy only necessary files and directories
files='/(assets|index.html|main.css|main.js)$'
for file in build/*; do
    if echo "$file" | grep -E "$files" >/dev/null 2>&1; then
        mv "$file" "docs/"
    fi
done

# github cname
cp CNAME docs/

# generate github pages
cp static/index.html docs/about.html
cp static/index.html docs/contact.html
mkdir -p docs/w

go run ./scripts/generate-pages

printf '\nOK ✔\n'
