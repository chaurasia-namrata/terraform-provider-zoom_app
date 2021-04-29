package zoom

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"terraform-provider-zoom_app/client"
	"terraform-provider-zoom_app/server"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	nameRegex := regexp.MustCompile("^[A-Za-z]\\w{5,29}$")

	if !(nameRegex.MatchString(k)) {
		errs = append(errs, fmt.Errorf("Expected name is not valid .Got %s", value))
		return warns, errs
	}
	return
}
func validateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)

	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid  %s", k))
		return warns, errs
	}
	return

}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{

			"email": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "emailId of new user",
				Required:     true,
				ValidateFunc: validateEmail,
			},
			"first_name": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "first name of new user",
				Required:     true,
				ValidateFunc: validateName,
			},
			"last_name": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "last name of new user",
				Required:     true,
				ValidateFunc: validateName,
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Status of user",
				Optional:    true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	apiClient := m.(*client.Client)

	user := server.Item{
		EmailId:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	err := apiClient.NewItem(&user)

	if err != nil {
		log.Println("[ERROR]: ", err)
		return diag.FromErr(err)
	}
	d.SetId(user.EmailId)
	resourceUserRead(ctx, d, m)
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	apiClient := m.(*client.Client)

	userId := d.Id()
	user, err := apiClient.GetItem(userId)

	if err != nil {
		log.Println("[ERROR]: ", err)
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	}

	if len(user.EmailId) > 0 {
		d.SetId(user.EmailId)
		d.Set("email", user.EmailId)
		d.Set("first_name", user.FirstName)
		d.Set("last_name", user.LastName)
	}

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var _ diag.Diagnostics
	apiClient := m.(*client.Client)
	var diags diag.Diagnostics
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})

		return diags
	}
	user := server.Item{
		EmailId:   d.Get("email").(string),
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
	}

	status := d.Get("status").(string)
	errDeac := apiClient.DeactivateUser(user.EmailId, status)
	log.Println(errDeac)

	err := apiClient.UpdateItem(&user)
	if err != nil {
		log.Printf("[Error] Error updating user :%s", err)
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)

	userId := d.Id()

	err := apiClient.DeleteItem(userId)
	if err != nil {
		log.Printf("[Error] Error deleting user :%s", err)
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
