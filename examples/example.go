package main

import (
	"fmt"
	"github.com/scalechamp/goss"
)

func main() {
	c := goss.NewClientFromToken("")

	plan, err := c.Plans.Find(&goss.PlanFindRequest{Cloud: "do", Region: "do-fra-1", Name: "hobby-100", Kind: "redis"})
	if err != nil {
		panic(err)
	}

	instance, err := c.Instances.Create(&goss.InstanceCreateRequest{
		Name:   "sdfsdf",
		PlanID: plan.ID,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(instance.ConnectionInfo.MasterHost)
}
