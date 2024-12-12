package Lem_in

import "fmt"

func PrintAnts(PathAnt [][]string) {
	// for i := 0; i < 1; i++ {
		for j := 0; j <= len(PathAnt[len(PathAnt)-1])-1; j++ {//7
// fmt.Println(len(PathAnt[len(PathAnt)-1])-1)
			for k := 0; k <= len(PathAnt)-1; k++ {
				if j <= len(PathAnt[k])-1 {
					if PathAnt[k][j] == "" {
						continue
					}
					fmt.Print(PathAnt[k][j] + " ")
				}
			}
			fmt.Println()
		}
	// }
}
