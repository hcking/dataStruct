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

func TestSortStudents(t *testing.T) {
	var students Students
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = fmt.Sprintf("stu%d", i)
		stu.Score = rand.Intn(100)
		students = append(students, stu)
	}
	fmt.Println(students)
	fmt.Println("sorting...")
	students.Sort()
	for k, v := range students {
		fmt.Println(k, v)
	}
}
