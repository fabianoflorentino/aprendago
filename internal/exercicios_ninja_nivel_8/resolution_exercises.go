// Package exercicios_ninja_nivel_8 contains various exercises demonstrating
// practical usage of Go language features such as JSON encoding/decoding,
// sorting slices, and custom sorting implementations. The exercises include
// examples of working with structs, marshaling and unmarshaling JSON data,
// encoding JSON to stdout, and sorting slices of integers, strings, and custom
// structs by different fields.
package exercicios_ninja_nivel_8

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
)

// user represents a person with a first name and age.
type user struct {
	First string
	Age   int
}

// user2 represents a user with a first name and age. The struct fields are
// annotated with JSON tags to specify how they should be serialized and
// deserialized.
type user2 struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

// user3 represents a user with a first name, last name, age, and a list of sayings.
type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// byFirst is a custom type that implements the sort.Interface for sorting a slice of user structs by their first name.
type byFirst []user

// byAge is a custom type that implements the sort.Interface for sorting a slice of user structs by their age.
type byAge []user

// ResolucaoNaPraticaExercicio1 demonstrates two examples of handling user data.
// In the first example, it iterates over a list of users, converts each user to JSON format,
// and prints the JSON string along with the user index.
// In the second example, it prints the result of the function resolucaoNaPraticaExercicio1().
//
// Package internal/exercicios_ninja_nivel_8 contains exercises and solutions for advanced Go programming concepts.
func ResolucaoNaPraticaExercicio1() {
	fmt.Printf("Exemplo 1 \n\n")

	for index, user := range users() {
		userToJson, _ := json.Marshal(user)
		fmt.Printf("User %d: %s\n", index+1, userToJson)
	}

	fmt.Printf("\nExemplo 2 \n\n%v", resolucaoNaPraticaExercicio1())
}

// ResolucaoNaPraticaExercicio2 unmarshals a JSON byte slice into a slice of user2 structs,
// and then iterates over the slice to print each user's first name and age.
// It relies on the resolucaoNaPraticaExercicio1 function to provide the JSON data.
// If an error occurs during unmarshalling, it prints the error.
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

// ResolucaoNaPraticaExercicio3 encodes a list of users to JSON and writes it to the standard output.
// It retrieves the list of users from the listUser3 function and uses the json.Encoder to encode each user.
// If an error occurs during encoding, it prints the error to the standard output.
func ResolucaoNaPraticaExercicio3() {
	enc := json.NewEncoder(os.Stdout)

	for _, user := range listUser3() {
		err := enc.Encode(user)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// ResolucaoNaPraticaExercicio4 demonstrates sorting slices of integers and strings.
// It first creates an unsorted slice of integers and prints it, then sorts the slice
// and prints the sorted result. The same process is repeated for a slice of strings.
// This function relies on the slices package for sorting operations.
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

// ResolucaoNaPraticaExercicio5 demonstrates sorting a slice of users by their first name and age.
// It prints the unsorted and sorted slices to the console for both sorting criteria.
func ResolucaoNaPraticaExercicio5() {
	fmt.Println("By First")
	fmt.Println("Slice de users desordenado: ", users())
	fmt.Println("Slice de users ordenado por First: ", sortByFirst(users()))

	fmt.Println("\nBy Age")
	fmt.Println("Slice de users desordenado: ", users())
	fmt.Println("Slice de users ordenado por Age: ", sortByAge(users()))
}

// resolucaoNaPraticaExercicio1 converts a list of users to a JSON string.
// It marshals the user data into JSON format and returns it as a string.
func resolucaoNaPraticaExercicio1() string {
	userToJson, _ := json.Marshal(users())

	return string(userToJson)
}

// Sorted by First from users struct
// Len returns the number of elements in the collection.
// It is a method of the byFirst type, which is used to implement
// the sort.Interface for sorting a collection based on the first element.
func (a byFirst) Len() int { return len(a) }

// Swap swaps the elements with indexes i and j.
// This method is part of the byFirst type, which is used to sort a collection based on the first element.
func (a byFirst) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares the First field of two elements in the byFirst slice
// and returns true if the First field of the element at index i is less
// than the First field of the element at index j.
func (a byFirst) Less(i, j int) bool { return a[i].First < a[j].First }

// Sorted by Age from users struct
// Len returns the number of elements in the byAge slice.
// It implements the Len method of the sort.Interface.
func (a byAge) Len() int { return len(a) }

// Swap swaps the elements with indexes i and j in the byAge slice.
// This method is part of the sort.Interface implementation for sorting by age.
func (a byAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less reports whether the element with index i should sort before the element with index j.
// It compares the Age field of the elements in the byAge slice to determine the order.
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// users returns a slice of user structs containing predefined user data.
// The function creates four user instances with different names and ages,
// and returns them in a slice in a specific order.
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

// listUser3 creates and returns a slice of user3 structs, each representing a different user
// with their respective first name, last name, age, and a list of sayings.
// The function initializes four user3 instances with predefined data and returns them as a slice.
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

// sliceInt returns a slice of integers containing a predefined set of values.
// The slice includes a mix of single and double-digit numbers, as well as a
// three-digit number. This function can be used for testing or demonstration
// purposes where a sample slice of integers is needed.
func sliceInt() []int {
	return []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
}

// sliceString returns a slice of strings containing a collection of random words.
// This function does not take any parameters and simply returns a predefined slice of strings.
func sliceString() []string {
	return []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
}

// sortByFirst sorts a slice of user structs by their first name in ascending order.
// It takes a slice of user structs as input and returns the sorted slice.
func sortByFirst(user []user) []user {
	sort.Sort(byFirst(user))

	return user
}

// sortByAge sorts a slice of user structs by their age in ascending order.
// It takes a slice of user structs as input and returns the sorted slice.
// The function uses the sort.Sort method with a custom byAge type that implements
// the sort.Interface to perform the sorting.
func sortByAge(user []user) []user {
	sort.Sort(byAge(user))

	return user
}
