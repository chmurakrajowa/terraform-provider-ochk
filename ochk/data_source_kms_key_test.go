package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type KMSKeyDataSourceTestData struct {
	ResourceName string
	DisplayName  string
	Version      int
}

func (c *KMSKeyDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_kms_key" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  version = "{{ .Version}}"
}
`, c)
}

func (c *KMSKeyDataSourceTestData) FullResourceName() string {
	return "data.ochk_kms_key." + c.ResourceName
}

func TestAccKMSKeyDataSource_read(t *testing.T) {
	kmsKey := &KMSKeyDataSourceTestData{
		ResourceName: "default",
		DisplayName:  testData.KMSKeyDisplayName,
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: kmsKey.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "display_name", kmsKey.DisplayName),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "version", "0"),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "algorithm", "AES"),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "size", "256"),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "state", "Active"),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "key_usage.0", "DECRYPT"),
					resource.TestCheckResourceAttr(kmsKey.FullResourceName(), "key_usage.1", "ENCRYPT"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "sha1_fingerprint"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "sha256_fingerprint"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "object_type"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "default_iv"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "activation_date"),
					resource.TestCheckResourceAttrSet(kmsKey.FullResourceName(), "created_at"),
				),
			},
		},
	})
}
