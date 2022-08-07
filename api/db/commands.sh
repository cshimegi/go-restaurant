#!/bin/bash

up() {
    cat ./db-migrations/migrations/*.up.sql > ./db-migrations/all_up.sql
    FILE=./db-migrations/all_up.sql
    exec_migration up $FILE
    rm $FILE
}

down() {
    cat ./db-migrations/migrations/*.down.sql > ./db-migrations/all_down.sql
    FILE=./db-migrations/all_down.sql
    exec_migration down $FILE
    rm $FILE
}

insert_test_data() {
    echo "Inserting test data"
    cat ./db-migrations/test-data/*.sql > ./db-migrations/all_test_data.sql
    FILE=./db-migrations/all_test_data.sql
    exec_mysql $FILE
    rm $FILE
    echo "Done inserting test data"
}

exec_mysql() {
    mysql -u root --password=${MYSQL_ROOT_PASSWORD} ${MYSQL_DATABASE} < $1
}

exec_migration() {
    echo "Doing migration for $1 from $2"
    exec_mysql $2
    echo "Done migration"
}

