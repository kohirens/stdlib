FROM golang:1.16-alpine3.13 AS dev

ARG USER_NAME
ARG USER_UID
ARG USER_GID
ARG USER_GROUP

ENV GOPATH /home/$USER_NAME
ENV GOFLAGS -mod=vendor
ENV CGO_ENABLED=0
ENV WORK_DIR /home/${USER_NAME}/src/${REPO}

COPY .docker/download-vs-code-server.sh .

# VS Code Requirements: openssh, musl, libgcc, libstdc++
RUN apk --no-progress --purge --no-cache upgrade \
 && apk --no-progress --purge --no-cache add --upgrade \
    curl \
    git \
    gnupg \
    libgcc \
    libstdc++ \
    openssh \
    tini \
 && apk --no-progress --purge --no-cache upgrade \
 && rm -vrf /var/cache/apk/* \
 && curl --version \
 && git --version

# Install vanilla GLibC: https://github.com/sgerrand/alpine-pkg-glibc
RUN curl -o /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
 && curl -LO https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.32-r0/glibc-2.32-r0.apk \
 && apk add glibc-2.32-r0.apk

# Add a non-root group and user, helpful if you dev on Linux.
RUN addgroup --system --gid $USER_GID $USER_GROUP \
 && adduser --system \
    --disabled-password \
    --ingroup $USER_GROUP \
    --uid $USER_UID \
    $USER_NAME

USER $USER_NAME

# Install VS Code Server.
RUN ls -la && ./download-vs-code-server.sh

# Install Go helpful dev tools.
RUN mkdir -p ~/bin \
 && curl -L -o ~/bin/git-chglog https://github.com/git-chglog/git-chglog/releases/download/0.9.1/git-chglog_linux_amd64 \
 && chmod +x ~/bin/git-chglog

RUN go get -u -v golang.org/x/tools/gopls \
    github.com/go-delve/delve/cmd/dlv \
    github.com/ramya-rao-a/go-outline \
    github.com/cweill/gotests/...

# VSCode Requirements for pre-installing extensions
RUN mkdir -p /home/$USER_NAME/.vscode-server/extensions /home/$USER_NAME/.vscode-server-insiders/extensions
#  && chown -R $USER_NAME /home/$USER_NAME/.vscode-server /home/$USER_NAME/.vscode-server-insiders

WORKDIR $WORK_DIR

ENTRYPOINT [ "tini", "--" ]

CMD [ "tail", "-f", "/dev/null" ]

FROM dev AS build

ENV GOPATH /home/$USER

COPY ./ /home/$USER/src
RUN whoami \
 && ls -la .

RUN go build

FROM alpine3.12 AS release

COPY  --chown=$USER_NAME:$USER_GROUP --from=build /home/$USER_NAME/bin /usr/local/bin

ENTRYPOINT [ "go-gitter" ]

CMD [ "-h" ]
