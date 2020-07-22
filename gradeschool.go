package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type students struct {
	name  string
	grade int
}

func main() {
	student := []students{}
	var instruction []string
	var query []string
	var result []string
	var operation string
	var studentgrade int
	query = append(query, "Add")
	query = append(query, "Which")
	query = append(query, "Who")

	instruction = append(instruction, "Add Jim to grade 7")
	instruction = append(instruction, "Add Bert to grade 6")
	instruction = append(instruction, "Add Sally to grade 2")
	instruction = append(instruction, "Add Bertha to grade 2")
	instruction = append(instruction, "Add Kev to grade 2")
	instruction = append(instruction, "Add Brian to grade 4")
	instruction = append(instruction, "Which students are in grade 2")
	instruction = append(instruction, "Who all is enrolled in school right now?")

	for _, x := range instruction {
		for _, v := range query {
			if strings.Contains(x, v) {
				operation = v
				break
			}
		}

		switch operation {
		case "Add":
			holdstr := x[4:]
			i := strings.Index(holdstr, " ")
			j := strings.LastIndex(holdstr, " ")
			g, _ := strconv.Atoi(holdstr[j+1:])
			student = append(student, students{holdstr[:i], g})
		case "Which":
			m := strings.LastIndex(x, " ")
			studentgrade, _ = strconv.Atoi(x[m+1:])
			for _, z := range student {
				if z.grade == studentgrade {
					result = append(result, z.name)
				}
			}
		case "Who":
			//			sort.Slice(student, func(i, j int) bool {
			//				return student[i].grade < student[j].grade
			//			})

			sort.Slice(student, func(i, j int) bool {
				if student[i].name < student[j].name {
					return true
				} else {
					return false
				}
				return student[i].grade < student[j].grade
			})

		}

	}
	fmt.Printf("Student %v\n", student)
	fmt.Printf("Students enrolled in grade %v are %v\n", studentgrade, result)
	fmt.Printf("Students enrolled in school %v\n", student)
}
