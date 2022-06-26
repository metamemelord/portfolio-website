FROM docker.io/library/golang:1-alpine as server-builder
WORKDIR /build
ENV GO111MODULE=on
ENV PATH=$PATH:$GOPATH/bin
RUN go install github.com/dmarkham/enumer@latest
COPY . .
RUN go generate ./...
RUN go build -o portfolio

FROM docker.io/library/node:16-alpine as ui-builder
RUN npm install --location=global npm
WORKDIR /build
COPY . .
RUN npx browserslist@latest --update-db
RUN npm install
RUN npm run build

FROM alpine
WORKDIR /portfolio
COPY --from=ui-builder /build/dist ./dist
COPY --from=server-builder /build/portfolio .
CMD ["./portfolio"]
