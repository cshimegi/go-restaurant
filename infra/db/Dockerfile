FROM mysql:8.0-debian

WORKDIR /

COPY ../../api/db ./db-migrations

RUN apt-get -y upgrade && \
    chmod 777 ./db-migrations/commands.sh
