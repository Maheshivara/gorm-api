#!/bin/bash

# Get the directory of the script
if [[ $BASH_SOURCE = */* ]]; then
    script_dir=${BASH_SOURCE%/*}
else
    script_dir=.
fi

# Generate mock for services
for file in $script_dir/src/services/*.go; do
  base_name=$(basename $file)
  mockgen -source="$file" -destination="$script_dir/src/services/mocks/$base_name"
done