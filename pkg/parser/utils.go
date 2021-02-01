package parser

import "sort"

func getKeysFromSlice(results []map[string]interface{}) (keys []string) {
	for k := range results[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
