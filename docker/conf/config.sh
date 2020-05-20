#!/bin/bash

set -e

PG_SETTING_FILE=/var/lib/postgresql/data/postgresql.conf

if [ ${DB_DEBUG} ];
    then
        echo "setting up logging..."
        # enable logging with duration
        sed -ri "s/#log_min_duration_statement = .*/log_min_duration_statement = 1/g" ${PG_SETTING_FILE}
        # log statement formatting
        sed -ri "s/#log_line_prefix = .*/log_line_prefix = '%t - %h - %d - %u - %a - %p : '/g" ${PG_SETTING_FILE}
    else
        echo "disabled logging..."
fi
sed -ri "s/shared_preload_libraries = .*/shared_preload_libraries = 'timescaledb, pg_stat_statements'/" ${PG_SETTING_FILE}
# add the following config to bottom of the postgresql.conf
{
    # echo "shared_preload_libraries = 'pg_stat_statements'"
    echo "pg_stat_statements.max = 10000"
    echo "pg_stat_statements.track = top"
    echo "#happy postgres.ing ...:)"
} >> ${PG_SETTING_FILE}
