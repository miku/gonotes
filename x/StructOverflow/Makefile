main: main.go
	time go build -ldflags="-w -s" main.go

main.go: gen.py
	python gen.py > main.go

