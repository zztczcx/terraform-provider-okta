package okta

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

var emailTemplatesDataSourceSchema = map[string]*schema.Schema{
	"brand_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Brand ID",
	},
	"email_templates": {
		Type:        schema.TypeSet,
		Computed:    true,
		Description: "List of `okta_email_template` belonging to a brand in the organization",
		Elem: &schema.Resource{
			Schema: emailTemplateDataSourceSchema,
		},
	},
}

var emailTemplateDataSourceSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The name of the email template",
	},
	"links": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Link relations for this object - JSON HAL - Discoverable resources related to the email template",
	},
}

func flattenEmailTemplate(emailTemplate *okta.EmailTemplate) map[string]interface{} {
	attrs := map[string]interface{}{}
	attrs["name"] = emailTemplate.Name
	links, _ := json.Marshal(emailTemplate.Links)
	attrs["links"] = string(links)

	return attrs
}
