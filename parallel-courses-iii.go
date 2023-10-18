// https://leetcode.com/problems/parallel-courses-iii

type Graph struct {
	Time        int
	Required    []*Graph
	NotLastNode bool
	Index       int32
}

func minimumTime(n int, relations [][]int, time []int) int {
	graphs := make([]Graph, n, n)

	for _, graph := range graphs {
		graph.Required = make([]*Graph, 0)
	}

	for _, relation := range relations {
		prev := relation[0]
		next := relation[1]

		graphs[next-1].Required = append(graphs[next-1].Required, &graphs[prev-1])

		graphs[prev-1].NotLastNode = true
	}

	for key, timeRequired := range time {
		graphs[key].Time = timeRequired
		graphs[key].Index = int32(key)
	}

	lastNodes := make([]*Graph, 0)

	for key := range graphs {
		if !graphs[key].NotLastNode {
			lastNodes = append(lastNodes, &graphs[key])
		}
	}

	return maxFromNodes(lastNodes)
}

var courseTimes = make(map[*Graph]int)

func maxFromNodes(graphs []*Graph) int {
	var result int

	for key, graph := range graphs {
		var courseTime int

		if time, ok := courseTimes[graph]; ok {
			courseTime = time
		}

		if _, ok := courseTimes[graph]; !ok {
			courseTime = graphs[key].Time + maxFromNodes(graphs[key].Required)
			courseTimes[graph] = courseTime
		}

		if courseTime > result {
			result = courseTime
		}
	}
	return result
}
