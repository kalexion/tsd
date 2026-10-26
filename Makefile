BINARY_NAME := tsd
CMD_PATH := ./cmd/tsd
INSTALL_DIR := $(HOME)/.local/bin

.PHONY: all build install clean

all: build install clean

build:
	go build -o $(BINARY_NAME) $(CMD_PATH)

install: build
	mkdir -p $(INSTALL_DIR)
	install -m 755 $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
