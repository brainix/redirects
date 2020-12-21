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
	-brew update
	-brew upgrade $(formulae)
	brew cleanup
	-heroku update
	go get -u -v $(packages)


.PHONY: format
format:
	gofmt -s -w .

.PHONY: run
run: format
	heroku local:run go run main.go
