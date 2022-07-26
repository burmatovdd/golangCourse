package models

type User struct {
	Name       string
	Surname    string
	Patronymic string
	Education  Education
	Work       []Work
	Stack      []Stack
}

type Education struct {
	MainEdu string
	Courses []Course
}

type Course struct {
	Name string
}

type Work struct {
	Company  string
	Projects []Project
}

type Project struct {
	Name string
}

type Stack struct {
	Name string
}
