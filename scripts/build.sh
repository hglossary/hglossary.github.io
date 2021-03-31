#!/usr/bin/env bash
set -eo pipefail

if ! grep nganhdulich <package.json >/dev/null 2>&1; then
    printf '\nERROR: must run this script in the root directory\n'
    exit 1
fi

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

printf '\nOK ✔\n'