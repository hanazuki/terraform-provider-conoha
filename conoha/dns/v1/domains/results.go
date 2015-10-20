package domains

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type Domain struct {
	ID string `mapstructure:"id"`
	Name string `mapstructure:"name"`
	Ttl int `mapstructure:"ttl"`
	Serial int `mapstructure:"serial"`
	Email string `mapstructure:"email"`
	Description string `mapstructure:"description"`
	GSLB int `mapstructure:"gslb"`
}

type domainResult struct {
	gophercloud.Result
}

func (r domainResult) Extract() (*Domain, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res Domain
	err := mapstructure.Decode(r.Body, &res)
	return &res, err
}

type CreateResult struct {
	domainResult
}

type GetResult struct {
	domainResult
}

type DeleteResult struct {
	gophercloud.ErrResult
}
