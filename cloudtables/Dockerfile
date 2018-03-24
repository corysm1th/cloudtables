# Build Tools
FROM python:2-alpine3.7 AS build
ENV PYTHONUNBUFFERED 1
RUN apk update && \
    apk add --virtual build-deps gcc python-dev musl-dev && \
    apk add postgresql-dev
RUN mkdir /code
WORKDIR /code
ADD requirements.txt /code/
RUN pip install -r requirements.txt

# Production Image
FROM python:2-alpine3.7
COPY --from=build /root/.cache /root/.cache
RUN mkdir /code && \
    mkdir /root/.aws &&\
    apk update && \
    apk add libpq && \
    apk add curl
WORKDIR /code
ADD . /code/
RUN pip install -r requirements.txt
