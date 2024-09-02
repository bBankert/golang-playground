package arrays

import "fmt"

func PracticeArrays() {
	hobbies := [3]string{"hobby 1", "hobby 2", "hobby 3"}

	//1
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5
	courseGoals := []string{"goal 1", "goal 2"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "goal 3"
	courseGoals = append(courseGoals, "goal 4")
	fmt.Println(courseGoals)
}
