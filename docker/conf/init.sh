#!/bin/bash

set -e

# execute the initial sqls required
psql -U reader pgtuner < /src/0010-configure.sql
psql -U reader pgtuner < /src/0011-resource-stats.sql
