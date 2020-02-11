#!/bin/bash

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

init() {
  for formula in $FORMULAS; do
    create_formulas_dir
    find_config_files
    copy_config_files
    copy_formula_bin
  done
}

init
