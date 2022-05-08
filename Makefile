build-c:
	docker build -f docker/Dockerfile.c -t emu-c:latest .
run-c:
	docker run -v $(PWD):/emu -it emu-c:latest bash

build-go:
	docker build -f docker/Dockerfile.go -t emu-go:latest .
run-go:
	docker run -v $(PWD):/emu -it emu-go:latest bash