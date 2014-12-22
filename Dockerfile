FROM ubuntu:14.04

ADD https://github.com/sosedoff/hipache-api/releases/download/v0.1.0/hipache-api_linux_amd64 /usr/local/bin/hipache-api
RUN chmod +x /usr/local/bin/hipache-api
CMD ["hipache-api"]