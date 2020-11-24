package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
)

func testStringInSet(resource, setKeyPrefix, expectedValue string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("resource key not found: %s", setKeyPrefix)
		}

		setLengthKey := fmt.Sprintf("%s.#", setKeyPrefix)
		setLengthStr, ok := rs.Primary.Attributes[setLengthKey]
		if !ok {
			return fmt.Errorf("attribute %s not found", setLengthKey)
		}

		setLength, _ := strconv.Atoi(setLengthStr)
		for i := 0; i < setLength; i++ {
			if value := rs.Primary.Attributes[fmt.Sprintf("%s.%d", setKeyPrefix, i)]; value == expectedValue {
				return nil
			}
		}

		return fmt.Errorf("value %s not found in set %s", expectedValue, setKeyPrefix)
	}
}
