FROM docker-hub.artifactory.xyz.net/golang:alpine as build-with-reports
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go test ./test

# go compiler params statically link runtime libs into the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o yaml-linter ./cmd/linter/
FROM scratch
COPY --from=build-with-reports /build/yaml-linter /build/internal/schemas/cicd-schema.json /app/
WORKDIR /app
ENTRYPOINT [ "./yaml-linter", "--data", "/tmp/cicd-metadata.yaml","--schema", "./cicd-schema.json" ]