package exercicios_ninja_nivel_8

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
)

type user struct {
	First string
	Age   int
}

type user2 struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byFirst []user
type byAge []user

func ResolucaoNaPraticaExercicio1() {
	fmt.Printf("Exemplo 1 \n\n")
	for index, user := range users() {
		userToJson, _ := json.Marshal(user)
		fmt.Printf("User %d: %s\n", index+1, userToJson)
	}

	fmt.Printf("\nExemplo 2 \n\n%v", resolucaoNaPraticaExercicio1())
}

func ResolucaoNaPraticaExercicio2() {
	fromJson := []byte(resolucaoNaPraticaExercicio1())
	var usersFromJson []user2

	err := json.Unmarshal(fromJson, &usersFromJson)
	if err != nil {
		fmt.Println(err)
	}

	for index, user := range usersFromJson {
		fmt.Printf("User %d: %s\n", index+1, fmt.Sprint("First: ", user.First, " Age: ", user.Age))
	}
}

func ResolucaoNaPraticaExercicio3() {
	enc := json.NewEncoder(os.Stdout)

	for _, user := range listUser3() {
		err := enc.Encode(user)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ResolucaoNaPraticaExercicio4() {
	sorted_by_int := sliceInt()

	fmt.Println("Slice de inteiros desordenado: ", sorted_by_int)
	slices.Sort(sorted_by_int)
	fmt.Println("Slice de inteiros ordenado: ", sorted_by_int)

	sorted_by_string := sliceString()

	fmt.Println("Slice de strings desordenado: ", sorted_by_string)
	slices.Sort(sorted_by_string)
	fmt.Println("Slice de strings ordenado: ", sorted_by_string)
}

func ResolucaoNaPraticaExercicio5() {
	fmt.Println("By First")
	fmt.Println("Slice de users desordenado: ", users())
	fmt.Println("Slice de users ordenado por First: ", sortByFirst(users()))

	fmt.Println("\nBy Age")
	fmt.Println("Slice de users desordenado: ", users())
	fmt.Println("Slice de users ordenado por Age: ", sortByAge(users()))
}

func resolucaoNaPraticaExercicio1() string {
	userToJson, _ := json.Marshal(users())

	return string(userToJson)
}

// ResolucaoNaPraticaExercicio1
func users() []user {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	u4 := user{
		First: "Q",
		Age:   64,
	}

	return []user{u3, u4, u2, u1}
}

func listUser3() []user3 {
	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to",
			"take care of that for you, James?",
		},
	}

	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't",
			"have to bring it back",
		},
	}

	u4 := user3{
		First: "Q",
		Last:  "See",
		Age:   64,
		Sayings: []string{
			"Have some of the cake, James",
			"It's all gone, because you were late",
		},
	}

	return []user3{u1, u2, u3, u4}
}

func sliceInt() []int {
	return []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
}

func sliceString() []string {
	return []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
}

func (a byFirst) Len() int           { return len(a) }
func (a byFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func sortByFirst(user []user) []user {
	sort.Sort(byFirst(user))

	return user
}

func sortByAge(user []user) []user {
	sort.Sort(byAge(user))

	return user
}
