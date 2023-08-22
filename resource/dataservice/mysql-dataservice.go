package dataservice

import (
	"fmt"
	"os"
)

var mysqlResource = ReqLimit{
	ReqMem:   "1500Mi",
	ReqCpu:   "100m",
	LimitMem: "2000Mi",
	LimitCpu: "1000m",
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
	VolumeMountName: "mysql-init",
	MountPath:       "/docker-entrypoint-initdb.d",
	SubPath:         "",
	Data:            map[string]string{
		/*
			"cosmo_api_manager.sql": ``,
			"cosmo_datasource.sql":  ``,
			"cosmo_gateway.sql":     ``,
			"cosmo_proxy.sql":       ``,
			"sdh_auth.sql":          ``,
			"sdh_auth_data.sql":     ``,
		*/
	},
}

func init() {
	// 读取有多少个文件
	files, err := os.ReadDir("./file")
	if err != nil {
		fmt.Println("Error reading directory ./file, err:", err)
		return
	}
	for _, i := range files {
		filename := i.Name()
		data, err := os.ReadFile(fmt.Sprintf("./file/%s", filename))
		if err != nil {
			fmt.Println("读取mysql-dataservice-cm出错:", err)
			return
		}
		mysqlCM2.Data[filename] = string(data)
	}
}
