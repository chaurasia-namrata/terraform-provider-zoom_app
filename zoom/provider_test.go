package zoom

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
	"log"
)



var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zoom": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ",err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ZOOM_ADDRESS"); v == "" {
		t.Fatal("ZOOM_ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("ZOOM_TOKEN"); v == "" {
		t.Fatal("ZOOM_TOKEN must be set for acceptance tests")
	}
}
