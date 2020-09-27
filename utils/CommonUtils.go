package utils

import (
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func CanShuTiShi(canShu string) []byte {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	ret, _ := json.Marshal(map[string]interface{}{
		"TiShi": canShu,
	})
	return ret
}
func HuoQuZiFuZhi(canShu interface{}) string {
	if canShu != nil && canShu.(string) != "" {
		return canShu.(string)
	}
	return ""
}
func ShiFouKongZiFu(canShu interface{}) bool {
	if HuoQuZiFuZhi(canShu) == "" {
		return true
	}
	return false
}

func ShiFouChaoChang(canShu interface{}, changDu int) bool {
	if strings.Count(HuoQuZiFuZhi(canShu), "")-1 > changDu {
		return true
	}
	return false
}

// 判断obj是否在target中，target支持的类型string,arrary,slice,map
func ShiFouBaoHan(target interface{}, obj interface{}) bool {
	targetValue := reflect.ValueOf(target)
	kindStr := reflect.TypeOf(target).Kind().String()
	if kindStr == reflect.Slice.String() || kindStr == reflect.Array.String() {
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	}
	if kindStr == reflect.Map.String() {
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	if kindStr == reflect.String.String() {
		if strings.Index(target.(string), obj.(string)) != -1 {
			return true
		}
	}

	return false
}
