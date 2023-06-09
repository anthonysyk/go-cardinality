test:
	go test -v ./...

bench:
	go test -run=^# -bench=. -count 5

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
