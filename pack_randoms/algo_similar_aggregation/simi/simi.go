package simi

import "fmt"

// Aggr 聚合list[map]数据中的相似元素
// data 需要聚合的数据
// stKeys 判断相似的标准，即 map 的键集合
// result map[int64][]map[string]string 以 int64 为键的map，值为被聚合后的list
// result 中各键对应的 list 中的元素不重复，合起来等同于 data
func Aggr(data []map[string]string, stKeys []string) map[int][]map[string]string {
	simi := map[string]int{}
	result := map[int][]map[string]string{}
	resultKNum := 0
	for _, a := range data {

		var rk = -1
		stMap := map[string]string{}   // ai 在 stKeys 中对应的各值
		stVKMap := map[string]string{} // ai 在 stKeys 中的各值，与stKeys反对应,
		for _, stKey := range stKeys {
			stMap[stKey] = a[stKey]
			stVKMap[a[stKey]] = stKey
		}

		foundMaps := map[string]string{} // stMap 的各值作为 key 在 simi 有的 key
		founds := []string{}

		for _, k := range stKeys {
			if _, v := simi[stMap[k]]; v { // -- 判断 map 存在 key 的方法 （key 有值）
				if _, v := foundMaps[stMap[k]]; v {
					continue
				}
				foundMaps[stMap[k]] = k
				founds = append(founds, stMap[k])
			}
		}

		if len(founds) == 0 { // 如果 founds 为空
			resultKNum++
			rk = resultKNum
			result[rk] = []map[string]string{} // result 新建一个rkey，新增一个空 list
		} else { // founds 有多个
			// 合并 得到合并后的 rkey 给 rk
			for _, fv := range founds {
				if -1 == rk {
					rk = simi[fv]
					continue
				}
				mergingKey := simi[fv]
				mergingList := result[mergingKey]
				delete(result, mergingKey)

				result[rk] = append(result[rk], mergingList...)
				for _, mapEle := range mergingList { // simi 中将 f 作为 value 的 key 全部指向 rk
					for ek, v := range mapEle {
						for _, k := range stKeys {
							if k == ek {
								simi[v] = rk
								break
							}
						}
					}
				}
			}
		}
		// ai 归入 result 中 rk 对应的 rkey 的 list
		result[rk] = append(result[rk], a)
		// simi 中 插入 stValues 的各值 到 rkey 的映射
		for _, v := range stMap {
			simi[v] = rk
		}

		fmt.Println(".")
	}

	fmt.Println("simi", simi)
	return result
}
