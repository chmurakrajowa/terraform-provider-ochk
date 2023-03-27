package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

type AutoNatTestData struct {
	ResourceName       string
	DisplayName        string
	NetworkDisplayName string
}

func (c *AutoNatTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_virtual_network" "auto-nat" {
  display_name = "{{ .NetworkDisplayName}}"
}

resource "ochk_auto_nat" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  virtual_network_id = data.ochk_virtual_network.auto-nat.id
}
`, c)
}

func (c *AutoNatTestData) FullResourceName() string {
	return "ochk_auto_nat." + c.ResourceName
}

func TestAccAutoNatResource_create(t *testing.T) {
	autoNat := AutoNatTestData{
		DisplayName:        devTestDataPrefix + "-auto-nat-1",
		ResourceName:       "default",
		NetworkDisplayName: testData.Network1Name,
	}

	resourceName := autoNat.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: autoNat.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", autoNat.DisplayName),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(autoNat.DisplayName),
		),
	})
}

type DNatTestData struct {
	ResourceName       string
	DisplayName        string
	NetworkDisplayName string
	Enabled            bool
	SourceNetwork      string
	TranslatedNetwork  string
	DestinationNetwork string
	Priority           string
}

func (c *DNatTestData) ToString() string {
	return executeTemplateToString(
		`
data "ochk_vrf" "project-vrf" {
  display_name = "`+testData.VRF+`"
}

resource "ochk_manual_nat" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  action = "DNAT"
  enabled = {{ .Enabled}}
  vrf_id = data.ochk_vrf.project-vrf.id
  source_network = "{{ .SourceNetwork}}"
  translated_network = "{{ .TranslatedNetwork}}"
  destination_network = "{{ .DestinationNetwork}}"
  priority = {{ .Priority}}
}
`, c)
}

func (c *DNatTestData) FullResourceName() string {
	return "ochk_manual_nat." + c.ResourceName
}

func TestAccDNatResource_create_update(t *testing.T) {
	dNat := DNatTestData{
		DisplayName:        devTestDataPrefix + "-dnat-1",
		ResourceName:       "dnat",
		NetworkDisplayName: testData.Network1Name,
		Enabled:            true,
		SourceNetwork:      "10.10.0.10",
		TranslatedNetwork:  "10.12.12.12",
		DestinationNetwork: testData.NatPublicIpAddr,
		Priority:           "837",
	}

	dNatUpdated := dNat
	dNatUpdated.DisplayName = devTestDataPrefix + "-dnat-1-updated"
	dNatUpdated.Enabled = false
	dNatUpdated.SourceNetwork = "10.20.0.10"
	dNatUpdated.TranslatedNetwork = "10.22.12.12"
	dNatUpdated.DestinationNetwork = testData.NatPublicIpAddr
	dNatUpdated.Priority = "840"

	resourceName := dNat.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: dNat.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", dNat.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "DNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(dNat.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", dNat.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "translated_network", dNat.TranslatedNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", dNat.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", dNat.Priority),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: dNatUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", dNatUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "DNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(dNatUpdated.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", dNatUpdated.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "translated_network", dNatUpdated.TranslatedNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", dNatUpdated.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", dNatUpdated.Priority),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(dNat.DisplayName),
			testAccKMSKeyResourceNotExists(dNatUpdated.DisplayName),
		),
	})

}

type SNatTestData struct {
	ResourceName       string
	DisplayName        string
	NetworkDisplayName string
	Enabled            bool
	SourceNetwork      string
	TranslatedNetwork  string
	DestinationNetwork string
	Priority           string
}

func (c *SNatTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_vrf" "project-vrf" {
  display_name = "`+testData.VRF+`"
}

resource "ochk_manual_nat" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  action = "SNAT"
  enabled = {{ .Enabled}}
  vrf_id = data.ochk_vrf.project-vrf.id
  source_network = "{{ .SourceNetwork}}"
  translated_network = "{{ .TranslatedNetwork}}"
  destination_network = "{{ .DestinationNetwork}}"
  priority = {{ .Priority}}
}
`, c)
}

func (c *SNatTestData) FullResourceName() string {
	return "ochk_manual_nat." + c.ResourceName
}

