package foo

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceFooBoring() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceFooBoringRead,
		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceFooBoringRead(_ *schema.ResourceData, _ interface{}) error {
	return nil
}
