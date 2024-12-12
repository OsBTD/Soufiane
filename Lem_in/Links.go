package Lem_in

func LinksFinder(Links [][]string) {
	for _, sl := range Links {
		for _, str := range sl {
			_, ok := Graph[str]
			if ok {
				continue
			}
			for _, sl1 := range Links {
				for i, str2 := range sl1 {
					if str == str2 {
						if i == 0 {
							Graph[str] = append(Graph[str], sl1[i+1])
						} else {
							Graph[str] = append(Graph[str], sl1[i-1])
						}
					}
				}
			}
		}
	}

	// fmt.Println(Graph)
}
