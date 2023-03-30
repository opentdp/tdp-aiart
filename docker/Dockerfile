FROM alpine AS builder

ADD . /ifs
RUN sh /ifs/deploy

########

FROM alpine

LABEL version="v0.6.1" \
    maintainer="wang@rehiy.com"

COPY --from=builder /usr/local/bin/tdp-* /usr/local/bin/

ENTRYPOINT ["tdp-aiart"]

EXPOSE 7700
