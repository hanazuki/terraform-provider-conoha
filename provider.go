package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_PASSWORD", nil),
			},
			"tenant_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_TENANT_NAME", nil),
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_REGION", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"conoha_compute_keypair": resourceComputeKeypair(),
			"conoha_dns_domain":      resourceDnsDomain(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		TenantName: d.Get("tenant_name").(string),
		Region:     d.Get("region").(string),
	}

	return config, nil
}
