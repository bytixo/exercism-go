package raindrops

import "strconv"

type Rain struct {
	Number int
	Sound  string
}

func Convert(number int) string {

	value := []Rain{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

	var data string

	for _, v := range value {
		if number%v.Number == 0 {
			data += v.Sound
		}
	}
	if data == "" {
		data += strconv.Itoa(number)
	}
	return data
}
