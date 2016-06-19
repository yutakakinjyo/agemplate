# Project information
TOP        = $(shell pwd)
REPOSITORY = $(shell basename $(TOP))
VERSION    = $(shell grep "const Version " $(TOP)/version.go | sed -E 's/.*"(.+)"$$/\1/')

# Build information
OUTPUT    = agemplate
BUILDDIR  = $(TOP)/pkg
XC_OS     = "darwin linux windows"
XC_ARCH   = "amd64"
DISTDIR   = $(BUILDDIR)/dist/$(VERSION)

help:
	@echo "Please type: make [target]"
	@echo "  build        Build $(OUTPUT) in to the pkg directory"
	@echo "  dist         Ship packages to release"
	@echo "  release      Create tag ($(VERSION)) and upload binaries to GitHub"
	@echo "  clean        Cleanup artifacts"
	@echo "  help         Show this help messages"

build:
	@echo "===> Beginning compile..."
	gox -os $(XC_OS) -arch $(XC_ARCH) -output "pkg/{{.OS}}_{{.Arch}}/$(OUTPUT)"

dist: build
	@echo "===> Shipping packages..."
	rm -rf $(DISTDIR)
	mkdir -p $(DISTDIR)
	@for dir in $$(find $(BUILDDIR) -mindepth 1 -maxdepth 1 -type d); do \
		platform=`basename $$dir`; \
		if [ $$platform = "dist" ]; then \
			continue; \
		fi; \
		archive=$(OUTPUT)_$(VERSION)_$$platform; \
		zip -j $(DISTDIR)/$$archive.zip $$dir/*; \
		pushd $(DISTDIR); \
		popd; \
	done

release:
	@echo "===> Publishing to GitHub..."
	ghr $(VERSION) $(DISTDIR)

clean:
	go clean ./...
	rm -rf $(BUILDDIR)

.PHONY: help clean release
