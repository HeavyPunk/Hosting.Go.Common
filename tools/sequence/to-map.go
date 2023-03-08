package tools_sequence

func ToMap[T any, TKey comparable, TValue any](arr []T, keySel func(T) TKey, valSel func(T) TValue) map[TKey]TValue {
	res := make(map[TKey]TValue)
	for i := 0; i < len(arr); i++ {
		res[keySel(arr[i])] = valSel(arr[i])
	}
	return res
}
