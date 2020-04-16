package goss

import (
	"fmt"
	"github.com/dghubble/sling"
)

type Plans struct {
	sling *sling.Sling
}

type Plan struct {
	ID     string `json:"id"`
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Cloud  string `json:"cloud"`
	Region string `json:"region"`
}

func (api *Plans) Find(plan *Plan) (*Plan, error) {
	response, err := api.sling.Path("/v1/plans/").Get("new").QueryStruct(plan).ReceiveSuccess(plan)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return plan, nil
}

func (api *Plans) Get(id string) (*Plan, error) {
	data := new(Plan)
	response, err := api.sling.Path("/v1/plans/").Get(id).ReceiveSuccess(data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return data, nil
}

func (api *Plans) List() ([]*Plan, error) {
	data := make([]*Plan, 0)
	response, err := api.sling.Get("/v1/plans/").ReceiveSuccess(&data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return data, nil
}
