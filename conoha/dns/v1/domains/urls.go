package domains

import (
	"github.com/rackspace/gophercloud"
)

const resourcePath = "domains"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func getURL(c *gophercloud.ServiceClient, domainID string) string {
	return c.ServiceURL(resourcePath, domainID)
}

func deleteURL(c *gophercloud.ServiceClient, domainID string) string {
	return c.ServiceURL(resourcePath, domainID)
}
