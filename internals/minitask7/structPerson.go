package minitask7

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewUser(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (u *Person) Cetak() {
	fmt.Printf("Nama : %s\nAlamat: %s\nNo.hp: %s\n", u.Name, u.Address, u.Phone)
}

func (u *Person) Greet(name string) {
	fmt.Printf("Hallo %s\n", name)
}

func (u *Person) SetName(newName string) {
	u.Name = newName
}
