package acmp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCompany() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCompanyRead,
		Schema: map[string]*schema.Schema{
			"company": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ParentAccountId": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"AccountId": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"AccountState": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"CompanyName": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"VATID": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						//Domain Map missing here
						"BillingStartDate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ContractID": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"Currency": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"Address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"City": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"Country": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"Zip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"Email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCompanyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
  
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/GetCompanies", "https://marketplace.also.de/SimpleAPI/SimpleAPIService.svc/rest"), nil)
	if err != nil {
	  return diag.FromErr(err)
	}
  
	r, err := client.Do(req)
	if err != nil {
	  return diag.FromErr(err)
	}
	defer r.Body.Close()
  
	company := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
	  return diag.FromErr(err)
	}
  
	if err := d.Set("company", company); err != nil {
	  return diag.FromErr(err)
	}
  
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
  
	return diags
  }