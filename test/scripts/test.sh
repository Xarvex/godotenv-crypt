#!/bin/sh

script_dir="$(readlink -f "$(dirname "$(readlink -f "${0}")")"/..)"
cd "${script_dir}" || exit 1

cp ../.env .env

mkdir -p bin
go build -o bin/test

bin/test
