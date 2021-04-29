package zoom

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.zoom_user.user1", "id", "user_mail@gmail.com"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`	  
	resource "zoom_user" "user1" {
		email        = "user_mail@gmail.com"
		first_name   = "firstname"
		last_name    = "lastname"
	  }
	data "zoom_user" "user1" {
		id = "user_mail@gmail.com"
	}
	`)
}
