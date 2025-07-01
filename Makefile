EXTENSION := .yaml
COMMAND := widdershins --omitHeader --language_tabs='shell:Shell' 'javascript:JavaScript' 'go:Go' --summary --outfile readme.md

# Recommended to run yamllint first and correct any errors to ensure proper output
# Ensure npm and widdershins are installed
.PHONY: install
install:
	@which npm >/dev/null 2>&1 || (echo "npm is not installed. Please install Node.js and npm first." && exit 1)
	@which widdershins >/dev/null 2>&1 || npm install -g widdershins

# Generate documentation with quiet output and error handling
.PHONY: generate
generate: install
	-find . -type f -name "*$(EXTENSION)" -execdir sh -c '$(COMMAND) "{}" > /dev/null || exit 1' \;

.PHONY: lint
lint:
	-find . -iname *.yaml | xargs yamllint
