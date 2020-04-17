package goss

import (
	"fmt"
	"github.com/dghubble/sling"
)

type Plans struct {
	sling *sling.Sling
}

type Plan struct {
	ID     string  `json:"id"`
	Kind   string  `json:"kind"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Cloud  string  `json:"cloud"`
	Region string  `json:"region"`
}

type PlanFindRequest struct {
	Kind   string `url:"kind"`
	Name   string `url:"name"`
	Cloud  string `url:"cloud"`
	Region string `url:"region"`
}

func (s *Plans) Find(planFindRequest *PlanFindRequest) (*Plan, error) {
	plan := new(Plan)
	response, err := s.sling.Get("/v1/plans/new").QueryStruct(planFindRequest).ReceiveSuccess(plan)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return plan, nil
}

func (s *Plans) Get(id string) (*Plan, error) {
	plan := new(Plan)
	response, err := s.sling.Path("/v1/plans/").Get(id).ReceiveSuccess(plan)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return plan, nil
}

func (s *Plans) List() ([]*Plan, error) {
	plan := make([]*Plan, 0)
	response, err := s.sling.Get("/v1/plans/").ReceiveSuccess(&plan)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return plan, nil
}
