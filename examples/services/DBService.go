package services

type DBService struct {
	DSN string
}

func NewDBService() *DBService {
	return &DBService{DSN: "mysql"}
}
