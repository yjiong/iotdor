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

#RUN mkdir /run/postgresql && chown postgres:postgres -R /run/postgresql &&\
#    su - postgres -c "initdb --username=postgres --encoding=UTF8 --data-checksums -D /var/lib/postgresql/data" &&\
#    su - postgres -c "pg_ctl start -D /var/lib/postgresql/data" &&\
#    su - postgres -c "psql -c \"CREATE USER ${POSTGRES_USER} WITH PASSWORD '${POSTGRES_PASSWORD}';\"" &&\
#    su - postgres -c "psql -c \"CREATE DATABASE ${POSTGRES_DB} OWNER ${POSTGRES_USER}\";" &&\
#    su - postgres -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} TO ${POSTGRES_USER};\"" &&\
#    su - postgres -c "pg_ctl stop -D /var/lib/postgresql/data" 

for dir in ${PGDARADIR} /run/postgresql
do
    if [ ! -d "${dir}" ]; then
        mkdir -p ${dir}
    fi
    chown postgres:postgres -R ${dir}
done

if [ ! -e "${PGDARADIR}/postgresql.conf" ]; then
  echo "[i] ${PGDARADIR} not found, creating initial DBs"
  tfile=`mktemp`
  if [ ! -f "$tfile" ]; then
      return 1
  fi

  cat << EOF > "$tfile"
initdb --username=postgres --encoding=UTF8 --data-checksums -D /var/lib/postgresql/data"
pg_ctl start -D /var/lib/postgresql/data"
psql -c "CREATE USER ${POSTGRES_USER} WITH PASSWORD '${POSTGRES_PASSWORD}';"
psql -c "CREATE DATABASE ${POSTGRES_DB} OWNER ${POSTGRES_USER};"
psql -c "GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} TO ${POSTGRES_USER};"
pg_ctl stop -D /var/lib/postgresql/data" 
EOF

  chmod 755 "$tfile"
  su - postgres -c "$tfile"
  rm -f "$tfile"

fi

su - postgres -c "postgres -D /var/lib/postgresql/data" 
tail -f /dev/null
