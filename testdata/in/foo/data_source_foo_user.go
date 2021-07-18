package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceFooUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFooUserRead,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ConflictsWith: []string{
					"username",
				},
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ConflictsWith: []string{
					"user_id",
				},
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFooUserRead(_ *schema.ResourceData, _ interface{}) error {
	return nil
}
