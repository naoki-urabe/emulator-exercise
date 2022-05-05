build:
	docker build -t emu:latest .
run:
	docker run -v $(PWD):/emu -it emu:latest bash