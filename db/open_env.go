package db

// OpenEnv connects using db.OpenEnv and returns the conn and patch id
// func OpenEnv(env env.Service) (*db.DB, int, error) {
// 	conn, err := db.OpenEnv(env)
// 	if conn == nil {
// 		log.StdOutService(log.LevelInfo).New().Warn("failed to open env")
// 		return conn, -1, err
// 	}
// 	patch, err := Patch(conn)
// 	return conn, patch, err
// }
