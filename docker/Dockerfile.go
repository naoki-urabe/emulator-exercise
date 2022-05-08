FROM golang:1.18.1-bullseye

RUN apt update \
    && apt install -y nasm

WORKDIR /emu