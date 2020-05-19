FROM timescale/timescaledb:latest-pg11
ADD docker/conf/*.sh /docker-entrypoint-initdb.d/
ADD backend/sqls/*.sql /src/
RUN chmod a+x /docker-entrypoint-initdb.d/init.sh /docker-entrypoint-initdb.d/config.sh
