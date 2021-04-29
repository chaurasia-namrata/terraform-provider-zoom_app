package zoom

import (
   "strings"
   "fmt"
   "terraform-provider-zoom/client"
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceUserRead,
    Schema: map[string]*schema.Schema{
	  "first_name": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
	  "last_name": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
	  "email": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
	  "id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
  }
}




func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
  
	userId := d.Get("id").(string)

	user, err := apiClient.GetItem(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", userId)
		}
	}

	
	d.SetId(user.EmailId)
	d.Set("email", user.EmailId)
	d.Set("first_name", user.FirstName)
	d.Set("last_name", user.LastName)
  
	return nil
  }
  
  
