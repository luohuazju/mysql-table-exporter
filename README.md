# MySQL Exporter to monitor One Table Activities

# Command to Build

Fetch libraries

> cd /home/carl/work/mysql-table-exporter/src/sillycat.com/mysql_table_exporter

On MAC OS, install dep if you do not have one

>brew install dep

Add location to the go path if needed

>export GOPATH=/Users/carl/work/go/mysql-table-exporter

> dep ensure -update

Build on CentOS7

> cd /home/carl/work/mysql-table-exporter

> env GOOS=linux GOARCH=amd64 go build -o bin/mysql_table_exporter -v sillycat.com/mysql_table_exporter

Run with ENV

>DATABASE_NAME='xxxx' DB_PASSWORD='password' HTTP_HOST=centos7-master bin/mysql_table_exporter

More ENV and Default Value

```
HTTP_PORT = 18081

HTTP_HOST = localhost

METRICS_PATH = /mysqltable/metrics

METRICS_PREFIX = mysql_table

DB_USERNAME = root

DB_PASSWORD = password

DB_SERVER = localhost

DB_PORT = 3306

DATABASE_NAME = mysql
```

Load dependency
```
go get -u
```
