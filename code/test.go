package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// mapMap()
	// jsonMar()
	arr := LeftJoinString([]string{"a", "b", "c"}, []string{})
	fmt.Println(arr)
}

func mapMap() {
	m := make(map[string]map[string]int)
	b := make(map[string]int)
	b["0-1"] += 1
	b["0-1"] += 1
	b["1-2"] += 1
	m["1234"] = b
	jsonStr, _ := json.Marshal(m)
	fmt.Println(m, string(jsonStr))

}

type test struct {
	dqa *DepartmentQuickAcceptNum
}

type DepartmentQuickAcceptNum struct {
	QuickAcceptNum map[string]map[string]int `json:"quick_accept_num,omitempty"` // 快速接单数 department:startHour_endHour:num
}

func (t DepartmentQuickAcceptNum) Value() (string, error) {
	b, err := json.Marshal(t)
	return string(b), err
}

func (t *DepartmentQuickAcceptNum) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &t)
}
func jsonMar() {
	dqa := &DepartmentQuickAcceptNum{}
	str, _ := dqa.Value()
	fmt.Println(str)
	dqa.Scan(str)
}
func LeftJoinString(lx, rx []string) []string {
	result := make([]string, 0, len(lx))
	rhash := hashSliceString(rx)

	for _, v := range lx {
		_, ok := rhash[v]
		if !ok {
			result = append(result, v)
		}
	}
	return result
}
func hashSliceString(arr []string) map[string]struct{} {
	hash := make(map[string]struct{}, len(arr))
	for _, i := range arr {
		hash[i] = struct{}{}
	}
	return hash
}
