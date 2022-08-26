package main

import "fmt"

type InventoryObject struct {
	name string
	Action
}

type BackPack struct {
	backpack map[InventoryObject]bool
	//backpackUsage map[InventoryObject]func(object ObjectsOfInteraction) Answer
}

var (
	Obj             = make(map[InventoryObject]map[ObjectsOfInteraction]bool)
	DoorIsOpen bool = false
)

func (bp BackPack) Add(object InventoryObject) {
	bp.backpack[object] = true
}
func (bp BackPack) CheckItem(object InventoryObject, person2 *Person) bool {
	return person2.backPack.backpack[object]
}

func (bp BackPack) UseItem(item InventoryObject, interaction ObjectsOfInteraction, person *Person) string {
	if !person.backPack.CheckItem(item, person) {
		return "нет предмета в инвентаре - " + item.name
	}
	//for _, i := range Obj {
	//	for j, t := range i {
	//		fmt.Println(j, t)
	//	}
	//}
	fmt.Println(item == InventoryObject{
		name:   "ключ",
		Action: Action{},
	})
	fmt.Println(len(Obj[InventoryObject{
		name:   "ключ",
		Action: Action{},
	}]))
	if !Obj[item][interaction] {
		return "не к чему применить"
	}
	DoorIsOpen = true
	return "дверь открыта"

}
