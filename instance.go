package goss

import (
	"fmt"
	"github.com/dghubble/sling"
	"time"
)

type Instances struct {
	sling *sling.Sling
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
	Name      string   `json:"name"`
	PlanID    string   `json:"plan_id,omitempty"`
	LicenseKey string `json:"license_key,omitempty"`
	EvictionPolicy string `json:"eviction_policy,omitempty"`
	Password  string   `json:"password,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
}

type InstanceUpdateRequest struct {
	ID         string    `json:"-"`
	Name       *string   `json:"name,omitempty"`
	Password   string    `json:"password,omitempty"`
	PlanID     string    `json:"plan_id,omitempty"`
	Whitelist  *[]string `json:"whitelist,omitempty"`
	LicenseKey string    `json:"license_key,omitempty"`
	EvictionPolicy string `json:"eviction_policy,omitempty"`
	Enabled    *bool     `json:"enabled,omitempty"`
}

func (s *Instances) Create(instanceCreqteRequest *InstanceCreateRequest) (*Instance, error) {
	instance := new(Instance)
	response, err := s.sling.Post("/v1/instances/").BodyJSON(instanceCreqteRequest).ReceiveSuccess(instance)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}
	return instance, nil
}

func (s *Instances) Update(instanceUpdateRequest *InstanceUpdateRequest) (*Instance, error) {
	instance := new(Instance)
	response, err := s.sling.Patch("/v1/instances/").Put(instanceUpdateRequest.ID).BodyJSON(instanceUpdateRequest).ReceiveSuccess(instance)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}
	return instance, nil
}

func (s *Instances) Get(id string) (*Instance, error) {
	data := new(Instance)
	response, err := s.sling.Path("/v1/instances/").Get(id).ReceiveSuccess(data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return data, nil
}

func (s *Instances) List() ([]*Instance, error) {
	instance := make([]*Instance, 0)
	response, err := s.sling.Get("/v1/instances/").ReceiveSuccess(&instance)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return instance, nil
}

func (s *Instances) Delete(id string) error {
	response, err := s.sling.Path("/v1/instances/").Delete(id).Receive(nil, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("err: status: %v", response.StatusCode)
	}
	return nil
}
