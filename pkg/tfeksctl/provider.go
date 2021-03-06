package tfeksctl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mumoshu/terraform-provider-eksctl/pkg/resource/cluster"
	"github.com/mumoshu/terraform-provider-eksctl/pkg/resource/iamserviceaccount"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"eksctl_cluster":           cluster.Resource(),
			"eksctl_iamserviceaccount": iamserviceaccount.Resource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{}

	return config.Client()
}
