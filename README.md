pktd
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://Copyfree.org)
![Master branch build Status](https://github.com/pkt-cash/pktd/actions/workflows/go.yml/badge.svg?branch=master)
![Develop branch build Status](https://github.com/pkt-cash/pktd/actions/workflows/go.yml/badge.svg?branch=develop)

`pktd` is the primary full node *PKT Cash* implementation, written in Go.

This is a mirror of the original project designed to work with the PacketCrypt Blockchain

## Building

Using `git`, clone the project from the repository:

`git clone https://github.com/PKTW3/pktd`

Use the `./do` shell script to build `pktd`, `pktwallet`, and `pktctl`.

NOTE: It is highly recommended to use only the toolchain Google distributes
at the [official Go homepage](https://golang.org/dl). Go toolchains provided
by Linux distributions often use different defaults or apply non-standard
patches to the official sources, usually to meet distribution-specific
requirements (for example, Red Hat backports, security fixes, and provides
a different default linker configuration vs. the upstream Google Go package.)

Support can only be provided for binaries compiled from unmodified sources,
using the official (upstream) Google Golang toolchain. We unfortunately are
unable to test and support every distribution specific combination. 

The official Google Golang installer for Linux is always available 
for download [here](https://storage.googleapis.com/golang/getgo/installer_linux).

## Documentation

The documentation for `pktd` is work-in-progress, and available in the [docs] folder.

## License

`pktd` is licensed under the [Copyfree](http://Copyfree.org) ISC License.
