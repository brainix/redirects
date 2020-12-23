# --------------------------------------------------------------------------- #
#   Makefile                                                                  #
#                                                                             #
#   Copyright © 2016-2020, Rajiv Bakulesh Shah, original author.              #
#   All rights reserved.                                                      #
# --------------------------------------------------------------------------- #


init upgrade: formulae := {go,heroku}
upgrade: packages := {github.com/gin-gonic/gin,github.com/go-redis/redis}


.PHONY: install
install: init upgrade

.PHONY: init
init:
	-xcode-select --install
	command -v brew >/dev/null 2>&1 || \
		ruby -e "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	brew analytics regenerate-uuid
	brew analytics off
	brew tap heroku/brew
	-brew install $(formulae)

.PHONY: upgrade
upgrade:
	@# export GOPATH="$HOME/Documents/Code/go
	@# export PATH="$PATH:$GOPATH/bin"
	-brew update
	-brew upgrade $(formulae)
	brew cleanup
	-heroku update
	@# -go mod init
	@# go get -u -v $(packages)
	@# go mod tidy


.PHONY: format
format:
	gofmt -s -w .

.PHONY: run
run: format
	go build -o app
	heroku local:run ./app
