package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func generateRandName(devTestDataPrefix string) string {
	return generateShortRandName(devTestDataPrefix)
}

func generateShortRandName(devTestDataPrefix string) string {
	return fmt.Sprintf("%s-%s", devTestDataPrefix, acctest.RandStringFromCharSet(4, acctest.CharSetAlphaNum))
}
