package goss

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type Plans struct {
	client *Client
}

type PlansServiceOp interface {
	Find(ctx context.Context, planFindRequest *PlanFindRequest) (*Plan, error)
	List(ctx context.Context) ([]*Plan, error)
	Get(ctx context.Context, id string) (*Plan, error)
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

func (s *Plans) Find(ctx context.Context, planFindRequest *PlanFindRequest) (*Plan, error) {
	values, err := query.Values(planFindRequest)
	if err != nil {
		return nil, err
	}
	u := url.URL{
		Path:       fmt.Sprintf("%s/new", plansUrl),
		RawQuery:   values.Encode(),
	}
	request, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	plan := new(Plan)
	if err := s.client.Do(ctx, request, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

const plansUrl = "/v1/plans"

func (s *Plans) Get(ctx context.Context, id string) (*Plan, error) {
	path := fmt.Sprintf("%s/%s", plansUrl, id)
	request, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	plan := new(Plan)
	if err := s.client.Do(ctx, request, plan); err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *Plans) List(ctx context.Context) ([]*Plan, error) {
	request, err := s.client.NewRequest(http.MethodGet, plansUrl, nil)
	if err != nil {
		return nil, err
	}
	plans := make([]*Plan, 0)
	if err := s.client.Do(ctx, request, plans); err != nil {
		return nil, err
	}
	return plans, nil
}
