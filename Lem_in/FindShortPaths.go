package Lem_in

var (
	Graph   = make(map[string][]string)
	Paths   [][]string
	Groupes = [][][]string{}
)

func FindPath(start, end string) [][]string {
	Paths := [][]string{}

	for _, st := range Graph[start] {
		visited := make(map[string]bool)
		Path := BFS(st, end, start, visited)
		if len(Path) == 0 {
			continue
		}
		Paths = append(Paths, Path)
	}
	return Paths
}

func BFS(st, end, start string, visited map[string]bool) []string {
	q := [][]string{{st}}

	visited[start] = true
	visited[st] = true

	for len(q) > 0 {

		Path := q[0]
		q = q[1:]

		node := Path[len(Path)-1]

		if node == end {
			check := map[string]bool{}

			for _, r := range Path {
				check[r] = true
			}

			for _, path := range q {
				for _, p := range path {
					if !check[p] {
						visited[p] = false
					}
				}
			}

			return Path
		}

		for _, ne := range Graph[node] {
			if !visited[ne] {
				if ne != end {
					visited[ne] = true
				}

				newPath := append([]string{}, Path...)
				newPath = append(newPath, ne)
				q = append(q, newPath)
			}
		}

	}

	return nil
}
