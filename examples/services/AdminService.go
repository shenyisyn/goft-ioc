package services

type AdminService struct {
	Order *OrderService `inject:"ServiceConfig.OrderService()"`
}

func NewAdminService() *AdminService {
	return &AdminService{}
}
