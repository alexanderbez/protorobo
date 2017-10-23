.PHONE: deps bench

deps:
	@go get -u github.com/cespare/prettybench

bench: deps
	@go test -v ./types --bench=. | prettybench
