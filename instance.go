package goss

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Instances struct {
	client *Client
}

type InstancesServiceOp interface {
	Create(ctx context.Context, instanceCreateRequest *InstanceCreateRequest) (*Instance, error)
	Update(ctx context.Context, instanceUpdateRequest *InstanceUpdateRequest) (*Instance, error)
	Get(ctx context.Context, id string) (*Instance, error)
	List(ctx context.Context) ([]*Instance, error)
	Delete(ctx context.Context, id string) error
}

type Instance struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Kind           string         `json:"kind"`
	Password       string         `json:"password"`
	State          string         `json:"state"`
	Enabled        bool           `json:"enabled"`
	Whitelist      []string       `json:"whitelist"`
	PlanID         string         `json:"plan_id"`
	LicenseKey     *string        `json:"license_key,omitempty"`
	EvictionPolicy *string        `json:"eviction_policy,omitempty"`
	ConnectionInfo ConnectionInfo `json:"connection_info"`
	CreatedAt      time.Time      `json:"created_at"`
}

type ConnectionInfo struct {
	MasterHost  string `json:"master_host"`
	ReplicaHost string `json:"replica_host"`
}

type InstanceCreateRequest struct {
	Name           string   `json:"name,omitempty"`
	Password       string   `json:"password,omitempty"`
	PlanID         string   `json:"plan_id,omitempty"`
	Whitelist      []string `json:"whitelist,omitempty"`
	LicenseKey     string   `json:"license_key,omitempty"`     // only for keydb-pro
	EvictionPolicy string   `json:"eviction_policy,omitempty"` // only for keydb-pro, redis, keydb
}

type InstanceUpdateRequest struct {
	ID             string   `json:"-"`
	Name           string   `json:"name,omitempty"`
	Password       string   `json:"password,omitempty"`
	PlanID         string   `json:"plan_id,omitempty"`
	Whitelist      []string `json:"whitelist,omitempty"`
	Enabled        *bool    `json:"enabled,omitempty"`
	LicenseKey     string   `json:"license_key,omitempty"`     // only for keydb-pro
	EvictionPolicy string   `json:"eviction_policy,omitempty"` // only for keydb-pro, redis, keydb
}

func (s *Instances) Create(ctx context.Context, instanceCreateRequest *InstanceCreateRequest) (*Instance, error) {
	request, err := s.client.NewRequest(http.MethodPost, instances, instanceCreateRequest)
	if err != nil {
		return nil, err
	}
	instance := new(Instance)
	if err := s.client.Do(ctx, request, instance); err != nil {
		return nil, err
	}
	return instance, nil
}

func (s *Instances) Update(ctx context.Context, instanceUpdateRequest *InstanceUpdateRequest) (*Instance, error) {
	path := fmt.Sprintf("%s/%s", instances, instanceUpdateRequest.ID)
	request, err := s.client.NewRequest(http.MethodPatch, path, instanceUpdateRequest)
	if err != nil {
		return nil, err
	}
	instance := new(Instance)
	if err := s.client.Do(ctx, request, instance); err != nil {
		return nil, err
	}
	return instance, nil
}

const instances = "/v1/instances"

func (s *Instances) Get(ctx context.Context, id string) (*Instance, error) {
	path := fmt.Sprintf("%s/%s", instances, id)
	request, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	instance := new(Instance)
	if err := s.client.Do(ctx, request, instance); err != nil {
		return nil, err
	}
	return instance, nil
}

func (s *Instances) List(ctx context.Context) ([]*Instance, error) {
	request, err := s.client.NewRequest(http.MethodGet, instances, nil)
	if err != nil {
		return nil, err
	}
	instances := make([]*Instance, 0)
	if err := s.client.Do(ctx, request, instances); err != nil {
		return nil, err
	}
	return instances, nil
}

func (s *Instances) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", instances, id)
	request, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return s.client.Do(ctx, request, instances)
}
