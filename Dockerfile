FROM golang:alpine as builder
ENV GO111MODULE=on
ENV PATH=$PATH:$GOPATH/bin

RUN apk add --update nodejs npm alpine-sdk
WORKDIR /build
COPY . .
RUN npx browserslist@latest --update-db
RUN npm install
RUN npm run build

RUN go get github.com/dmarkham/enumer
RUN go install github.com/dmarkham/enumer
RUN go generate ./...
RUN go build -o portfolio

FROM alpine
WORKDIR /portfolio
COPY --from=builder /build/dist ./dist
COPY --from=builder /build/portfolio .
CMD ["./portfolio"]
