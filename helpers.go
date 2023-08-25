package set

func findSmallestAndBigestIndex[T comparable](items []*Set[T]) (smallest int, biggest int) {
	if len(items) == 0 {
		return
	}

	smallestSize := len(items[smallest].c)
	biggestSize := len(items[biggest].c)

	for i := range items {
		if length := len(items[i].c); length > biggestSize {
			biggest = i
		} else if length < smallestSize {
			smallest = i
		}
	}

	return
}
