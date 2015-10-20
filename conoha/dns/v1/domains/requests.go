package domains

import (
	"errors"

	"github.com/rackspace/gophercloud"
)

type CreateOptsBuilder interface {
	ToDomainCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Name string
	Email string
	Ttl int
	Description string
	GSLB int
}

func (opts CreateOpts) ToDomainCreateMap() (map[string]interface{}, error) {
	if opts.Name == "" {
		return nil, errors.New("Missing field required for domain creation: Name")
	}
	if opts.Email == "" {
		return nil, errors.New("Missing field required for domain creation: Email")
	}

	domain := make(map[string]interface{})

	domain["name"] = opts.Name
	domain["email"] = opts.Email
	if opts.Ttl != 0 {
		domain["ttl"] = opts.Ttl
	}
	if opts.Description != "" {
		domain["description"] = opts.Description
	}
	domain["gslb"] = opts.GSLB

	return domain, nil
}

func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToDomainCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return res
}

func Get(client *gophercloud.ServiceClient, name string) GetResult {
	var res GetResult

	_, res.Err = client.Get(getURL(client, name), &res.Body, nil)
	return res
}

func Delete(client *gophercloud.ServiceClient, name string) DeleteResult {
	var res DeleteResult

	_, res.Err = client.Delete(deleteURL(client, name), &gophercloud.RequestOpts{
		OkCodes: []int{200, 202, 204},
	})
	return res
}
