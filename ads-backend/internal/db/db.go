package db

type DB struct{}

func NewDB(conn Connecter) *DB {
	db, err := conn.Connect()
	if err != nil {
		panic(err)
	}

	return &db
}
