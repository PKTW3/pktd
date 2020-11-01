#!/usr/bin/env sh
die() { printf '%s\n' "Error: ${*:?}" >&2; exit 1; }
build() { l="${1:?}"; printf '%s\n' "Building ${l?${unset:?}}"; go build -trimpath -ldflags="${PKTD_LDFLAGS:?${unset:?}}" -o "${bindir:?${unset:?}}"/"${l?${unset:?}}" || die "Failed building ${l?${unset:?}}"; }
export GO111MODULE="on" && export unset="Error: Variable is unset; aborting."
export bindir="./bin" && export PKTD_TESTFLAGS="-count=1 -cover -parallel=1"
PKTD_GIT_ID=$(git update-index -q --refresh 2>/dev/null; git describe --tags HEAD 2>/dev/null)
if ! git diff --quiet 2>/dev/null; then
    if [ -n "${PKT_FAIL_DIRTY:-}" ]; then { git diff 2>/dev/null; die "Build is dirty, aborting."; }; fi
    export PKTD_GIT_ID="${PKTD_GIT_ID:?${unset:?}}-dirty"
fi
export PKTD_LDFLAGS="-X github.com/pkt-cash/pktd/pktconfig/version.appBuild=${PKTD_GIT_ID:?${unset:?}}"
mkdir -p "${bindir:?${unset:?}}" || die "Failed to create output directory; aborting."
build pktd && build pktwallet && build pktctl
printf '%s\n' "Running tests"; # shellcheck disable=SC2086
go test ${PKTD_TESTFLAGS:?${unset:?}} ./... || die "One or more tests failed."
"${bindir?${unset:?}}/pktd" --version || die "Unable to run compiled pktd executable."; # shellcheck disable=SC2250
printf '%s\n' "Success! $( (cd "${bindir:?${unset:?}}" 2>/dev/null && d=$(pwd -P 2>/dev/null) && printf '%s\n' "Compiled output is located at ${d:?${bindir:?$unset}}." 2>/dev/null) 2>/dev/null )"

