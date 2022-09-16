#!/bin/sh
#########################################################################
# File Name: pgup.sh
# Author: yjiong
# mail: 4418229@qq.com
# Created Time: 2022-09-16 01:28:43
#########################################################################
PGDARADIR=/var/lib/postgresql/data
POSTGRES_USER=${POSTGRES_USER:-"iotdor"}
POSTGRES_DB=${POSTGRES_DB:-"iotdor"}
POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-"yaojiong"}

for dir in ${PGDARADIR} /run/postgresql
do
    if [ ! -d "${dir}" ]; then
        mkdir -p ${dir}
    fi
    chown postgres:postgres -R ${dir}
done

if [ ! -e "${PGDARADIR}/postgresql.conf" ]; then
    echo "[i] ${PGDARADIR} not found, creating initial DBs"
    su - postgres -c "initdb --username=postgres --encoding=UTF8 --data-checksums -D /var/lib/postgresql/data"
    su - postgres -c "pg_ctl start -D /var/lib/postgresql/data"
    su - postgres -c "psql -c \"CREATE USER ${POSTGRES_USER} WITH PASSWORD '${POSTGRES_PASSWORD}';\""
    su - postgres -c "psql -c \"CREATE DATABASE ${POSTGRES_DB} OWNER ${POSTGRES_USER}\";"
    su - postgres -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} TO ${POSTGRES_USER};\""
    su - postgres -c "pg_ctl stop -D /var/lib/postgresql/data" 
    sed -i \$a'host all  all    0.0.0.0/0   md5' "${PGDARADIR}/pg_hba.conf"
    sed -i \$a"listen_addresses = '*'" "${PGDARADIR}/postgresql.conf"
fi

su - postgres -c "postgres -D /var/lib/postgresql/data" 
