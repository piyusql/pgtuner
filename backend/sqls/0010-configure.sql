CREATE EXTENSION
IF NOT EXISTS pg_stat_statements;

CREATE TABLE client
(
    client_id serial primary key,
    hostname varchar(64) not null,
    dbname varchar(32) not null,
    username varchar(32) not null,
    password varchar(16) not null,
    ssl_connect bool not null default 'f'
);

CREATE UNIQUE INDEX idx_client on client(hostname, dbname);

CREATE TYPE chart_type AS ENUM
('INDEX', 'RESOURCE', 'TABLE', 'VACUUM');

CREATE TABLE charts
(
    chart_id serial primary key,
    type chart_type not null,
    name varchar(32) not null unique
);

CREATE TABLE metrics
(
    client_id int references client(client_id) not null,
    chart_id int references charts(chart_id) not null,
    timestamp timestamptz not null,
    data jsonb not null
);
ALTER TABLE metrics ADD primary key (client_id, chart_id, timestamp);

SELECT create_hypertable('metrics', 'timestamp', 'client_id', 4);

INSERT INTO client
    (client_id, hostname, dbname, username, password, ssl_connect)
VALUES
    (1, 'localhost', 'pgtuner', 'reader', 'pass', 'f');

INSERT INTO charts
    (chart_id, type, name)
VALUES
    (1, 'RESOURCE', 'cpu-utilization'),
    (2, 'RESOURCE', 'memory-utilization');
