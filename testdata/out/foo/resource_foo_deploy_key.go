package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceFooDeployKey() *schema.Resource {
	return &schema.Resource{
		Description: "This resource allows you to blah.",

		Create: resourceFooDeployKeyCreate,
		Read:   resourceFooDeployKeyRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The name or id of the project to add the deploy key to.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"title": {
				Description: "A title to describe the deploy key with.",
				Type:        schema.TypeString,
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
