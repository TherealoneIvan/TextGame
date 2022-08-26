package main

import "fmt"

type Person struct {
	Name         string
	isBackPackOn bool
	backPack     BackPack
	e            Environment
}

func (p Person) Walk(environment Environment, person *Person) Answer {
	Frst = true
	if environment.name == "улица" && !DoorIsOpen {
		return Answer{answer: "дверь закрыта"}
	}
	_, f := person.e.ways[environment.name]
	if !f {
		return Answer{"нет пути в " + environment.name}
	} else {
		person.e = environment
		return person.LookAround(environment, *person)
	}
	return Answer{}
}
func (p Person) LookAround(environment Environment, person Person) Answer {
	fmt.Println(environment.name)
	return person.e.lookAroundFunc[environment.name](environment)
}
func (p Person) TakeBackPack(person *Person) Answer {
	if person.e.inventoryObject["рюкзак"] {
		person.isBackPackOn = true
		person.e.inventoryObject["рюкзак"] = false
		return Answer{"вы одели: рюкзак"}
	}
	return Answer{"нет такого"}
}

func (p Person) GrabItem(object InventoryObject, person *Person) Answer {
	if person.isBackPackOn && person.e.inventoryObject[object.name] {
		person.backPack.Add(object)
		person.e.inventoryObject[object.name] = false
		return Answer{"предмет добавлен в инвентарь: " + object.name}
	} else if !person.isBackPackOn {
		return Answer{"некуда класть"}
	} else if !person.e.inventoryObject[object.name] {
		return Answer{"нет такого"}
	}
	return Answer{}
}

func (p Person) Use(item InventoryObject, iterable ObjectsOfInteraction, person *Person) Answer {
	result := person.backPack.UseItem(item, iterable, person)
	return Answer{result}
}
