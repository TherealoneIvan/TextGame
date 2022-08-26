package main

//пути куда можно пойти
type Ways struct {
	ways   map[string]Environment
	output []string
}

//Описание места
type Environment struct {
	name            string
	inventoryObject map[string]bool                       // предметы которые может взять игрок
	lookAroundFunc  map[string]func(e Environment) Answer //
	objInt          map[string]bool                       // предметы для взаимодействия игрока например дверь
	Ways
	//inventoryOutput []InventoryObject
	ObjMap map[InventoryObject]map[ObjectsOfInteraction]bool
}
type ObjectsOfInteraction struct {
	name string
}

//ты в своей комнате. можно пройти - коридор"},
//		{4, "осмотреться", "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор"},
func LookAroundRoom(e Environment) Answer {
	var w string
	if Frst {
		Frst = false
		return Answer{"ты в своей комнате. можно пройти - коридор"}
	} else {
		for _, i := range e.output {
			w = w + i + ", "
			w = w[:len(w)-2]
		}
		//return
	}
	var item string = "на столе: "
	for i := range e.inventoryObject {
		if i == "рюкзак" {
			item = item + "на стуле - " + i
		}
		item = item + " " + i
	}
	if len(e.inventoryObject) != 0 {
		return Answer{item + ". можно пройти - " + w}
	}
	return Answer{"пустая комната. можно пройти - " + w}
}

func waysToString(e Environment) string {
	var w string
	if len(e.ways) == 1 {
		for i := range e.ways {
			w = w + " " + i
		}
	} else {
		for i := range e.ways {
			w = w + i + ", "
		}
		w = w[:len(w)-2]
	}
	return w
}
func LookAroundKitchen(e Environment) Answer {
	w := waysToString(e)
	return Answer{"ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти -" + w}
}
func LookAroundCorridor(e Environment) Answer {
	w := waysToString(e)
	return Answer{"ничего интересного. можно пройти - " + w}
}
func LookAroundStreet(e Environment) Answer {
	w := waysToString(e)
	return Answer{"на улице весна.можно пройти - " + w}
}
func WayReceiver(way string) Environment {
	switch way {
	case "кухня":
		return Kitchen
	case "коридор":
		return Corridor
	case "комната":
		return Room
	case "улица":
		return Street
	}
	return Environment{}
}
