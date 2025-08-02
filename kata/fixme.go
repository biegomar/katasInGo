package kata

import (
	"fmt"
	"strings"
)

type Dinglemouse struct {
	name            string
	age             int
	sex             rune
	orderMap        map[int]string
	orderMapReverse map[string]int
	actualPosition  int
}

func NewDinglemouse() *Dinglemouse {
	return &Dinglemouse{
		orderMap:        make(map[int]string),
		orderMapReverse: make(map[string]int),
		actualPosition:  0,
	}

}

func (d *Dinglemouse) SetAge(age int) *Dinglemouse {
	d.setPosition("age")
	d.age = age
	return d
}

func (d *Dinglemouse) SetSex(sex rune) *Dinglemouse {
	d.setPosition("sex")
	d.sex = sex
	return d
}

func (d *Dinglemouse) SetName(name string) *Dinglemouse {
	d.setPosition("name")
	d.name = name
	return d
}

func (d *Dinglemouse) Hello() string {
	sexText := "male"
	if d.sex == 'F' {
		sexText = "female"
	}

	result := strings.Builder{}
	result.Grow(100)
	result.WriteString("Hello.")

	for i := 1; i <= len(d.orderMap); i++ {
		attribute := d.orderMap[i]
		if attribute == "name" {
			result.WriteString(fmt.Sprintf(" My name is %s.", d.name))
		}
		if attribute == "sex" {
			result.WriteString(fmt.Sprintf(" I am %s.", sexText))
		}
		if attribute == "age" {
			result.WriteString(fmt.Sprintf(" I am %d.", d.age))
		}
	}

	return result.String()
}

func (d *Dinglemouse) setPosition(attribute string) {
	if d.orderMapReverse[attribute] == 0 {
		d.actualPosition++
		d.orderMapReverse[attribute] = d.actualPosition
		d.orderMap[d.actualPosition] = attribute
	}
}
