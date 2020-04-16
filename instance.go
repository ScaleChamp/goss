package goss

import (
	"fmt"
	"github.com/dghubble/sling"
	"time"
)

//func (api *API) waitUntilReady(id string) (map[string]interface{}, error) {
//	data := make(map[string]interface{})
//	failed := make(map[string]interface{})
//	for {
//		response, err := api.sling.Path("/v1/instances/").Get(id).Receive(&data, &failed)
//		if err != nil {
//			return nil, err
//		}
//		if response.StatusCode != 200 {
//			return nil, errors.New(fmt.Sprintf("waitUntilReady failed, status: %v, message: %s", response.StatusCode, failed))
//		}
//		if data["state"] == "running" {
//			return data, nil
//		}
//
//		time.Sleep(10 * time.Second)
//	}
//}

type Instances struct {
	sling *sling.Sling
}

type Instance struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	State     string    `json:"state,omitempty"`
	Enabled   bool      `json:"enabled,omitempty"`
	Whitelist []string  `json:"whitelist,omitempty"`
	PlanID    string    `json:"plan_id,omitempty"`

	ConnectionInfo *struct {
		MasterHost  string `json:"master_host"`
		ReplicaHost string `json:"replica_host"`
	} `json:"connection_info,omitempty"` // read-only
}

func (api *Instances) Create(params *Instance) (*Instance, error) {
	response, err := api.sling.Post("/v1/instances/").BodyJSON(params).ReceiveSuccess(params)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}
	return params, nil
}

func (api *Instances) Get(id string) (*Instance, error) {
	data := new(Instance)
	response, err := api.sling.Path("/v1/instances/").Get(id).ReceiveSuccess(data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return data, nil
}

func (api *Instances) List() ([]*Instance, error) {
	data := make([]*Instance, 0)
	response, err := api.sling.Get("/v1/instances/").ReceiveSuccess(&data)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: status: %v", response.StatusCode)
	}

	return data, nil
}

func (api *Instances) Delete(id string) error {
	response, err := api.sling.Path("/v1/instances/").Delete(id).ReceiveSuccess(nil)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("err: status: %v", response.StatusCode)
	}
	return nil
}
