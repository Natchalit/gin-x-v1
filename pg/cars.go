package pg

var CARS = &Connect{
	dbName: `dev_liyl`,
}

// func Cars() (*sqlx.DB, error) {
// 	res, ex := connections.ConnectionSql(`dev_liyl`)
// 	if ex != nil {
// 		return nil, ex
// 	}
// 	return res, nil
// }
