FROM registry.services.mts.ru/scs-data-platform/devops/images/go-build:v1.1.0 as build
USER root

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=http://nexus.services.mts.ru/repository/go-proxy/ \
    GOSUMDB="sum.golang.org https://nexus.services.mts.ru/repository/go-sum" \
    GOINSECURE=nexus.services.mts.ru

WORKDIR /opt/build

COPY go.mod .
COPY go.sum .
COPY pkg/ ./pkg/
COPY *.go ./

RUN go mod tidy
RUN go build -o appsec.recdep .

FROM registry.services.mts.ru/scs-data-platform/devops/images/base/mts-alpine:v1.0.0-latest as base
USER root

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://proxy.golang.org,direct \
    ANDROID_SDK_ROOT=/usr/lib/android-sdk \
    PATH="/usr/local/go/bin:${PATH}"

WORKDIR /tmp/mts
RUN wget http://pki.mts.ru/root.crt && \
    wget http://pki.mts.ru/class2root.crt && \
    wget http://pki.mts.ru/WinCAG2.crt

WORKDIR /usr/local/share/ca-certificates/mts
RUN cp /tmp/mts/* /usr/local/share/ca-certificates/mts/.

RUN ls /usr/local/share/ca-certificates/mts
RUN cat /usr/local/share/ca-certificates/mts/root.crt >> /etc/ssl/certs/ca-certificates.crt
RUN cat /usr/local/share/ca-certificates/mts/class2root.crt >> /etc/ssl/certs/ca-certificates.crt
RUN cat /usr/local/share/ca-certificates/mts/WinCAG2.crt >> /etc/ssl/certs/ca-certificates.crt

ENV REQUESTS_CA_BUNDLE=/etc/ssl/certs/ca-certificates.crt

RUN apk add py3-pip=23.1.2-r0 sudo=1.9.13_p3-r2 dirmngr wget=1.21.4-r0 android-tools=34.0.0-r6 gradle=8.0.2-r0 npm=9.6.6-r0 git=2.40.1-r0 go=1.20.10-r0
RUN pip install -i https://nexus.services.mts.ru/repository/pypi/simple --trusted-host nexus.services.mts.ru cyclonedx-bom findimports --break-system-packages

ENV http_proxy=""
ENV https_proxy=""

WORKDIR /opt/build
COPY --from=build /opt/build/appsec.recdep .

ENTRYPOINT [ "/opt/build/appsec.recdep" ]