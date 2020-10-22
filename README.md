# goft-ioc
基于Golang的轻量级IoC容器。支持单例、多例、运行时获取bean

## 安装
go get -u github.com/shenyisyn/goft-ioc@v0.5.3

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
## 单例注入写法
```go
type UserService struct {
	Order *OrderService `inject:"-"`
}
```
## 多例注入写法
```go
type UserService struct {
	Order *OrderService `inject:"ServiceConfig.OrderService()"`
}
```
多例就是类名+方法名执行，每次都会创建一个新实例
## 使用 初始化
```go
	serviceConfig:=Config.NewServiceConfig()
	Injector.BeanFactory.Config(serviceConfig) //展开方法
	//  BeanFactory.Set()
	{
 
		userService:=services.NewUserService()
		Injector.BeanFactory.Apply(userService) //处理依赖
		fmt.Println(userService.Order)
		 

	}
```
