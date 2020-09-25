package gongjus

import jsoniter "github.com/json-iterator/go"

func CanShuTiShi(canShu string) []byte {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	ret, _ := json.Marshal(map[string]interface{}{
		"TiShi": canShu,
	})
	return ret
}
