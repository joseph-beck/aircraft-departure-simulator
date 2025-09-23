package db

type Connecter interface {
	Connect() (DB, error)
}
