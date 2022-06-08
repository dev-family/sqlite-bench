.PHONY: test
test:
	@go test -bench=. -benchtime=50x
