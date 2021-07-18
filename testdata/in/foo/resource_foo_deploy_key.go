package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceFooDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceFooDeployKeyCreate,
		Read:   resourceFooDeployKeyRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFooDeployKeyCreate(_ *schema.ResourceData, _ interface{}) error {
	return nil
}

func resourceFooDeployKeyRead(_ *schema.ResourceData, _ interface{}) error {
	return nil
}
