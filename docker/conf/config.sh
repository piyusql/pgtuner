#!/bin/bash

set -e

PG_SETTING_FILE=/var/lib/postgresql/data/postgresql.conf

if [ ${DB_DEBUG} ];
    then
        echo "setting up logging..."
        # enable logging with duration
        sed -ri "s/#log_statement = 'none'/log_statement = 'all'/g" ${PG_SETTING_FILE}
        sed -ri "s/#log_duration = off/log_duration = on/g" ${PG_SETTING_FILE}
        # log statement formatting
        sed -ri "s/#log_line_prefix = .*/log_line_prefix = '%t - %h - %d - %a '/g" ${PG_SETTING_FILE}
    else
        echo "disabled logging..."
fi

# add the following config to bottom of the postgresql.conf
{
    echo "shared_preload_libraries = 'pg_stat_statements'"
    echo "pg_stat_statements.max = 10000"
    echo "pg_stat_statements.track = all"
    echo "#happy postgres.ing ...:)"
} >> ${PG_SETTING_FILE}