package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSortHeroSlice(t *testing.T) {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{}
		hero.Name = fmt.Sprintf("hero%d", i)
		hero.Age = rand.Intn(100)
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	fmt.Println("sorting...")
	heroes.Sort()
	for k, v := range heroes {
		fmt.Println(k, v)
	}
}
