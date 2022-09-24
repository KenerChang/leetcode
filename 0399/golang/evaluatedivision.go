package evaluatedivision

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// we can use a map to store maps of values for each variable
	// If two variables are connected, and they belong to different map
	// then we need merge them into one map, and update values of variables which are merged in to new map

	maps := map[string]map[string]float64{}
	for i, equation := range equations {
		l, r := equation[0], equation[1]

		lmap, rmap := maps[l], maps[r]
		if lmap == nil && rmap == nil {
			// create a new map
			m := map[string]float64{
				l: values[i],
				r: 1.0,
			}

			maps[l] = m
			maps[r] = m
		} else if lmap == nil {
			rmap[l] = rmap[r] * values[i]
			maps[l] = rmap
		} else if rmap == nil {
			lmap[r] = lmap[l] / values[i]
			maps[r] = lmap
		} else {
			// merge two maps
			for k, v := range lmap {
				rmap[k] = rmap[r] * values[i] * (v / lmap[l])
				maps[k] = rmap
			}
		}
	}

	results := []float64{}
	for _, q := range queries {
		l, r := q[0], q[1]
		lmap, rmap := maps[l], maps[r]
		if lmap == nil || rmap == nil {
			results = append(results, -1.0)
			continue
		}

		if lmap[l] != rmap[l] || lmap[r] != rmap[r] {
			results = append(results, -1.0)
			continue
		}

		results = append(results, lmap[l]/lmap[r])
	}
	return results
}
