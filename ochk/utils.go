package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func generateRandName() string {
	return fmt.Sprintf("tf-acc-test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}

func generateShortRandName() string {
	return fmt.Sprintf("tf-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}
