 # Author: Stanislav Yakush <st.yakush@yandex.ru>

TARGET = XScan

PWD = $(shell pwd)

BUILD_DIR = ${PWD}/bin

CMDS = xscan

.PHONY: build clean
.DEFAULT_GOAL := build

build:
	@echo "Build..."
	@$(foreach cmd, $(CMDS), APP_NAME=${cmd} APP_DIR=${PWD}/cmd/${cmd} BUILD_DIR=${BUILD_DIR} ./scripts/build.sh || exit;)
	@echo "Done!"

clean:
	@echo "Clean..."
	@rm -rf ${BUILD_DIR}
	@echo "Done!"
