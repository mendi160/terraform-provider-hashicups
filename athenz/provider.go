package athenz

import (
	"fmt"
	"os"

	"git.ouroath.com/athenz/terraform_provider_athenz/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"zms_url": {
				Type:        schema.TypeString,
				Description: fmt.Sprintf("Athenz API URL"),
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ATHENZ_ZMS_URL", nil),
			},
			"cert": {
				Type:        schema.TypeString,
				Description: fmt.Sprintf("Athenz client certificate"),
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ATHENZ_CERT", os.Getenv("HOME")+"/.athenz/cert"),
			},
			"key": {
				Type:        schema.TypeString,
				Description: fmt.Sprintf("Athenz client key"),
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ATHENZ_KEY", os.Getenv("HOME")+"/.athenz/key"),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"athenz_domain": DataSourceDomain(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"athenz_sub_domain": ResourceSubDomain(),
		},

		ConfigureFunc: configProvider,
	}
}

func configProvider(d *schema.ResourceData) (interface{}, error) {
	zms := client.ZmsConfig{
		Url:  d.Get("zms_url").(string),
		Cert: d.Get("cert").(string),
		Key:  d.Get("key").(string),
	}
	if zms.Url == "localhost" {
		return client.AccTestZmsClient()
	}
	return client.NewClient(zms.Url, zms.Cert, zms.Key)
}
