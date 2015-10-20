package main

import (
	"fmt"

	"github.com/hanazuki/terraform-provider-conoha/conoha/dns/v1/domains"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rackspace/gophercloud"
)

func resourceDnsDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsDomainCreate,
		Read:   resourceDnsDomainRead,
		Delete: resourceDnsDomainDelete,
		Exists: resourceDnsDomainExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDnsDomainCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).dnsClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa dns client: %s", err)
	}

	createOpts := domains.CreateOpts{
		Name:  d.Get("name").(string),
		Email: d.Get("email").(string),
	}

	domain, err := domains.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa domain: %s", err)
	}

	d.SetId(domain.ID)

	return resourceDnsDomainRead(d, meta)
}

func resourceDnsDomainRead(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).dnsClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa dns client: %s", err)
	}

	domain, err := domains.Get(client, d.Id()).Extract()
	if err != nil {
		return err
	}

	d.Set("name", domain.Name)
	d.Set("email", domain.Email)

	return nil
}

func resourceDnsDomainDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).dnsClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa dns client: %s", err)
	}

	err = domains.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ConoHa domain: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsDomainExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client, err := meta.(*Config).dnsClient()
	if err != nil {
		return false, fmt.Errorf("Error creating ConoHa dns client: %s", err)
	}

	_, err = domains.Get(client, d.Id()).Extract()
	if err != nil {
		if code, ok := err.(*gophercloud.UnexpectedResponseCodeError); ok && code.Actual == 404 {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
