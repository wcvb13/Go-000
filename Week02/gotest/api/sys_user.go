package api

import (
	"fmt"
	"gotest/service"
)

func FoundUser(userid int32)  {
	user, err := service.GetUser(userid)
	if err != nil {
		fmt.Printf("404 user not found with err:%+v\n",err)
		return
	}
	fmt.Printf("200 user name:%s age:%d\n",user.Name,user.Age)
}