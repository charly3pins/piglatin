FROM golang:1.10

RUN mkdir -p /etc/piglatin

COPY . /etc/piglatin/.

WORKDIR /etc/piglatin/

RUN make deps build

CMD [ "/etc/piglatin/cmd/server/piglatin" ]