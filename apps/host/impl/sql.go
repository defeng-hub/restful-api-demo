package impl

//sql合集

const (
	insertResourceSQL = `INSERT INTO host_resource (
		id,vendor,region,create_at,expire_at,type,
		name,description,status,update_at,sync_at,account,
        public_ip,private_ip
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	insertDescribeSQL = `INSERT INTO host_describe (
		resource_id,cpu,memory,gpu_amount,gpu_spec,
        os_type,os_name,serial_number
	) VALUES (?,?,?,?,?,?,?,?);`

	queryHostSQL = `SELECT * FROM host_resource as r LEFT JOIN host_describe h ON r.id=h.resource_id`
)
