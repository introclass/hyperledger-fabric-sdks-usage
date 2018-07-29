// Create: 2018/07/16 15:13:00 Change: 2018/07/28 13:27:31
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func main() {

	//读取配置文件，创建SDK
	configProvider := config.FromFile("./config.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("create sdk fail: %s\n", err.Error())
	}

	//读取配置文件(config.yaml)中的组织(member1.example.com)的用户(Admin)
	mspClient, err := mspclient.New(sdk.Context(),
		mspclient.WithOrg("member1.example.com"))
	if err != nil {
		log.Fatalf("create msp client fail: %s\n", err.Error())
	}

	adminIdentity, err := mspClient.GetSigningIdentity("Admin")
	if err != nil {
		log.Fatalf("get admin identify fail: %s\n", err.Error())
	} else {
		fmt.Println("AdminIdentify is found:")
		fmt.Println(adminIdentity)
	}

	//调用合约
	channelProvider := sdk.ChannelContext("mychannel",
		fabsdk.WithUser("Admin"),
		fabsdk.WithOrg("member1.example.com"))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		log.Fatalf("create channel client fail: %s\n", err.Error())
	}

	var args [][]byte
	args = append(args, []byte("key1"))

	request := channel.Request{
		ChaincodeID: "mycc",
		Fcn:         "query",
		Args:        args,
	}
	response, err := channelClient.Query(request)
	if err != nil {
		log.Fatal("query fail: ", err.Error())
	} else {
		fmt.Printf("response is %s\n", response.Payload)
	}
}
