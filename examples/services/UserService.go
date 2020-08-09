package services

import (
	"fmt"
)

type UserService struct {
	Order *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService) GetUserInfo(uid int) {
	fmt.Println("GetUserInfo")

}
func (this *UserService) GetOrderInfo(uid int) {
	//this.Order.GetOrderInfo(uid)
	fmt.Println("获取用户ID=", uid, "的订单信息")
}
