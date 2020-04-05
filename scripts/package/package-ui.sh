#!/usr/bin/env bash
# Do not use this script manually. Use Makefile.

rm -f service/rice-box.go
rice embed-go -i service/server.go
