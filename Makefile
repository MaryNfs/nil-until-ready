.PHONY: new test

new:
	@if [ -z "$(NAME)" ]; then \
		echo 'Usage: make new NAME=valid_palindrome'; \
		exit 1; \
	fi
	./scripts/newlc.sh "$(NAME)" "go/Leetcode"

test:
	go test ./...
