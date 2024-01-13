#!/bin/sh

script_dir="$(dirname "${0}")/.."

mkdir -p "${script_dir}/bin"
go build "${script_dir}" -o "${script_dir}bin/test"

"${script_dir}/bin/test"
