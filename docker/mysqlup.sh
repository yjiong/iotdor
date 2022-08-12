#!/bin/sh
#########################################################################
# File Name: mysqlup.sh
# Author: yjiong
# mail: 4418229@qq.com
# Created Time: 2021-01-16 01:28:43
#########################################################################
if [ ! -d "/run/mysqld" ]; then
  mkdir -p /run/mysqld
fi
if [ ! -d /var/lib/mysql/${MYSQL_DATABASE} ]; then
  echo "[i] ${MYSQL_DATABASE} not found, creating initial DBs"
  mysql_install_db --user=root --ldata=/var/lib/mysql
  if [ "$MYSQL_ROOT_PASSWORD" = "" ]; then
    MYSQL_ROOT_PASSWORD=yaojiong
    echo "[i] MySQL root Password: $MYSQL_ROOT_PASSWORD"
  fi

  MYSQL_DATABASE=${MYSQL_DATABASE:-""}
  MYSQL_USER=${MYSQL_USER:-""}
  MYSQL_PASSWORD=${MYSQL_PASSWORD:-""}

  tfile=`mktemp`
  if [ ! -f "$tfile" ]; then
      return 1
  fi


# UPDATE user set plugin='mysql_native_password' where user='root' and host='localhost';
# UPDATE user set authentication_string=password("$MYSQL_ROOT_PASSWORD"),plugin='mysql_native_password' where user='root';
# old version
#FLUSH PRIVILEGES;
#USE mysql;
#GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '$MYSQL_ROOT_PASSWORD' WITH GRANT OPTION;
#GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION;
#ALTER USER 'root'@'localhost' IDENTIFIED BY '$MYSQL_ROOT_PASSWORD';
#UPDATE user set plugin='mysql_native_password' where user='root' and host='localhost';
#FLUSH PRIVILEGES;

#FLUSH PRIVILEGES;
#USE mysql;
#CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'yaojiong'; 
#ALTER USER 'root'@'localhost' IDENTIFIED BY 'yaojiong';
#ALTER USER 'root'@'%' IDENTIFIED BY 'yaojiong';
#GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
#GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION;
#FLUSH PRIVILEGES;
  cat << EOF > $tfile
FLUSH PRIVILEGES;
USE mysql;
CREATE USER IF NOT EXISTS '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';
ALTER USER '$MYSQL_USER'@'localhost' IDENTIFIED BY '$MYSQL_PASSWORD';
ALTER USER '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';
GRANT ALL PRIVILEGES ON *.* TO '$MYSQL_USER'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO '$MYSQL_USER'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF

  if [ "$MYSQL_DATABASE" != "" ]; then
    echo "CREATE DATABASE IF NOT EXISTS $MYSQL_DATABASE DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;" >> $tfile
    echo "[i] Creating database: $MYSQL_DATABASE"
    echo "USE $MYSQL_DATABASE;" >> $tfile
    if [ "$MYSQL_USER" != "" ]; then
      echo "[i] Creating user: $MYSQL_USER with password $MYSQL_PASSWORD"
      echo "GRANT ALL PRIVILEGES ON $MYSQL_DATABASE.* TO '$MYSQL_USER'@'%' WITH GRANT OPTION;" >> $tfile
    fi
  fi

  /usr/bin/mysqld --user=root --bootstrap --verbose=0 < $tfile
  rm -f $tfile

fi
if [ ! -d /sql ];then
    exec /usr/bin/mysqld --user=root --console
else
    /usr/bin/mysqld --user=root --console &
    sleep 3
    /usr/bin/mysql -uroot -proot $MYSQL_DATABASE -e "SOURCE /sql/iotd.sql;"
    rm -rf /sql
    tail -f /dev/null
fi
