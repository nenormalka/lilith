package methods

func ArrayDiff[S ~[]T, T comparable](first, second S) S {
	if len(first) == 0 {
		return second
	}

	if len(second) == 0 {
		return first
	}

	var diff S

	firstMap := ArrayToMapValues[S](first)
	secondMap := ArrayToMapValues[S](second)

	for key := range firstMap {
		if _, ok := secondMap[key]; !ok {
			diff = append(diff, key)
		} else {
			delete(secondMap, key)
		}
	}

	if len(secondMap) == 0 {
		return diff
	}

	for key := range secondMap {
		diff = append(diff, key)
	}

	return diff
}
