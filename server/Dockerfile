FROM golang AS build

WORKDIR /go/src/github.com/oxisto/track-expenses/server

# install dep utility
RUN go get -u github.com/golang/dep/cmd/dep

# copy dependency information and fetch them
COPY Gopkg.* ./
RUN dep ensure --vendor-only

# copy sources
COPY . .

# build and install (without C-support, otherwise there issue because of the musl glibc replacement on Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/server/server.go

FROM alpine:latest
# update CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/track-expenses
COPY --from=build /go/src/github.com/oxisto/track-expenses/server .
CMD ["./server"]
