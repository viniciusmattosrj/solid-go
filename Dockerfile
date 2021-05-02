FROM go:1.3

LABEL maintainer="Vinicius Mattos vinimattos.rj@gmail.com"

RUN apk update && apk add make

WORKDIR /solid-go

COPY . .

RUN make build

FROM alpine:3.12

COPY --from=builder /solid-go/bin/linux_amd64/solid-go/solid-go /usr/bin/solid-go

ENV PORT=5000

ENTRYPOINT [ "/usr/bin/solid-go" ]

EXPOSE 5000