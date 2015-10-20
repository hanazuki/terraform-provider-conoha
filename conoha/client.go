package conoha

import (
	"github.com/rackspace/gophercloud"
)

func NewDnsV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	eo.ApplyDefaults("dns")
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return nil, err
	}

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint: url,
		ResourceBase: url + "v1/",
	}, nil
}
