package dataservice

import "fmt"

var mysqlResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var mysqlCMList = CMDataList{
	mysqlCM1,
	mysqlCM2,
}

var mysqlCM1 = CMData{
	Name:            fmt.Sprintf("%s", Mysql),
	VolumeMountName: "mysql-ds-cnf",
	MountPath:       "/etc/mysql/conf.d",
	SubPath:         "",
	Data: map[string]string{
		"my.cnf": `
[mysqld]
datadir=/var/lib/mysql
default-storage-engine=INNODB
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init_connect='SET NAMES utf8mb4'
lower_case_table_names=1
table_open_cache=128
max_connections=5000
max_connect_errors=6000
innodb_file_per_table=1
innodb_buffer_pool_size=1G
max_allowed_packet=64M
transaction_isolation=READ-COMMITTED
innodb_flush_method=O_DIRECT
innodb_lock_wait_timeout=1800
innodb_flush_log_at_trx_commit=0
sync_binlog=0
group_concat_max_len=1024000
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
skip-name-resolve

[client]
default-character-set=utf8mb4
`,
	},
}

var mysqlCM2 = CMData{
	Name:            fmt.Sprintf("%s-init", Mysql.String()),
	VolumeMountName: "mysql-ds-cnf",
	MountPath:       "/etc/mysql/conf.d",
	SubPath:         "",
	Data: map[string]string{
		"my.cnf": ``,
	},
}
