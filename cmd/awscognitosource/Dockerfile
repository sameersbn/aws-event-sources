FROM golang:1.14-alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /go/src/awscognitosource

RUN apk add --no-cache openssh-client ca-certificates

COPY ./ .
RUN go install -installsuffix cgo ./cmd/awscognitosource

FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL name "Triggermesh AWS Cognito Event Source"
LABEL vendor "Triggermesh"
LABEL version "v0.1.0"
LABEL release "1"
LABEL summary "The Triggermesh Cognito Source"
LABEL description "This is the Triggermesh Knative Event Source for AWS Cognito"

COPY licenses/ /licenses/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/awscognitosource /go/bin/awscognitosource

ENTRYPOINT ["/go/bin/awscognitosource"]
