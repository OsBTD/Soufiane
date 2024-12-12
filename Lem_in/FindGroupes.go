package Lem_in

func FindGroupes(Paths [][]string) [][][]string {
	Groupe := [][]string{}

	for i := 0; i <= len(Paths)-1; i++ {
		cou := 0
		check := make(map[string]bool)
		for j := 0; j <= len(Paths[i])-2; j++ {
			check[Paths[i][j]] = true
			cou++
		}

		Groupe = append(Groupe, Paths[i])

		for _, st := range Graph[Ants.Start] {
			if !check[st] {
				Path := BFS(st, Ants.End, Ants.Start, check)
				if len(Path) == 0 || cou == 0 {
					continue
				}
				Groupe = append(Groupe, Path)
			}
		}
		Groupes = append(Groupes, Groupe)
		Groupe = [][]string{}
	}

	return Groupes
}

