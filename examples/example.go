package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scalechamp/goss"
)

func main() {
	c := goss.NewClientFromToken("4a2d340ed6b15b5e946f938a8c96d7ad")

	plan, err := c.Plans.Find(context.TODO(), &goss.PlanFindRequest{Cloud: "do", Region: "fra1", Name: "hobby-100", Kind: "redis"})
	if err != nil {
		panic(err)
	}

	p := &goss.InstanceCreateRequest{
		Name:   "minewood-generated",
		PlanID: plan.ID,
	}
	instance, err := c.Instances.Create(context.TODO(), p)
	if err != nil {
		panic(err)
	}

	x, _ := json.MarshalIndent(instance, ">", "  ")
	fmt.Println(string(x))
	fmt.Println(instance.ConnectionInfo.MasterHost)
}
