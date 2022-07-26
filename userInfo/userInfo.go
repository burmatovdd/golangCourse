package userInfo

import (
	"fmt"
	"homework/models"
)

type Service struct{}

type UserInfo interface {
	AddUser() models.User
}

// todo: add input check

func addProject(name string) models.Project {

	project := models.Project{
		name,
	}
	return project
}

func addWork(company string) models.Work {
	projects := []models.Project{}
	var name string

	for {
		fmt.Println("Enter project name (use CamelCase) for ", company)
		fmt.Scan(&name)
		if name == "0" {
			fmt.Println("here")
			break
		}
		projects = append(projects, addProject(name))
	}

	work := models.Work{
		company,
		projects,
	}
	return work
}

func addStack() []models.Stack {
	stack := []models.Stack{}

	stack = append(stack, models.Stack{
		"Golang",
	})
	stack = append(stack, models.Stack{
		"Vue3",
	})
	stack = append(stack, models.Stack{
		"Docker",
	})
	stack = append(stack, models.Stack{
		"Git",
	})
	stack = append(stack, models.Stack{
		"NodeJS",
	})
	return stack
}

func (service *Service) AddUser() models.User {
	courses := []models.Course{}
	courses = append(courses, models.Course{"Basic Golang"})
	courses = append(courses, models.Course{"Vue3"})
	courses = append(courses, models.Course{"Docker"})

	works := []models.Work{}

	works = append(works, addWork("yandex"))
	works = append(works, addWork("MailGroup"))
	works = append(works, addWork("Google"))

	edu := models.Education{"Mirea", courses}
	user := models.User{
		Name:       "John",
		Surname:    "Smith",
		Patronymic: "No",
		Education:  edu,
		Work:       works,
		Stack:      addStack(),
	}

	return user
}
