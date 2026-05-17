FROM alpine:3.20 AS builder

RUN apk add --no-cache ca-certificates curl

ARG ARCH=linux/amd64
ARG KUBECTL_VER=1.31.0

RUN curl -sL https://dl.k8s.io/release/v${KUBECTL_VER}/bin/${ARCH}/kubectl \
    -o /usr/local/bin/kubectl && chmod +x /usr/local/bin/kubectl

RUN kubectl version --client=true

FROM alpine:3.20 AS kuberik-cli

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/kubectl /usr/local/bin/
COPY --chmod=755 kuberik /usr/local/bin/

USER 65534:65534
ENTRYPOINT ["kuberik"]
