FROM timescale/timescaledb:latest-pg11
ADD docker/conf/config.sh /docker-entrypoint-initdb.d/
RUN chmod a+x /docker-entrypoint-initdb.d/config.sh
ADD backend/sqls/*.sql /src/
RUN chmod a+w -R /src
