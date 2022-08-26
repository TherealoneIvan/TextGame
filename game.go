package main

import (
	"fmt"
	"strings"
)

var (
	actions  map[Action]func()
	person   Person
	Kitchen  Environment
	Room     Environment
	Corridor Environment
	Street   Environment
	Key      InventoryObject
	Conspect InventoryObject
	Backpack InventoryObject
	Door     ObjectsOfInteraction
	Frst     bool = true
)

func itemReceiver(it string) InventoryObject {
	switch it {
	case "ключи":
		return Key
	default:
		return InventoryObject{
			name:   it,
			Action: Action{},
		}
	}
}
func handleCommand(command string) interface{} {
	commandSlice := strings.Fields(command)
	switch commandSlice[0] {
	case "осмотреться":
		return person.LookAround(person.e, person).answer
	case "идти":
		return person.Walk(WayReceiver(commandSlice[1]), &person).answer
	case "взять":
		return person.GrabItem(InventoryObject{
			name:   commandSlice[1],
			Action: Action{},
		}, &person).answer
	case "применить":
		return person.backPack.UseItem(InventoryObject{
			name:   commandSlice[1],
			Action: Action{},
		}, ObjectsOfInteraction{name: commandSlice[2]}, &person)
	case "одеть":
		return person.TakeBackPack(&person).answer
	default:
		return "неизвестная команда"
	}
	return nil
}
func initGame() {
	Street = Environment{
		name:            "улица",
		inventoryObject: nil,
		lookAroundFunc: map[string]func(e Environment) Answer{
			"кухня":   LookAroundKitchen,
			"коридор": LookAroundCorridor,
			"комната": LookAroundRoom,
			"улица":   LookAroundStreet,
		},
		objInt: nil,
		Ways: Ways{ways: map[string]Environment{
			"коридор": Corridor,
		}},
	}
	Room = Environment{
		name: "комната",
		inventoryObject: map[string]bool{
			"рюкзак":    true,
			"ключи":     true,
			"конспекты": true,
		},
		lookAroundFunc: map[string]func(e Environment) Answer{
			"кухня":    LookAroundKitchen,
			"корридор": LookAroundCorridor,
			"комната":  LookAroundRoom,
			"улица":    LookAroundStreet,
		},
		objInt: nil,
		Ways: Ways{ways: map[string]Environment{
			"коридор": Corridor,
		}},
	}
	Corridor = Environment{
		name:            "коридор",
		inventoryObject: nil,
		lookAroundFunc: map[string]func(e Environment) Answer{
			"кухня":   LookAroundKitchen,
			"коридор": LookAroundCorridor,
			"комната": LookAroundRoom,
			"улица":   LookAroundStreet,
		},
		objInt: map[string]bool{
			"дверь": true,
		},
		Ways: Ways{ways: map[string]Environment{
			"кухня":   Kitchen,
			"комната": Room,
			"улица":   Street,
		},
			output: []string{"кухня", "комната", "улица"},
		},
	}
	Kitchen = Environment{
		name:            "кухня",
		inventoryObject: nil,
		lookAroundFunc: map[string]func(e Environment) Answer{
			"кухня":   LookAroundKitchen,
			"коридор": LookAroundCorridor,
			"комната": LookAroundRoom,
			"улица":   LookAroundStreet,
		},
		objInt: nil,
		Ways: Ways{ways: map[string]Environment{
			"коридор": Corridor,
		}},
	}
	fmt.Println(Kitchen.ways["коридор"].name)

	person = Person{
		Name:         "Первый игрок",
		isBackPackOn: false,
		backPack: BackPack{
			backpack: map[InventoryObject]bool{},
		},
		e: Kitchen,
	}
	Key = InventoryObject{
		name:   "ключи",
		Action: Action{},
	}
	Conspect = InventoryObject{
		name:   "конспекты",
		Action: Action{},
	}
	Backpack = InventoryObject{
		name:   "рюкзак",
		Action: Action{},
	}
	Door = ObjectsOfInteraction{name: "дверь"}

	Obj[Key] = map[ObjectsOfInteraction]bool{
		Door: true,
	}
	Obj[Conspect] = make(map[ObjectsOfInteraction]bool)
	Obj[Backpack] = make(map[ObjectsOfInteraction]bool)
	DoorIsOpen = false
}
func checkWays(environment Environment) {
	for i := range environment.ways {
		fmt.Println(i + " kkkk")
	}
}

