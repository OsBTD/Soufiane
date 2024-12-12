package Lem_in

import (
	"strconv"
)

var AntsRoom struct {
	LenPath int
	Rooms   []string
}

var Ant []struct {
	LenPath int
	Rooms   []string
}

func PathAnts(Path [][]string, AntCount int) [][]string {
	for i := 0; i <= len(Path)-1; i++ {
		AntsRoom.LenPath = len(Path[i])
		for j := 0; j <= len(Path[i])-1; j++ {
			AntsRoom.Rooms = append(AntsRoom.Rooms, Path[i][j])
		}
		Ant = append(Ant, AntsRoom)
		AntsRoom.Rooms = []string{}
	}

	PathWalk := [][]string{}

	for l := 1; l <= AntCount; l++ {

		cur := 0
		for index, path := range Ant {
			if Ant[cur].LenPath > path.LenPath {
				cur = index
			}
		}

		p := make([]string, Ant[cur].LenPath-len(Ant[cur].Rooms))
		for _, room := range Ant[cur].Rooms {
			path := "L" + strconv.Itoa(l) + "-" + room
			p = append(p, path)
		}

		PathWalk = append(PathWalk, p)

		Ant[cur].LenPath++
	}
	return PathWalk
}
