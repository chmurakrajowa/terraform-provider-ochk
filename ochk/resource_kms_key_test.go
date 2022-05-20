package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type KMSKeyTestData struct {
	ResourceName         string
	DisplayName          string
	KeyUsage             []string
	Algorithm            string
	Size                 int
	Material             string
	PrivateKeyIDToUnwrap string
}

func (c *KMSKeyTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_kms_key" "{{ .ResourceName }}" {
  	display_name = "{{ .DisplayName }}"
	key_usage = {{ StringsToTFList .KeyUsage }}
	algorithm = "{{ .Algorithm }}"
	size = {{ .Size }}
	{{ if .Material }}
	material = "{{ .Material }}"
	{{ end }}
	{{ if .PrivateKeyIDToUnwrap }}
	private_key_id_to_unwrap = "{{ .PrivateKeyIDToUnwrap }}"
	{{ end }}
}
`, c)
}

func (c *KMSKeyTestData) FullResourceName() string {
	return "ochk_kms_key." + c.ResourceName
}

func TestAccKMSKeyResource_create_update(t *testing.T) {
	kmsKey := KMSKeyTestData{
		ResourceName: "key",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		KeyUsage:     []string{"ENCRYPT", "DECRYPT"},
		Algorithm:    "AES",
		Size:         256,
	}

	// there is no update, we test only recreate
	kmsKeyUpdated := kmsKey
	kmsKeyUpdated.DisplayName = kmsKey.DisplayName + "-upd"

	KMSKeyResourceName := kmsKey.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: kmsKey.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(KMSKeyResourceName, "display_name", kmsKey.DisplayName),
					resource.TestCheckResourceAttr(KMSKeyResourceName, "algorithm", kmsKey.Algorithm),
					resource.TestCheckResourceAttr(KMSKeyResourceName, "size", fmt.Sprintf("%d", kmsKey.Size)),
				),
			},
			{
				Config: kmsKeyUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(KMSKeyResourceName, "display_name", kmsKeyUpdated.DisplayName),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(kmsKeyUpdated.DisplayName),
			testAccKMSKeyResourceNotExists(kmsKey.DisplayName),
		),
	})
}

func TestAccKMSKeyResource_import(t *testing.T) {
	kmsKey := KMSKeyTestData{
		ResourceName:         "key",
		DisplayName:          generateShortRandName(devTestDataPrefix),
		KeyUsage:             []string{"ENCRYPT", "DECRYPT"},
		Algorithm:            "RSA",
		Size:                 2048,
		Material:             getTestKMSPrivateKey(),
		PrivateKeyIDToUnwrap: getTestKMSWrapAESKey(),
	}

	KMSKeyResourceName := kmsKey.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: kmsKey.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(KMSKeyResourceName, "display_name", kmsKey.DisplayName),
					resource.TestCheckResourceAttr(KMSKeyResourceName, "algorithm", kmsKey.Algorithm),
					resource.TestCheckResourceAttr(KMSKeyResourceName, "size", fmt.Sprintf("%d", kmsKey.Size)),
				),
			},
			{
				ResourceName:      KMSKeyResourceName,
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(kmsKey.DisplayName),
		),
	})
}

func getTestKMSPrivateKey() string {
	return "-----BEGIN RSA PRIVATE KEY-----\\nMIIEpQIBAAKCAQEA8vXmbi9IAPnkIQTwTijtX3tdNRokMJ7WA2pwMOswEihck5n3\\nf+tbtBTGwWuat2UiPUgWfDtNiqHeK8CTWZuGFyITxF8n2iiQM9CTcR8AJUkp08Fn\\nftNcopnEjYN92wV6ehuPo8k88R9OYA6N5SgxvqFraGCTNE6h+02M+rLeetVVMatw\\n+ZvCz3I8bHFchI9zR+jE3h02z6DvHG9TcaDG2qBn/azWAVY+6sO/6xQgMN1k+jMo\\nPMH6cS62mBdcK+xNq9NhIDWlptnNFuUw97n7a7VVvsTGIS/rjYZuO/ym1s9gXr6B\\n+OoiJmuurxSGer8nYlpYq2VpYlXNAcw1UGMhHQIDAQABAoIBAQCGabYm5S9/osAr\\n6FCN3SSdu2EwfJri7yzVTPBuj97TXNMCsZ50faAJO6lN3psEtQXBQ311E1XtyWlh\\naTPb0ifX6nlnHYGttt04XT8EyTLKbfSe+xOn3YUVS96qr8FUB27f2RmZcj6t4zT3\\n/XVQ/vCuVx1V7H/j41DH9/pzw7tD1mhL0rCjt9NKLXyrFT4ii63q9iLE2GPqfeKV\\nqbG9A6XGhRcumvsIkVxBTh5ofLK2bxLXdHyRk5HEDq0SmKXEeXPULR5vnfZ/lskP\\nHNx2rA3yXS8sEBeigCYtQOngByCWSfXy2R5HVKp2t27PWgWYR9gx58T8WOGOhooC\\nxgWNZmMdAoGBAPsU3VKRzPE2BCrO+hrpYofmZIniHTJGMhVG4mo9WvT8K4eta97Q\\nqE+4I2YrumdFlQ3/ZttRR0KPV0PkmxloOe18oGsp2qQEmUBeVvguFrkwJ7faXacS\\nX/WXLyI+mX6p1GPWx62MalYW9bHRn5Fh1TrTVMDXs84zpcxp90HT7PpLAoGBAPe4\\nT2sPienL3bwl9srRfqo/moQR4xsW7lMIzVgYOZIqhcDf0zPRChW6+uYqn36spDnQ\\nsyZq2LGhCQqhputEbt29jrBsXYTo23FgIojxrrFUcsSnWqWuHhUswAqfSEBVodCr\\n3g2I4jXu11ZIimweFD/nrvHiqdVA+Vv4dRXrEzE3AoGAFaVc10t+kaUIgvBJG5zX\\nQ8QXEtQNlFH861yrFGGpv8klr5LB4/m1KPpFAv/uGA0lpolIQswlCpX0/gNtY6la\\npSDDa7m2AcHrvGLluIuwpdKC+hS3UjoBT9jy1U70SLk3eEwB9vJOEJH2KJhb21rF\\n2UZy3hU6iSJmvtK74E729TUCgYEAtDgr1yjL3gLKQ3qfLHjzHOr+//x/bBLnuhMa\\nSW/+Wl+DRYnQ/s6i9qI8rLzvoln5dHJoE5gCJGCS0mA+rsTvn3Sr3aBI/UvncnlN\\nRrIFtM9KW9WhNg4RprgS0ueEygFCoyyWdORUJoantQc7ZWMQullUxndvtUz63TVK\\nAXMvWEsCgYEAhQ6Ylqpuco/60ybjRzUNhGGsyVMZsi03uyYfjzOsu7UIcdurOZTO\\nA2UhIbAMpBvNLgqqtBwVcwxJkdMW7Z7TGnfqiyh11Pr8yZQcbfMP3FMLgnsTmnrC\\n0yG2J8e9jUhdd70nTWsAQikQaaszrZLdJ/Hk1zm0VIU+ohBaWmg02m4=\\n-----END RSA PRIVATE KEY-----"
}

func getTestKMSWrapAESKey() string {
	return "Er8BJNFpGPIfeBDZuwWELx6R6VKnKHikG81Tz6oDLrQtvrq6/CFJrwkU8gDS1yozKkHOvVroJvbDtUBgb8P4gTj2eTP7p1Ii/NO+fx7ZwpZAbfH9MK9iHqIcr6c6V0UpwYUGJ9ppNC5tPsi3eIc5aDCfOvvBMyPISKGGb04uWS80ff9haYD67vVdBkT5F9HHEny6W1eqUpqPVLPl6AvtVHVlLIfE1I2gCa1/glLalbHPOUIevaOT0ZBofD1rPgt8N3qKj61YT92D0txMn/NvFGzIKk9iCaUl/Z/ZOK6LM3hD2FDC5jFy3vzKEy6l/Rptt9dzywGjvc/1e3k7F/6DVg=="
}

func testAccKMSKeyResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).KMSKeys

		KMSKeys, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(KMSKeys) > 0 {
			return fmt.Errorf("kms_key %s still exists", KMSKeys[0].ID)
		}

		return nil
	}
}
