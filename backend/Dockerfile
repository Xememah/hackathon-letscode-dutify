FROM golang:1.8

WORKDIR /go/src/repo.letscode.sii.pl/wroclaw/three/backend
COPY . .

RUN go-wrapper download 
RUN go-wrapper install

EXPOSE 6000

CMD ["go-wrapper", "run"]