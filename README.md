# goft-ioc
基于Golang的轻量级IoC容器。支持单例、多例、运行时获取bean

## 安装
go get -u github.com/shenyisyn/goft-ioc@v0.5.0

## 使用 -- 创建配置类
```go
type ServiceConfig struct {
}
func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func(this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}  

```
## 使用 初始化
```go
	serviceConfig:=Config.NewServiceConfig()
	Injector.BeanFactory.Config(serviceConfig) //展开方法
	//  BeanFactory.Set()
	{
 
		userService:=services.NewUserService()
		Injector.BeanFactory.Apply(userService) //处理依赖
		fmt.Println(userService.Order.Name())
		userService.GetUserInfo(3)

	}
```
