-- extension to enable CPU/Memory consumption data

CREATE EXTENSION
IF NOT EXISTS file_fdw;

CREATE SERVER fileserver FOREIGN DATA WRAPPER file_fdw;

CREATE FOREIGN TABLE cpuloadavg
(one text, five text, fifteen text, scheduled text, pid text)
SERVER fileserver
OPTIONS
(filename '/proc/loadavg', format 'text', delimiter ' ');

CREATE FOREIGN TABLE meminfo
(stat text, value text)
SERVER fileserver
OPTIONS
(filename '/proc/meminfo', format 'csv', delimiter ':');

-- Just be sure that the data is available
SELECT *
FROM cpuloadavg;
SELECT *
FROM meminfo;
