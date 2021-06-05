package model

type Commands struct {
	Path, Command string
}

type ProjectModel struct {
	Name     string
	Commands []Commands
}
