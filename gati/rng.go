package gati

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
	sgc7http "github.com/zhs007/slotsgamecore7/http"
	sgc7utils "github.com/zhs007/slotsgamecore7/utils"
	"go.uber.org/zap"
)

// RngInfo - rng infomation
type RngInfo struct {
	Bits  int `json:"bits"`
	Range int `json:"range"`
	Value int `json:"value"`
}

// GetRngs - get rngs
func GetRngs(rngURL string, gameID int, nums int) ([]int, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	url := sgc7utils.AppendString(rngURL, "?size=", strconv.Itoa(nums))

	mapHeader := make(map[string]string)
	mapHeader["X-Game-ID"] = strconv.Itoa(gameID)
	code, body, err := sgc7http.HTTPGet(url, mapHeader)
	if err != nil {
		sgc7utils.Error("gati.GetRngs:HTTPGet",
			zap.Error(err),
			zap.String("url", url))

		return nil, err
	}

	if code != 200 {
		sgc7utils.Error("gati.GetRngs:HTTPGet",
			zap.Int("code", code),
			zap.String("url", url))

		return nil, err
	}

	// client := &http.Client{}

	// req, err := http.NewRequest("GET", sgc7utils.AppendString(rngURL, "?size=", strconv.Itoa(nums)), nil)
	// if err != nil {
	// 	return nil, err
	// }

	// req.Header.Add("X-Game-ID", strconv.Itoa(gameID))

	// resp, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	lst := []int{}
	err = json.Unmarshal(body, &lst)
	if err != nil {
		return nil, err
	}

	return lst, nil
}
