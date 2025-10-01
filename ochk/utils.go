package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func generateRandName(devTestDataPrefix string) string {
	return generateShortRandName(devTestDataPrefix)
}

func generateShortRandName(devTestDataPrefix string) string {
	return fmt.Sprintf("%s-%s", devTestDataPrefix, acctest.RandStringFromCharSet(4, acctest.CharSetAlphaNum))
}

func NewNullableString(s string) openapi.NullableString {
	var ns openapi.NullableString
	ns.Set(&s)
	return ns
}

func NewNullableInt64(s int64) openapi.NullableInt64 {
	var ns openapi.NullableInt64
	ns.Set(&s)
	return ns
}

func NewNullableBool(s bool) openapi.NullableBool {
	var ns openapi.NullableBool
	ns.Set(&s)
	return ns
}

func GetIntValue(s openapi.NullableInt32) *int32 {
	if s.IsSet() {
		return s.Get()
	}
	return nil
}

func castInt32ToInt(ptr *int32) int {
	if ptr != nil {
		var x = int(*ptr)
		return x
	}
	return 0
}
