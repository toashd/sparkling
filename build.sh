#!/usr/bin/env sh

set -e

repo_path="github.com/toashd"

revision=$( git rev-parse --short HEAD 2> /dev/null || echo 'unknown' )
branch=$( git rev-parse --abbrev-ref HEAD 2> /dev/null || echo 'unknown' )
host=$( hostname -f )
build_date=$( date +%Y%m%d%H%M%S )
go_version=$( go version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/' )

if [ "$(go env GOOS)" = "windows" ]; then
	ext=".exe"
fi

ldflags="
  -X ${repo_path}/sparkling.Build=${revision}
  -X ${repo_path}/sparkling.BuildDate=${build_date}
  -X ${repo_path}/sparkling.GoVersion=${go_version}"

echo "> build sparkling"

go build -ldflags "${ldflags}" -o bin/sparkling${ext} ./cmd

exit 0
