package main

import "fmt"

type User struct {
	name string
}

type Key struct {
	ID   int
	name string
}

func main() {
	m := map[string]*User{
		"HR":  {"polette"},
		"CEO": {"polo"},
		"CFO": {"perin"},
	}

	fmt.Println(m)
	fmt.Println(m["HR"])

	//variable intermediaire (get polo if we change *User to User)
	hr := m["HR"]
	hr.name = "mou"

	fmt.Println(m["HR"])

	//assign new value

	m["CFO"] = &User{"test"}
	fmt.Println(m["CFO"])

	//double value - struct key

	res := make(map[Key]string)
	res[Key{1, "polo"}] = "file1"
	k := Key{2, "a new one"}
	res[k] = "file2"

	delete(res, Key{1, "polo"})
	fmt.Println(res)

}
