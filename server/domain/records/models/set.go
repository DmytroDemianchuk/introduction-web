package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	OriginSets        map[string][]int
	ModeKey           string
	ModeIndexes       []int
	ModeSimilarity    map[string]float32
	UnitedSimilarSets map[string][]int
}

func SetNew(rows [][]string, margin float32) Set {
	var set Set
	set.InitSet(rows)
	set.FindMode()
	set.CountSimilarity()
	set.CombineSets(margin)
	return set
}

func (s *Set) InitSet(rows [][]string) {
	s.OriginSets = make(map[string][]int)

	for iRow, row := range rows {
		if len(row) != 0 {
			var indexes []int
			for iColumn, cell := range row {
				if cell != "" {
					indexes = append(indexes, iColumn+1)
				}
			}
			if len(indexes) > 0 {
				indexStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(indexes)), " "), "[]")
				s.OriginSets[indexStr] = append(s.OriginSets[indexStr], iRow+1)
			}
		}
	}
}

func (s *Set) FindMode() error {
	max := 0
	var index string
	for i, val := range s.OriginSets {
		if len(val) > max {
			max = len(val)
			index = i
		}
	}
	s.ModeKey = index
	indexes, err := s.stringToNumbers(index)
	if err != nil {
		return err
	}
	s.ModeIndexes = indexes
	return nil

}

func (s *Set) CombineSets(margin float32) {
	s.UnitedSimilarSets = make(map[string][]int)
	for k, val := range s.ModeSimilarity {
		if val > margin {
			s.UnitedSimilarSets[s.ModeKey] = append(s.UnitedSimilarSets[s.ModeKey], s.OriginSets[k]...)
		} else {
			s.UnitedSimilarSets[k] = s.OriginSets[k]
		}
	}
}

func (s *Set) CountSimilarity() error {
	s.ModeSimilarity = make(map[string]float32)

	for i := range s.OriginSets {
		indexes, err := s.stringToNumbers(i)
		if err != nil {
			return err
		}
		commonCount := 0
		for _, index := range indexes {
			if s.contains(index) {
				commonCount++
			}
		}
		allCount := s.UnionCount(indexes)
		s.ModeSimilarity[i] = (float32(commonCount) / float32(allCount))
	}
	return nil
}

func (s *Set) UnionCount(slice []int) int {
	unique := make(map[int]bool)

	var result int
	for _, item := range append(s.ModeIndexes, slice...) {
		if _, exists := unique[item]; !exists {
			result++
			unique[item] = true
		}
	}
	return result
}

func (s *Set) contains(value int) bool {
	for _, v := range s.ModeIndexes {
		if v == value {
			return true
		}
	}
	return false
}

func (s *Set) stringToNumbers(index string) ([]int, error) {
	var numbers []int
	substrings := strings.Split(index, " ")

	for _, str := range substrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