type gameCase struct {
	step    int
	command string
	answer  string
}

var game0cases = [][]gameCase{
	[]gameCase{
		{1, "осмотреться", "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"}, // действие осмотреться
		{2, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},                                         // действие идти
		{3, "идти комната", "ты в своей комнате. можно пройти - коридор"},
		{4, "осмотреться", "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор"},
		{5, "одеть рюкзак", "вы одели: рюкзак"},                   // действие одеть
		{6, "взять ключи", "предмет добавлен в инвентарь: ключи"}, // действие взять
		{7, "взять конспекты", "предмет добавлен в инвентарь: конспекты"},
		{8, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{9, "применить ключи дверь", "дверь открыта"}, // действие применить
		{10, "идти улица", "на улице весна. можно пройти - домой"},
	},

	[]gameCase{
		{1, "осмотреться", "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
		{2, "завтракать", "неизвестная команда"},  // придёт топать в универ голодным :(
		{3, "идти комната", "нет пути в комната"}, // через стены ходить нельзя
		{4, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{5, "применить ключи дверь", "нет предмета в инвентаре - ключи"},
		{6, "идти комната", "ты в своей комнате. можно пройти - коридор"},
		{7, "осмотреться", "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор"},
		{8, "взять ключи", "некуда класть"}, // надо взять рюкзак сначала
		{9, "одеть рюкзак", "вы одели: рюкзак"},
		{10, "осмотреться", "на столе: ключи, конспекты. можно пройти - коридор"}, // состояние изменилось
		{11, "взять ключи", "предмет добавлен в инвентарь: ключи"},
		{12, "взять телефон", "нет такого"},                                // неизвестный предмет
		{13, "взять ключи", "нет такого"},                                  // предмента уже нет в комнатеы - мы его взяли
		{14, "осмотреться", "на столе: конспекты. можно пройти - коридор"}, // состояние изменилось
		{15, "взять конспекты", "предмет добавлен в инвентарь: конспекты"},
		{16, "осмотреться", "пустая комната. можно пройти - коридор"}, // состояние изменилось
		{17, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{18, "идти кухня", "кухня, ничего интересного. можно пройти - коридор"},
		{19, "осмотреться", "ты находишься на кухне, на столе чай, надо идти в универ. можно пройти - коридор"}, // состояние изменилось
		{20, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{21, "идти улица", "дверь закрыта"},                                  //условие не удовлетворено
		{22, "применить ключи дверь", "дверь открыта"},                       //состояние изменилось
		{23, "применить телефон шкаф", "нет предмета в инвентаре - телефон"}, // нет предмета
		{24, "применить ключи шкаф", "не к чему применить"},                  // предмет есть, но применить его к этому нельзя
		{25, "идти улица", "на улице весна. можно пройти - домой"},
	},
}

func main() {
	a := InventoryObject{
		name:   "1",
		Action: Action{},
	}
	b := InventoryObject{
		name:   "1",
		Action: Action{},
	}
	fmt.Println(a == b)
	for caseNum, commands := range game0cases {
		initGame()
		fmt.Println("testing1")
		for _, item := range commands {
			answer := handleCommand(item.command)
			if answer != item.answer {
				fmt.Println("case:", caseNum, item.step,
					"\n\tcmd:", item.command,
					"\n\tresult:  ", answer,
					"\n\texpected:", item.answer)
			}
		}
	}

}
