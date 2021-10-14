package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c Course) PrintClasses() {
	text := "Las clases del curso " + c.Name + " son: \n"
	for _, v := range c.Classes {
		text += "- " + v + "\n"
	}
	fmt.Println(text)
}

func (c *Course) AddUserId(id uint) []uint {
	c.UserIDs = append(c.UserIDs, id)
	return c.UserIDs
}
