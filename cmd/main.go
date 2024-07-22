package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
)

func main() {
	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}
	err := SquareValidation(mtx1)
	if err != nil {
		log.Fatal(err)
	}

	ua := []int{ /*4, 1, 0, 2*/ }

	err = LoopValidation(mtx1)

	if err != nil {
		log.Fatal(err)
	}

	err = UniqeValidation(ua)

	if err != nil {
		log.Fatal(err)
	}

	err = RangeValidation(mtx1, ua)

	if err != nil {
		log.Fatal(err)
	}

	maxGrade := calMaxGrade(mtx1)
	userGrade := calcUserGrade(mtx1, ua)

	fmt.Println(maxGrade)
	fmt.Println(userGrade)

	fmt.Println("main have finished")

}

func EvalSequence(matrix [][]int, userAnswer []int) (int, error) {

	err := Validation(matrix, userAnswer)

	if err != nil {
		return 0, err
	}

	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	percent := userGrade * 100 / maxGrade

	return percent, err
}

func Validation(matrix [][]int, userAnswer []int) error {
	err := SquareValidation(matrix)
	if err != nil {
		return err
	}

	err = LoopValidation(matrix)

	if err != nil {
		return err
	}

	err = UniqeValidation(userAnswer)

	if err != nil {
		return err
	}

	err = RangeValidation(matrix, userAnswer)

	if err != nil {
		return err
	}
	return nil
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {

	var g GraphLists

	g.FillAdjList(matrix)

	result := MaxUserAnswerScore(g.adjList, userAnswer)

	return result
}

func SquareValidation(matrix [][]int) error {
	height := len(matrix)
	for _, width := range matrix {
		if height != len(width) {
			return errors.New("matrix is not square")
		}
	}
	return nil
}

func LoopValidation(matrix [][]int) error {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][i] != 0 {
			return errors.New("matrix has loop")
		}
	}
	return nil
}

func UniqeValidation(userAnswer []int) error {
	tempMap := make(map[int]struct{})
	for _, answer := range userAnswer {
		tempMap[answer] = struct{}{}
	}
	if len(userAnswer) != len(tempMap) {
		return errors.New("user's answers are not unique")
	}
	return nil
}

func RangeValidation(matrix [][]int, userAnswer []int) error {
	if len(userAnswer) == 0 {
		return errors.New("user's answers are empty")
	}
	workSlice := slices.Clone(userAnswer)
	slices.Sort(workSlice)
	maxAnswer := userAnswer[len(userAnswer)-1]
	if maxAnswer > len(matrix) {
		return errors.New("user's answers are out of range")
	}
	return nil

}

func calMaxGrade(mat [][]int) int {
	var g GraphLists

	g.FillAdjList(mat)

	result := 0

	for key := range mat {
		maxAnswer := NewMaxAnswer()
		mapForSearch := MakeCopyMap(g.adjList) //maps, COPY and Clone do not work
		FindWayIn(mapForSearch, key, maxAnswer)
		if maxAnswer.Weight > result {
			result = maxAnswer.Weight
		}

	}

	return result

}
