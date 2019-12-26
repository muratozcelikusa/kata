package Kata



func Score (rolls []int) int{
 var sum int

	for i:=0;i<len(rolls);i++ {
		 var pins=rolls[i]

		if IsStrike(pins){
			if i<len(rolls){
				sum += GetNextTwo(rolls,i)
			}
		}else if IsSpare(rolls,pins){
			sum += GetNextTwo(rolls,i)
			i++
		}
		sum+=rolls[i]
	}
	return sum
}

func IsStrike(pins int) bool {
return pins == 10
}

func IsSpare(rolls []int,i int) bool {
return rolls[i]+GetNextOne(rolls,i) ==10
}

func GetNextOne (rolls []int,i int) int {
	nextRoll := i+1

	if nextRoll < len(rolls){
		return rolls[nextRoll]
	}
	return 0
}

func GetNextTwo (rolls []int,i int) int {
	nextRoll := i+1
	nextTwoRoll:=i+2
	if nextTwoRoll < len(rolls){
		return rolls[nextRoll] + rolls [nextTwoRoll]
	}
	return 0
}

