# ======== BUILD =========
build: package-ui
	@./scripts/build/build.sh

build-linux: package-ui
	@./scripts/build/build.sh linux/amd64

build-docker: package-docker

# ======== PACKAGE =========
package-ui:
	@./scripts/package/package-ui.sh

package-docker: build-linux
	docker build -t lifehub:1.0.0 .

# ======== RUN =========
run:
	@./scripts/run/run.sh
