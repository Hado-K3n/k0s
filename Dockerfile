ARG ARCH
FROM ${ARCH}alpine:3.15
ARG TARGETARCH

ADD docker-entrypoint.sh /entrypoint.sh
COPY ./k0s-${TARGETARCH}/k0s /usr/local/bin/k0s

RUN apk add --no-cache bash coreutils findutils curl tini

ENV KUBECONFIG=/var/lib/k0s/pki/admin.conf


ENTRYPOINT ["/sbin/tini", "--", "/bin/sh", "/entrypoint.sh" ]


CMD ["k0s", "controller", "--enable-worker"]
