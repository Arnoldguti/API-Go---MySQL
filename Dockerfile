# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# File Author / Maintainer
MAINTAINER Arnol Gutierrez


# Copy the local package files to the container's workspace.
ADD . /go/src/posts

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install posts

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/posts


# Document that the service listens on port 8080.
EXPOSE 8080