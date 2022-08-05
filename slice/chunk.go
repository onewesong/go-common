package slice

func ChunksSlice[T any](l []T, size int) (chunks [][]T) {
	n := len(l)
	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		chunks = append(chunks, l[i:end])
	}
	return
}
