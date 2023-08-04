package utils

const (
	INSERT_SERVICE      = "INSERT INTO ms_service(name, uom, price) VALUES($1, $2, $3)"
	GET_SERVICE_BY_NAME = "SELECT id, name, uom, price FROM ms_service WHERE name = $1"
	GET_SERVICE_BY_ID   = "SELECT id, name, uom, price FROM ms_service WHERE id = $1"
	INSERT_USER         = "INSERT INTO user_credential(id,user_name, password, is_active) VALUES($1, $2, $3, $4)"
	GET_USER_BY_NAME    = "SELECT id, user_name , password, is_active FROM user_credential WHERE user_name = $1"
	GET_USER_BY_ID      = "SELECT id, user_name,  is_active FROM user_credential WHERE id = $1"
	GET_ALL_USER        = "SELECT * FROM user_credential"
)
