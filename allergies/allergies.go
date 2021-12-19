package allergies

// Map Allergen to their respective score
var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(n uint) []string {
	allergies := []string{}
	for v, i := range allergens {
		if n&i != 0 {
			allergies = append(allergies, v)
		}
	}
	return allergies
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&allergens[allergen] != 0
}
