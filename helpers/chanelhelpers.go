package helpers

func SumChanelValues[V int64 | float64 | int32 | int](ch chan V, size int) V {
	var sum V
	for i := 0; i < size; i++ {
		sum += <-ch
	}

	return sum
}
