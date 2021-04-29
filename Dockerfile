FROM golang:latest

WORKDIR /sw-system-backend
COPY . /sw-system-backend
RUN go build .

EXPOSE 3000
ENTRYPOINT ["/sw-system-backend/sw-system-backend"]
