package Lem_in

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Ants struct {
	Count int
	Start string
	End   string
	Rooms []string
	Links [][]string
}

var File string

func ReadFile(s string) error {
	var cordinate []string
	var lines []string
	text, err := os.ReadFile(s)
	if err != nil {
		return err
	}

	if len(string(text)) == 0 {
		return fmt.Errorf("%s", "ERROR: invalid data format, empty file !")
	}


	line := strings.Split(string(text), "\n")

	for i := 0; i < len(line); i++ {
		line := strings.TrimSpace(line[i])
		lines = append(lines, line)
	}
	for i := 0; i <= len(lines)-1; i++ {
		if lines[i] == "" || len(lines[i]) > 1 && string(lines[i][0]) == "#" && lines[i] != "##start" && lines[i] != "##end" {
			lines = append(lines[0:i], lines[i+1:]...)
		}
	}

	n, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("%s", "ERROR: invalid data format, invalid number of Ants")
	}

	if n <= 0 {
		return fmt.Errorf("%s", "ERROR: invalid data format, invalid number of Ants")
	}

	Ants.Count = n
	var start, end int
	for i := 0; i <= len(lines)-1; i++ {

		if lines[i] == "##start" {
			start++
		} else if lines[i] == "##end" {
			end++
		}

		if len(lines[i]) != 0 && string(lines[i][0]) == "L" {
			return fmt.Errorf("%s", "ERROR: invalid data format, rooms will never start with the letter L")
		}

		if i > 0 && i < len(lines)-1 && len(strings.Split(lines[i], " ")) < len(strings.Split(lines[i+1], " ")) && string(lines[i+1][0]) != "#" && lines[i] != "##start" && lines[i] != "##end" {
			return fmt.Errorf("%s", "ERROR: invalid data format")
		}

		if i > 0 && len(strings.Split(lines[i], " ")) == 1 && !strings.Contains(lines[i], "-") && lines[i] != "##start" && lines[i] != "##end" {
			return fmt.Errorf("%s", "ERROR: invalid data format")
		}

		if i < len(lines)-1 && len(lines[i+1]) != 0 && (lines[i] == "##start" || lines[i] == "##end") && len(strings.Split(lines[i+1], " ")) != 3 && string(lines[i+1][0]) != "#" {
			return fmt.Errorf("%s", "ERROR: invalid data format")
		}

		cou := 0
		if i != len(lines)-1 && lines[i] == "##start" {
			for j := 0; j <= len(lines)-1; j++ {
				if i+1+j < len(lines)-1 && len(lines[i+1+j]) != 0 && string(lines[i+1+j][0]) == "#" {
					if string(lines[i+1+j][1]) == "#" {
						return fmt.Errorf("%s", "ERROR: invalid data format, no start room !!!")
					}
					continue
				} else if i+1+j < len(lines)-1 && len(lines[i+1+j]) != 0 {
					if len(strings.Split(lines[i+1+j], " ")) == 3 {
						cou++
						Ants.Start = strings.Fields(lines[i+1+j])[0]
						break
					} else {
						return fmt.Errorf("%s", "ERROR: invalid data format, no start room !!!")
					}
				}
			}
			if cou == 0 {
				return fmt.Errorf("%s", "ERROR: invalid data format, no start room !!!")
			}
		}

		if i != len(lines)-1 && lines[i] == "##end" {
			for j := 0; j <= len(lines)-1; j++ {
				if i+1+j < len(lines)-1 && len(lines[i+1+j]) != 0 && string(lines[i+1+j][0]) == "#" {
					if string(lines[i+1+j][1]) == "#" {
						return fmt.Errorf("%s", "ERROR: invalid data format, no end room !!!")
					}
					continue
				} else if i+1+j < len(lines)-1 && len(lines[i+1+j]) != 0 {
					if len(strings.Split(lines[i+1+j], " ")) == 3 {
						Ants.End = strings.Fields(lines[i+1+j])[0]
						cou++
						break
					} else {
						return fmt.Errorf("%s", "ERROR: invalid data format, no end room !!!")
					}
				}
			}
			if cou == 0 {
				return fmt.Errorf("%s", "ERROR: invalid data format, no end room !!!")
			}

		}

		if strings.Contains(lines[i], "-") {
			if len(strings.Split(lines[i], "-")) != 2 {
				return fmt.Errorf("%s", "ERROR: invalid data format, bad links !")
			}

			xx := strings.Split(lines[i], "-")
			if xx[0] == xx[1] {
				return fmt.Errorf("%s", "ERROR: invalid data format, bad links !!!!!!!")
			} else if xx[0] == "" || xx[1] == "" {
				return fmt.Errorf("%s", "ERROR: invalid data format, bad links !!!!!!!")
			}

			Ants.Links = append(Ants.Links, strings.Split(lines[i], "-"))
		} else if strings.Contains(lines[i], " ") {
			k := strings.Split(lines[i], " ")
			y := strings.Join(k[1:], "")
			cordinate = append(cordinate, y)
			Ants.Rooms = append(Ants.Rooms, k[0])

		}

		if len(strings.Split(lines[i], " ")) != 1 && len(strings.Split(lines[i], " ")) != 3 {
			return fmt.Errorf("%s", "ERROR: invalid data format")
		}

	}

	if start != 1 || end != 1 {
		return fmt.Errorf("%s", "ERROR: invalid data format, ##start or ##end repeted or doesn't exist !")
	}

	for i := 0; i <= len(cordinate)-1; i++ {
		for j := 0; j <= len(cordinate)-1; j++ {
			if i != j && cordinate[i] == cordinate[j] {
				return fmt.Errorf("%s", "ERROR: invalid data format, rooms have the same cordinate !")
			}
		}
	}

	links := ""

	for _, t := range Ants.Links {
		L := strings.Join(t, "")
		links += L
	}

	for i, link := range Ants.Links {
		roomchecker := map[string]bool{}
		for _, room := range link {
			roomchecker[room] = true
		}
		for j, link := range Ants.Links {
			c := 0
			for _, room := range link {
				if i != j && roomchecker[room] {
					c++
				}
			}
			if c == 2 {
				return fmt.Errorf("%s", "ERROR: invalid data format, link repeted !!!")
			}

		}
	}

	for _, cor := range cordinate {
		_, err := strconv.Atoi(cor)
		if err != nil {
			return fmt.Errorf("%s", "ERROR: invalid data format, bad cordinate !!!")
		}
	}
	f := strings.Join(Ants.Rooms, "")
	for _, x := range Ants.Links {
		for _, y := range x {
			if !strings.Contains(f, y) {
				return fmt.Errorf("%s", "ERROR: invalid data format, room in the links dosn't exist !")
			}
		}
	}
	for j, y := range Ants.Rooms {
		if !strings.Contains(links, y) {
			return fmt.Errorf("%s", "ERROR: invalid data format, room dosn't have a links !")
		}
		for i := 0; i < len(Ants.Rooms); i++ {
			if i != j && y == Ants.Rooms[i] {
				return fmt.Errorf("%s", "ERROR: invalid data format, room repeted !")
			}
		}
	}
	File = string(text) + "\n"
	return nil
}
