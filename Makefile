all: build
.PHONY: all

ifdef OS_GIT_VERSION
SOURCE_GIT_TAG := ${OS_GIT_VERSION}
endif

# Include the library makefile
include $(addprefix ./vendor/github.com/openshift/build-machinery-go/make/, \
	golang.mk \
	targets/openshift/images.mk \
	targets/openshift/rpm.mk \
	targets/openshift/deps-gomod.mk \
)

KUBE_GIT_MINOR_VERSION := "33"
KUBE_GIT_VERSION := "v1.33.3"

GO_LD_EXTRAFLAGS :=-X k8s.io/component-base/version.gitMajor="1" \
                   -X k8s.io/component-base/version.gitMinor=$(KUBE_GIT_MINOR_VERSION) \
                   -X k8s.io/component-base/version.gitVersion=$(KUBE_GIT_VERSION) \
                   -X k8s.io/component-base/version.gitCommit="$(SOURCE_GIT_COMMIT)" \
                   -X k8s.io/component-base/version.buildDate="$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')" \
                   -X k8s.io/component-base/version.gitTreeState="clean" \
                   -X k8s.io/client-go/pkg/version.gitVersion="$(SOURCE_GIT_TAG)" \
                   -X k8s.io/client-go/pkg/version.gitCommit="$(SOURCE_GIT_COMMIT)" \
                   -X k8s.io/client-go/pkg/version.buildDate="$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')" \
                   -X k8s.io/client-go/pkg/version.gitTreeState="$(SOURCE_GIT_TREE_STATE)"

GO_BUILD_PACKAGES :=$(strip \
	./cmd/... \
	$(wildcard ./tools/*) \
)
# These tags make sure we can statically link and avoid shared dependencies
GO_BUILD_FLAGS :=-tags 'include_gcs include_oss containers_image_openpgp gssapi'
GO_BUILD_FLAGS_DARWIN :=-tags 'include_gcs include_oss containers_image_openpgp'
GO_BUILD_FLAGS_WINDOWS :=-tags 'include_gcs include_oss containers_image_openpgp'
GO_BUILD_FLAGS_LINUX_CROSS :=-tags 'include_gcs include_oss containers_image_openpgp'

OUTPUT_DIR :=_output
CROSS_BUILD_BINDIR :=$(OUTPUT_DIR)/bin
RPM_VERSION :=$(shell set -o pipefail && echo '$(SOURCE_GIT_TAG)' | sed -E 's/v([0-9]+\.[0-9]+\.[0-9]+)-.*/\1/')
RPM_EXTRAFLAGS := \
	--define 'local_build true' \
	--define 'os_git_vars ignore' \
	--define 'version $(RPM_VERSION)' \
	--define 'release 1'

IMAGE_REGISTRY :=registry.ci.openshift.org

# This will call a macro called "build-image" which will generate image specific targets based on the parameters:
# $1 - target name
# $2 - image ref
# $3 - Dockerfile path
# $4 - context
$(call build-image,ocp-cli,$(IMAGE_REGISTRY)/ocp/4.2:cli,./images/cli/Dockerfile.rhel,.)

$(call build-image,ocp-cli-artifacts,$(IMAGE_REGISTRY)/ocp/4.2:cli-artifacts,./images/cli-artifacts/Dockerfile.rhel,.)

$(call verify-golang-versions,images/cli/Dockerfile.rhel)

image-ocp-cli-artifacts: image-ocp-cli

$(call build-image,ocp-deployer,$(IMAGE_REGISTRY)/ocp/4.2:deployer,./images/deployer/Dockerfile.rhel,.)
image-ocp-deployer: image-ocp-cli

$(call build-image,ocp-recycler,$(IMAGE_REGISTRY)/ocp/4.2:recycler,./images/recycler/Dockerfile.rhel,.)
image-ocp-recycler: image-ocp-cli

oc: GO_BUILD_PACKAGES :=./cmd/oc
oc: build
.PHONY: oc

update: update-generated-completions
.PHONY: update

verify: verify-cli-conventions verify-generated-completions verify-kube-version
.PHONY: verify

verify-cli-conventions:
	go run ./tools/clicheck
.PHONY: verify-cli-conventions

update-generated-completions: build
	hack/update-generated-completions.sh
.PHONY: update-generated-completions

verify-generated-completions: build
	hack/verify-generated-completions.sh
.PHONY: verify-generated-completions

verify-kube-version: build
	KUBE_GIT_VERSION=$(KUBE_GIT_VERSION) hack/verify-kube-version.sh
.PHONY: verify-kube-version

generate-docs:
	go run ./tools/gendocs
.PHONY: generate-docs

generate-docs-admin:
	go run ./tools/gendocs --admin
.PHONY: generate-docs-admin

generate-docs-microshift:
	go run ./tools/gendocs --microshift
.PHONY: generate-docs-microshift

generate-docs-admin-microshift:
	go run ./tools/gendocs --admin --microshift
.PHONY: generate-docs-admin-microshift

generate-versioninfo:
	SOURCE_GIT_TAG=$(SOURCE_GIT_TAG) hack/generate-versioninfo.sh
.PHONY: generate-versioninfo

cross-build-darwin-amd64:
	+@GOOS=darwin GOARCH=amd64 $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_DARWIN)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/darwin_amd64
.PHONY: cross-build-darwin-amd64

cross-build-darwin-arm64:
	+@GOOS=darwin GOARCH=arm64 $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_DARWIN)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/darwin_arm64
.PHONY: cross-build-darwin-arm64

cross-build-windows-amd64: generate-versioninfo
	+@GOOS=windows GOARCH=amd64 $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_WINDOWS)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/windows_amd64
	$(RM) cmd/oc/oc.syso
.PHONY: cross-build-windows-amd64

cross-build-linux-amd64:
	+@GOOS=linux GOARCH=amd64 $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_LINUX_CROSS)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/linux_amd64
.PHONY: cross-build-linux-amd64

cross-build-linux-arm64:
	+@GOOS=linux GOARCH=arm64 $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_LINUX_CROSS)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/linux_arm64
.PHONY: cross-build-linux-arm64

cross-build-linux-ppc64le:
	+@GOOS=linux GOARCH=ppc64le $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_LINUX_CROSS)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/linux_ppc64le
.PHONY: cross-build-linux-ppc64le

cross-build-linux-s390x:
	+@GOOS=linux GOARCH=s390x $(MAKE) --no-print-directory build GO_BUILD_PACKAGES:=./cmd/oc GO_BUILD_FLAGS:="$(GO_BUILD_FLAGS_LINUX_CROSS)" GO_BUILD_BINDIR:=$(CROSS_BUILD_BINDIR)/linux_s390x
.PHONY: cross-build-linux-s390x

cross-build: cross-build-darwin-amd64 cross-build-darwin-arm64 cross-build-windows-amd64 cross-build-linux-amd64 cross-build-linux-arm64 cross-build-linux-ppc64le cross-build-linux-s390x
.PHONY: cross-build

clean-cross-build:
	$(RM) -r '$(CROSS_BUILD_BINDIR)'
	$(RM) cmd/oc/oc.syso
	if [ -d '$(OUTPUT_DIR)' ]; then rmdir --ignore-fail-on-non-empty '$(OUTPUT_DIR)'; fi
.PHONY: clean-cross-build

clean: clean-cross-build
