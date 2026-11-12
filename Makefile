EXTENSION := .yaml
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# Ensure npm is installed and local dependencies are present
.PHONY: install
install:
	@which npm >/dev/null 2>&1 || (echo "npm is not installed. Please install Node.js and npm first." && exit 1)
	@test -f "$(ROOT_DIR)node_modules/.bin/openapi-to-md" || (cd "$(ROOT_DIR)" && npm install)

# Generate documentation with quiet output and error handling
.PHONY: generate
generate: install
	-find . -type f -name "*$(EXTENSION)" -exec sh -c 'cd "$$(dirname "$$1")" && "$(ROOT_DIR)node_modules/.bin/openapi-to-md" "$$(basename "$$1")" readme.md' _ {} \;
