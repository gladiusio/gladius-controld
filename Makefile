##
## Makefile to test and build the gladius binaries
##

##
# GLOBAL VARIABLES
##

# if we are running on a windows machine
# we need to append a .exe to the
# compiled binary
BINARY_SUFFIX=
ifeq ($(OS),Windows_NT)
	BINARY_SUFFIX=.exe
endif

ifeq ($(GOOS),windows)
	BINARY_SUFFIX=.exe
endif

# code source and build directories
SRC_DIR=./cmd
DST_DIR=./build

CLI_SRC=$(SRC_DIR)/gladius-cli
NET_SRC=$(SRC_DIR)/gladius-networkd
# control daemon source is not yet available
CTL_SRC=$(SRC_DIR)/gladius-controld

CLI_DEST=$(DST_DIR)/gladius-cli$(BINARY_SUFFIX)
NET_DEST=$(DST_DIR)/gladius-networkd$(BINARY_SUFFIX)
# control daemon source is not yet available
CTL_DEST=$(DST_DIR)/gladius-controld$(BINARY_SUFFIX)

# commands for go
GOBUILD=go build
GOTEST=go test
##
# MAKE TARGETS
##

# general make targets
all: controld

clean:
	rm -rf ./build/*
	go clean

# dependency management
dependencies:
	# install go packages
	dep ensure

	# Deal with the ethereum cgo bindings
	go get github.com/ethereum/go-ethereum

	cp -r \
	"${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
	"vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"

test: $(NET_SRC)
	$(GOTEST) $(NET_SRC)

controld: test
	$(GOBUILD) -o $(CTR_DEST) $(CTL_SRC)
