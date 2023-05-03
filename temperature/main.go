package main

import "fmt"

type TemperatureForecast struct {
	value int
	next  int
	last  bool
}

func main() {
	t1 := TemperatureForecast{
		value: 13,
	}
	t2 := TemperatureForecast{
		value: 12,
	}
	t3 := TemperatureForecast{
		value: 15,
	}
	t4 := TemperatureForecast{
		value: 11,
	}
	t5 := TemperatureForecast{
		value: 9,
	}
	t6 := TemperatureForecast{
		value: 12,
	}
	t7 := TemperatureForecast{
		value: 16,
		last:  true,
	}

	allTemp := make([]TemperatureForecast, 0, 7)
	allTemp = append(allTemp, t1)
	allTemp = append(allTemp, t2)
	allTemp = append(allTemp, t3)
	allTemp = append(allTemp, t4)
	allTemp = append(allTemp, t5)
	allTemp = append(allTemp, t6)
	allTemp = append(allTemp, t7)

	//for i := 0; i < len(allTemp); i++ {
	//	Find(i, allTemp)
	//}
	for i := 0; i < len(allTemp); i++ {
		FindMe(allTemp[i:])
	}

	fmt.Println(allTemp)
}

func FindMe(t []TemperatureForecast) {
	step := 1
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if j+1 == len(t) {
				fmt.Println(t[i].value)
				fmt.Println("LAST")
				t[i].next = step
				return
			}

			if t[i].value < t[j].value {
				t[i].next = step

				return
			}
			step++
			//fmt.Println(t[i].value)
		}
	}

}
