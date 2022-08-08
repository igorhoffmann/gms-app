# FROM golang:1.18.4-buster

# RUN go version
# ENV GOPATH=/
# COPY ./ ./

# # install psql
# RUN apt-get update
# RUN apt-get -y install postgresql-client

# # make wait-for-postgres.sh executable
# RUN chmod +x wait-for-postgres.sh

# # build go app
# RUN go mod download
# RUN go build -o gms-app ./cmd/main.go

# CMD ["./gms-app"]

FROM golang:1.18.4-buster AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

# build go app
RUN go mod download; CGO_ENABLED=0 go build -o /gms-app ./cmd/main.go


FROM alpine:latest

# copy go app, config and wait-for-postgres.sh
COPY --from=build /gms-app /gms-app
COPY ./configs/ /configs/
COPY ./wait-for-postgres.sh ./

# install psql and make wait-for-postgres.sh executable
RUN apk --no-cache add postgresql-client && chmod +x wait-for-postgres.sh

CMD ["/gms-app"]