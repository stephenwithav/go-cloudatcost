package main

import (
	"fmt"
	"github.com/masayukioguni/go-cloudatcost/cloudatcost"
	"os"
)

func main() {
	Login := os.Getenv("CLOUDATCOST_API_LOGIN")
	Key := os.Getenv("CLOUDATCOST_API_KEY")
	client, _ := cloudatcost.NewClient(&cloudatcost.Option{Login: Login, Key: Key})

	listservers, hr, err := client.ServersService.List()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return
	}

	fmt.Printf("%v,%v\n", listservers, err)

}
