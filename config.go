package main

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type Config struct {
	Username   string
	Password   string
	TenantName string
	Region     string
}

func (c *Config) client() (*gophercloud.ProviderClient, error) {
	options := gophercloud.AuthOptions{
		IdentityEndpoint: fmt.Sprintf("https://identity.%s.conoha.io/v2.0", c.Region),
		Username:         c.Username,
		Password:         c.Password,
		TenantName:       c.TenantName,
	}

	return openstack.AuthenticatedClient(options)
}

func (c *Config) computeClient() (*gophercloud.ServiceClient, error) {
	providerClient, err := c.client()
	if err != nil {
		return nil, err
	}

	return openstack.NewComputeV2(providerClient,
		gophercloud.EndpointOpts{Region: c.Region})
}
