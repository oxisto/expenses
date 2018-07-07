FROM node AS build-frontend

# upgrade yarn
RUN npm install -g yarn@1.8

WORKDIR /tmp

ADD frontend/*.json ./
ADD frontend/*.lock ./
RUN yarn install

ADD frontend/. .
RUN yarn run lint
RUN yarn run build --prod

FROM golang AS build-server

WORKDIR /go/src/github.com/oxisto/expenses

# install dep utility
RUN go get -u github.com/golang/dep/cmd/dep

# copy dependency information and fetch them
COPY Gopkg.* ./
RUN dep ensure --vendor-only

# copy sources
COPY . .

# build and install (without C-support, otherwise there issue because of the musl glibc replacement on Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/server/server.go
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/useradd/useradd.go

FROM alpine:latest
# update CA certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/share/expenses
COPY --from=build-frontend /tmp/dist ./frontend/dist
COPY --from=build-server /go/src/github.com/oxisto/expenses/server .
COPY --from=build-server /go/src/github.com/oxisto/expenses/useradd .
CMD ["./server"]
