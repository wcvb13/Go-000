package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "projecttest/api/app/user/v1"
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect fail with err:%s\n",err.Error())
	}
	defer conn.Close()

	client := v1.NewUserServiceClient(conn)
	req := new(v1.UserRequest)

	req.Id = 5

	resp,err1 := client.GetUserInfo(context.Background(),req)
	if err1 != nil {
		fmt.Printf("response err:%s\n",err1.Error())
	}
	fmt.Println(resp)
}