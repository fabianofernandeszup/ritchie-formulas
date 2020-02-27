#!/bin/sh
find formulas -name "*.zip" | while read filename; do unzip -o -d "`dirname "$filename"`" "$filename"; done;
