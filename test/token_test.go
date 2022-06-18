package test

import (
	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"fmt"
	"log"
	"testing"
)

func TestToken(t *testing.T) {
	//uc := define.UserClaim{
	//	Id:       1,
	//	Identity: "USER_1",
	//	Name:     "name",
	//}
	token, err := helper.GenerateToken(1, "USER_1", "name", define.TokenExpire)
	if err != nil {
		log.Println("生成token失败", err)
	}
	fmt.Println(token)
}
