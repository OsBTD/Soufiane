package main

import (
	"fmt"
	"os"
	"sort"

	Lem_in "lem-in/Lem_in"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid args !")
		return
	}

	err := Lem_in.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	Lem_in.LinksFinder(Lem_in.Ants.Links)
	paths := Lem_in.FindPath(Lem_in.Ants.Start, Lem_in.Ants.End)

	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format, no path found !!!")
		return
	}

	Groupes := Lem_in.FindGroupes(paths)

	var sl []int

	TheBest := make(map[int][][][]string)

	for _, ThePath := range Groupes {
		pathAnt := Lem_in.PathAnts(ThePath, Lem_in.Ants.Count)
		TheBest[len(pathAnt[len(pathAnt)-1])] = append(TheBest[len(pathAnt[len(pathAnt)-1])], ThePath)
		sl = append(sl, len(pathAnt[len(pathAnt)-1]))
	}
	sort.Ints(sl)

	FinalPath := Lem_in.PathAnts(TheBest[sl[0]][0], Lem_in.Ants.Count)

	fmt.Println(Lem_in.File)

	Lem_in.PrintAnts(FinalPath)
}
