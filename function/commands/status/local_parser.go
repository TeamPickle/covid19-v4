package status

import (
	"unicode/utf8"
)

func aliasRegionName(regionName string) string {
	if regionName == "오스트레일리아" {
		return "호주"
	}
	if regionName == "우리나라" || regionName == "한국" {
		return "대한민국"
	}
	return regionName
}

func findRegionName(dict map[string]string, regionName string) string {
	for k, v := range dict {
		if v == regionName && utf8.RuneCountInString(k) == 2 {
			return k
		}
	}
	return ""
}

func isKoreanRegion(regionName string) bool {
	return utf8.RuneCountInString(regionName) == 2 && len(regionName) != 2 && regionName != "기타"
}

func findKoreanRegions(dict map[string]string) (result []string) {
	for k := range dict {
		if isKoreanRegion(k) {
			result = append(result, k)
		}
	}
	return
}
