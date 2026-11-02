package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
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

func NewNullableFloat32(s float32) openapi.NullableFloat32 {
	var ns openapi.NullableFloat32
	ns.Set(&s)
	return ns
}

func NewNullableInt64(s int64) openapi.NullableInt64 {
	var ns openapi.NullableInt64
	ns.Set(&s)
	return ns
}

func NewNullableInt32(s int32) openapi.NullableInt32 {
	var ns openapi.NullableInt32
	ns.Set(&s)
	return ns
}

func NewNullableBool(s bool) openapi.NullableBool {
	var ns openapi.NullableBool
	ns.Set(&s)
	return ns
}

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

//func GetIntValue(s openapi.NullableInt32) *int32 {
//	if s.IsSet() {
//		return s.Get()
//	}
//	return nil
//}

func castInt32ToInt(ptr *int32) int {
	if ptr != nil {
		var x = int(*ptr)
		return x
	}
	return 0
}

func NullableStringToUUID(ns openapi.NullableString) (*strfmt.UUID, error) {
	if !ns.IsSet() || ns.Get() == nil {
		return nil, nil
	}

	uuid := strfmt.UUID(*ns.Get())
	return &uuid, nil
}

func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func StringToUUID(ns string) (strfmt.UUID, error) {
	uuid := strfmt.UUID(ns)
	return uuid, nil
}
