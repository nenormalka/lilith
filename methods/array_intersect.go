package methods

func ArrayIntersect[S ~[]T, T comparable](first, second S) S {
	var intersect S

	if len(first) == 0 || len(second) == 0 {
		return intersect
	}

	firstMap := ArrayToMapValues[S](first)
	secondMap := ArrayToMapValues[S](second)

	for key := range firstMap {
		if _, ok := secondMap[key]; !ok {
			continue
		}

		intersect = append(intersect, key)
	}

	return intersect
}
