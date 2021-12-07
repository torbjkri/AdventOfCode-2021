package dec6

import (
	"my_utils"
)

type LanternFish struct {
	clones_today []int
	total_days   int
	current_day  int

	total_fishies int
}

func NewLanternFish(days_total int) *LanternFish {
	f := new(LanternFish)
	f.clones_today = make([]int, days_total+9)
	f.total_days = days_total
	f.current_day = 0

	for idx, _ := range f.clones_today {
		f.clones_today[idx] = 0
	}
	f.total_fishies = 0
	return f
}

func (f *LanternFish) AddSpawner(the_day int) {
	f.clones_today[the_day] += 1
	f.total_fishies += 1
}

func (f *LanternFish) Update() {

	f.total_fishies += f.clones_today[f.current_day]
	f.clones_today[f.current_day+7] += f.clones_today[f.current_day]
	f.clones_today[f.current_day+9] += f.clones_today[f.current_day]

	f.current_day += 1
}

func GetInitialData(filename string) []int {
	data := my_utils.ReadFile(filename)

	return my_utils.StringToNumList(data, ",")
}

func Solve1() int {

	num_days := 18

	initial_fish := GetInitialData("C:/Dev/Go/AdventOfCode-2021/6/data")

	fishie := NewLanternFish(num_days)

	for _, input := range initial_fish {
		fishie.AddSpawner(input)
	}

	for day := 0; day < num_days; day++ {
		fishie.Update()
	}
	total_fishies := fishie.total_fishies

	return total_fishies
}

func Solve2() int {

	num_days := 256

	initial_fish := GetInitialData("C:/Dev/Go/AdventOfCode-2021/6/data")

	fishie := NewLanternFish(num_days)

	for _, input := range initial_fish {
		fishie.AddSpawner(input)
	}

	for day := 0; day < num_days; day++ {
		fishie.Update()
	}
	total_fishies := fishie.total_fishies

	return total_fishies
}
