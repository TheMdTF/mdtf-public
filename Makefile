EXTENSION := .yaml

# Ensure npm is installed
.PHONY: install
install:
	@which npm >/dev/null 2>&1 || (echo "npm is not installed. Please install Node.js and npm first." && exit 1)

# Generate documentation with quiet output and error handling
.PHONY: generate
generate: install
	-find . -type f -name "*$(EXTENSION)" -exec sh -c 'cd "$$(dirname "$$1")" && npx openapi-to-md "$$(basename "$$1")" readme.md' _ {} \;
