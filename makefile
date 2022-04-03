simulate:
	go run main.go $(ARGS)

simulate-example:
	go run main.go ./resources/example-map.txt -i 10000 -a 100 -v

test:
	go test -v -cover ./...
