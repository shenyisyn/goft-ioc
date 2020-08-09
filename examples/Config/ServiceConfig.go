package Config

import (
	"github.com/shenyisyn/goft-ioc/examples/services"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}
