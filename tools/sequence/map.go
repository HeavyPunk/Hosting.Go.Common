package tools_sequence

func Mapper[TIn any, TOut any](input []TIn, mapper func(TIn) TOut) []TOut {
	res := make([]TOut, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = mapper(input[i])
	}
	return res
}

func MapperIt[TIn any, TOut any](input []TIn, mapper func(int, TIn) TOut) []TOut {
	res := make([]TOut, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = mapper(i, input[i])
	}
	return res
}
