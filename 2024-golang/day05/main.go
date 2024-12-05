package day05 // Change this for each new day

import (
	"slices"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type PageRule struct {
	firstPage  int
	secondPage int
}

func NewPageRule(str string) PageRule {
	parts := strings.Split(str, "|")
	first, _ := utils.GetIntFromString(parts[0])
	second, _ := utils.GetIntFromString(parts[1])
	return PageRule{firstPage: first, secondPage: second}
}

type Update struct {
	pages []int
}

func NewUpdate(str string) Update {
	parts := strings.Split(str, ",")
	pages := make([]int, 0)
	for _, part := range parts {
		page, _ := utils.GetIntFromString(part)
		pages = append(pages, page)
	}
	return Update{pages: pages}
}

func BuildPageRulesMap(rules []PageRule) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		if _, ok := rulesMap[rule.firstPage]; !ok {
			rulesMap[rule.firstPage] = make([]int, 0)
		}
		rulesMap[rule.firstPage] = append(rulesMap[rule.firstPage], rule.secondPage)
	}
	return rulesMap
}

func IsValidUpdate(update Update, rulesMap map[int][]int) bool {
	if len(update.pages) == 1 {
		return true
	}
	for _, page := range update.pages {

		if slices.Contains(rulesMap[page], update.pages[1]) {
			return IsValidUpdate(Update{pages: update.pages[1:]}, rulesMap)
		} else {
			return false
		}
	}
	return false
}

func BuildPageRulesMapFromInput(input []string) map[int][]int {
	rules := make([]PageRule, 0)
	for _, line := range input {
		if strings.Contains(line, "|") {
			rules = append(rules, NewPageRule(line))
		}
	}
	return BuildPageRulesMap(rules)
}

func BuildUpdatesFromInput(input []string) []Update {
	updates := make([]Update, 0)
	for _, line := range input {
		if strings.Contains(line, ",") {
			updates = append(updates, NewUpdate(line))
		}
	}
	return updates
}

func Solve1(input []string) int {
	sum := 0
	updates := BuildUpdatesFromInput(input)
	rulesMap := BuildPageRulesMapFromInput(input)
	for _, update := range updates {
		if IsValidUpdate(update, rulesMap) {
			sum += update.pages[(len(update.pages)-1)/2]
		}
	}

	return sum
}

func SortUpdate(update Update, rulesMap map[int][]int) Update {
	if IsValidUpdate(update, rulesMap) {
		return update
	}
	for i := 0; i < len(update.pages)-1; i++ {
		if !slices.Contains(rulesMap[update.pages[i]], update.pages[i+1]) {
			update.pages[i], update.pages[i+1] = update.pages[i+1], update.pages[i]
			return SortUpdate(update, rulesMap)
		}
	}
	return update
}

func Solve2(input []string) int {
	sum := 0
	updates := BuildUpdatesFromInput(input)
	rulesMap := BuildPageRulesMapFromInput(input)
	for _, update := range updates {
		if !IsValidUpdate(update, rulesMap) {

			sum += SortUpdate(update, rulesMap).pages[(len(update.pages)-1)/2]
		}
	}

	return sum
}
