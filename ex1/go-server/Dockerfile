## Build stage
FROM golang:1.10.3 as build

# Install go dep
# RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/github.com/mpreu/cq-oc-training/ex1/go-server
# COPY Gopkg.toml Gopkg.lock ./
# RUN dep ensure --vendor-only
COPY . ./

# Install dependencies if necessary
#RUN go get github.com/golang/dep/cmd/dep
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure -v -vendor-only
# Build and install
RUN go install -v github.com/mpreu/cq-oc-training/ex1/go-server

## Run stage
FROM golang:1.10.3
EXPOSE 64912
COPY --from=build /go/bin/go-server /bin/go-server

ENTRYPOINT ["/bin/go-server"]