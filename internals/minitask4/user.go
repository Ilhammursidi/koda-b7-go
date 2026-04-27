package minitask4

type User struct {
	// id        int
	Name      string
	Image     string
	Email     string
	Umur      int
	Phone     string
	IsMarried bool
	Education Education
}

type Education struct {
	Name  string
	Major string
}
