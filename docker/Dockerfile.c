FROM debian:bullseye-slim

RUN apt update \
    && apt install -y build-essential nasm

WORKDIR /emu