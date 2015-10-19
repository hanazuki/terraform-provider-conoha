package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
)

func resourceComputeKeypair() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeKeypairCreate,
		Read:   resourceComputeKeypairRead,
		Delete: resourceComputeKeypairDelete,
		Exists: resourceComputeKeypairExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceComputeKeypairCreate(d *schema.ResourceData, meta interface{}) error {
	computeClient, err := meta.(*Config).computeClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa compute client: %s", err)
	}

	createOpts := keypairs.CreateOpts{
		Name:      d.Get("name").(string),
		PublicKey: d.Get("public_key").(string),
	}

	keypair, err := keypairs.Create(computeClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa keypair: %s", err)
	}

	d.SetId(keypair.Name)

	return resourceComputeKeypairRead(d, meta)
}

func resourceComputeKeypairRead(d *schema.ResourceData, meta interface{}) error {
	computeClient, err := meta.(*Config).computeClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa compute client: %s", err)
	}

	keypair, err := keypairs.Get(computeClient, d.Id()).Extract()
	if err != nil {
		return err
	}

	d.Set("name", keypair.Name)
	d.Set("public_key", keypair.PublicKey)

	return nil
}

func resourceComputeKeypairDelete(d *schema.ResourceData, meta interface{}) error {
	computeClient, err := meta.(*Config).computeClient()
	if err != nil {
		return fmt.Errorf("Error creating ConoHa compute client: %s", err)
	}

	err = keypairs.Delete(computeClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ConoHa keypair: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceComputeKeypairExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	computeClient, err := meta.(*Config).computeClient()
	if err != nil {
		return false, fmt.Errorf("Error creating ConoHa compute client: %s", err)
	}

	_, err = keypairs.Get(computeClient, d.Id()).Extract()
	if err != nil {
		if code, ok := err.(*gophercloud.UnexpectedResponseCodeError); ok && code.Actual == 404 {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
