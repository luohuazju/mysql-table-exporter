# MySQL Exporter to monitor One Table Activities

Set Up go ENV
```
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
tar zxvf go1.21.0.linux-amd64.tar.gz
mv go ~/tool/go-1.21.0
sudo ln -s /home/centos/tool/go-1.21.0 /opt/go-1.21.0
sudo rm -fr /opt/go
sudo ln -s /opt/go-1.21.0 /opt/go

. ~/.bash_profile

go version
go version go1.21.0 linux/amd64
```

Build and Run

```
cd ~/work/mysql-table-exporter/

go get -u
go: downloading github.com/prometheus/client_golang v1.16.0
go: downloading github.com/go-sql-driver/mysql v1.7.1
go: downloading github.com/prometheus/common v0.44.0
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/prometheus/client_model v0.4.0
go: downloading github.com/cespare/xxhash/v2 v2.2.0
go: downloading github.com/prometheus/procfs v0.11.1
go: downloading github.com/cespare/xxhash v1.1.0
go: downloading golang.org/x/sys v0.11.0
go: downloading google.golang.org/protobuf v1.31.0
go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.4
go: downloading github.com/golang/protobuf v1.5.3

go build -o bin/mysql-table-exporter -v ./

ENVIRONMENT=DEV go run main.go

ENVIRONMENT=PROD bin/mysql-table-exporter
```

Run with ENV

```
DATABASE_NAME='xxxx' DB_PASSWORD='password' HTTP_HOST=centos7-master bin/mysql_table_exporter
```

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
