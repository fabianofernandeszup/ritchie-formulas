#!/bin/sh a

FORMULAS="$1"

create_formulas_dir() {
  mkdir -p formulas/"$formula"
}

find_config_files() {
  files=$(find "$formula" -type f -name "*config.json")
}

copy_config_files() {
  for file in $files; do
    cp "$file" formulas/"$formula"
  done
}

copy_formula_bin() {
  cp -rf "$formula"/bin formulas/"$formula"
}

create_formula_checksum() {
  find "${formula}"/bin -type f -exec md5sum {} \; | sort -k 2 | md5sum | cut -f1 -d ' ' > formulas/"${formula}.md5"
}

init() {
  for formula in $FORMULAS; do
    create_formulas_dir
    find_config_files
    copy_config_files
    create_formula_checksum
    copy_formula_bin
  done
}

init
