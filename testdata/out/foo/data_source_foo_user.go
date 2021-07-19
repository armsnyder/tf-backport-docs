package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceFooUser() *schema.Resource {
	return &schema.Resource{
		Description: "Provides blah for `thing`.\n\n" +
			"Also blah for \"stuff\".",

		Read: dataSourceFooUserRead,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Description: "The ID of the user, for example `foo\\:\\/\\/`.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ConflictsWith: []string{
					"username",
				},
			},
			"username": {
				Description: "The username of the user.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ConflictsWith: []string{
					"user_id",
				},
			},
			"email": {
				Description: "The e-mail address of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dupe_same_desc": {
				Description: "A description.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dupe_diff_desc": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFooUserRead(_ *schema.ResourceData, _ interface{}) error {
	return nil
}
