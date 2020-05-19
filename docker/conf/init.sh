#!/bin/bash

set -e

# execute the initial sqls required
psql -U reader pgtuner < /src/configure.sql
psql -U reader pgtuner < /src/resource-stats.sql
