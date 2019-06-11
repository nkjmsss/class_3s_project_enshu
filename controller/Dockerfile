FROM python:3.6.8-alpine3.9
WORKDIR /src
RUN apk add --no-cache --virtual .builddeps \
      gcc \
      gfortran \
      g++ \
    && \
    pip install numpy && \
    apk del .builddeps
