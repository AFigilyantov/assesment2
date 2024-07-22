package main

func findMaxIntvalueIn(m map[int]int) (int, int) {
	var resKey, resVal int

	for key, val := range m {
		if val > resVal {
			resKey = key
			resVal = val
		}

	}

	return resKey, resVal
}

func FindWayIn(mapForSearch map[int]map[int]int, currentNode int, maxAnswer *MaxAnswer) {
	var weight int
	maxAnswer.Way[currentNode] = struct{}{}
	val, ok := mapForSearch[currentNode]
	if maxAnswer.PrevNode != -1 {
		delete(val, maxAnswer.PrevNode)
	}
	if !ok {
		return
	}
	maxAnswer.PrevNode = currentNode
	currentNode, weight = findMaxIntvalueIn(val)

	if weight == 0 {
		return
	}

	_, okk := maxAnswer.Way[currentNode]

	if okk {
		return
	}
	maxAnswer.Weight += weight
	FindWayIn(mapForSearch, currentNode, maxAnswer)

}

type MaxAnswer struct {
	Way      map[int]struct{}
	Weight   int
	PrevNode int
}

func NewMaxAnswer() *MaxAnswer {

	var m MaxAnswer
	m.Way = make(map[int]struct{})
	m.Weight = 0
	m.PrevNode = -1

	return &m

}

func MakeCopyMap(mapSrc map[int]map[int]int) map[int]map[int]int {
	mapDest := make(map[int]map[int]int)

	for i := 0; i < len(mapSrc); i++ {
		mapDest[i] = make(map[int]int)
	}

	for key, line := range mapSrc {
		for index, val := range line {
			m := mapDest[key]
			m[index] = val
		}
	}
	return mapDest
}

func MaxUserAnswerScore(mapSrc map[int]map[int]int, ua []int) int {

	score := 0

	for i := 0; i < len(ua)-1; i++ {

		val, ok := mapSrc[ua[i]]
		if !ok {
			continue
		}
		score += val[ua[i+1]]
	}
	return score
}
