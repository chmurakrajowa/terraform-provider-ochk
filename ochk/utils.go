package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func generateRandName() string {
	return generateShortRandName()
}

func generateShortRandName() string {
	return fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(4, acctest.CharSetAlphaNum))
}
