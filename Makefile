EXTENSION := .yaml
COMMAND := widdershins --language_tabs='shell:Shell' 'javascript:JavaScript' 'go:Go' --summary --outfile readme.md

generate:
	find . -type f -name "*$(EXTENSION)" -execdir sh -c '$(COMMAND) "{}" > /dev/null || exit 1' \;
