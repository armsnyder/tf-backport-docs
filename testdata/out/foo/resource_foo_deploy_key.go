package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceFooDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceFooDeployKeyCreate,
		Read:   resourceFooDeployKeyRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:        schema.TypeString,
				Description: "The name or id of the project to add the deploy key to.",
				Required:    true,
				ForceNew:    true,
			},
			"title": {
				Type:        schema.TypeString,
				Description: "A title to describe the deploy key with.",
				Required:    true,
				ForceNew:    true,
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
