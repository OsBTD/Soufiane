package Lem_in

func FindPaths(Paths [][]string) [][]string {
	visited := make(map[string]bool)

	for i := 0; i <= len(Paths[0])-2; i++ {
		visited[Paths[0][i]] = true
	}
	for i := 1; i <= len(Paths)-1; i++ {
		for j := 0; j <= len(Paths[i])-1; j++ {
			if visited[Paths[i][j]] {
				thePath := BFS(Paths[i][0], Ants.End, Ants.Start, visited)
				if len(thePath) == 0 {
					continue
				}
				Paths = append(Paths, thePath)

			}
		}
	}

	// fmt.Println(Paths)
	return Paths
}

// func BFS2(st, end, start string, visited map[string]bool) []string {
// 	q := [][]string{{st}}

// 	visited[start] = true
// 	visited[st] = true

// 	for len(q) > 0 {

// 		Path := q[0]
// 		q = q[1:]

// 		node := Path[len(Path)-1]

// 		if node == end {
// 			return Path
// 		}

// 		for _, ne := range Graph[node] {
// 			if !visited[ne] {
// 				if ne != end {
// 					visited[ne] = true
// 				}

// 				newPath := append([]string{}, Path...)
// 				newPath = append(newPath, ne)
// 				q = append(q, newPath)

// 			}
// 		}

// 	}

// 	return nil
// }