func TestAccSNatResource_create_update(t *testing.T) {
	sNat := SNatTestData{
		DisplayName:        devTestDataPrefix + "-snat-1",
		ResourceName:       "snat",
		NetworkDisplayName: testData.Network1Name,
		Enabled:            true,
		SourceNetwork:      "10.10.0.10",
		TranslatedNetwork:  testData.NatPublicIpAddr,
		DestinationNetwork: "10.12.12.12",
		Priority:           "838",
	}

	sNatUpdated := sNat
	sNatUpdated.DisplayName = devTestDataPrefix + "-snat-1-updated"
	sNatUpdated.Enabled = false
	sNatUpdated.SourceNetwork = "10.12.0.10"
	sNatUpdated.TranslatedNetwork = testData.NatPublicIpAddr
	sNatUpdated.DestinationNetwork = "10.12.12.14"
	sNatUpdated.Priority = "841"

	resourceName := sNat.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: sNat.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", sNat.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "SNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(sNat.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", sNat.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "translated_network", sNat.TranslatedNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", sNat.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", sNat.Priority),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: sNatUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", sNatUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "SNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(sNatUpdated.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", sNatUpdated.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "translated_network", sNatUpdated.TranslatedNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", sNatUpdated.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", sNatUpdated.Priority),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(sNat.DisplayName),
			testAccKMSKeyResourceNotExists(sNatUpdated.DisplayName),
		),
	})
}

type NoSNatTestData struct {
	ResourceName       string
	DisplayName        string
	NetworkDisplayName string
	Enabled            bool
	SourceNetwork      string
	DestinationNetwork string
	Priority           string
}

func (c *NoSNatTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_vrf" "project-vrf" {
  display_name = "`+testData.VRF+`"
}

resource "ochk_manual_nat" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  action = "NO_SNAT"
  enabled = {{ .Enabled}}
  vrf_id = data.ochk_vrf.project-vrf.id
  source_network = "{{ .SourceNetwork}}"
  destination_network = "{{ .DestinationNetwork}}"
  priority = {{ .Priority}}
}
`, c)
}

func (c *NoSNatTestData) FullResourceName() string {
	return "ochk_manual_nat." + c.ResourceName
}

func TestAccNoSNatResource_create_update(t *testing.T) {
	noSNat := NoSNatTestData{
		DisplayName:        devTestDataPrefix + "-no-snat-1",
		ResourceName:       "no-snat",
		NetworkDisplayName: testData.Network1Name,
		Enabled:            true,
		SourceNetwork:      "10.10.0.10",
		DestinationNetwork: "10.12.12.12",
		Priority:           "843",
	}

	noSNatUpdated := noSNat
	noSNatUpdated.DisplayName = devTestDataPrefix + "-no-snat-1-updated"
	noSNatUpdated.Enabled = false
	noSNatUpdated.SourceNetwork = "10.12.0.10"
	noSNatUpdated.DestinationNetwork = "10.12.12.14"
	noSNatUpdated.Priority = "844"

	resourceName := noSNat.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: noSNat.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", noSNat.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "NO_SNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(noSNat.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", noSNat.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", noSNat.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", noSNat.Priority),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: noSNatUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", noSNatUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "vrf_id", "data.ochk_vrf.project-vrf", "id"),
					resource.TestCheckResourceAttr(resourceName, "action", "NO_SNAT"),
					resource.TestCheckResourceAttr(resourceName, "enabled", strconv.FormatBool(noSNatUpdated.Enabled)),
					resource.TestCheckResourceAttr(resourceName, "source_network", noSNatUpdated.SourceNetwork),
					resource.TestCheckResourceAttr(resourceName, "destination_network", noSNatUpdated.DestinationNetwork),
					resource.TestCheckResourceAttr(resourceName, "priority", noSNatUpdated.Priority),
				),
			},
		},
		CheckDestroy: resource.ComposeTestCheckFunc(
			testAccKMSKeyResourceNotExists(noSNat.DisplayName),
			testAccKMSKeyResourceNotExists(noSNatUpdated.DisplayName),
		),
	})
}
