FROM golang:alpine as build-env

# Try hard to get a minimal image to reduce footprint
RUN apk update && \
    apk add go git gcc libc-dev libltdl libtool libgcc && \
    # Install go libraries
	go get github.com/spf13/cobra && \
    go get github.com/spf13/viper && \
    # Install mac2company client
    git clone --depth=1 "https://github.com/mudpuppy12/mac2company.git" "${GOPATH}/src/github.com/mudpuppy12/mac2company" && \
    cd "${GOPATH}/src/github.com/mudpuppy12/mac2company" && \
    go build -o /usr/bin/mac2company ./src/mac2company && \
    cp ./src/config.toml /tmp && \
    apk del go git gcc libc-dev libtool libgcc


FROM alpine
COPY --from=build-env /usr/bin/mac2company /bin/mac2company
COPY --from=build-env /tmp/config.toml /etc/mac2company/config.toml
VOLUME [ "/etc/mac2company" ]
ENTRYPOINT ["/bin/mac2company"]