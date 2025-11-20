package todo

type CreateTodoInput struct {
	Name        string
	Description string
}

type PatchTodoInput struct {
	ID          int
	Name        *string
	Description *string
	Completed   *bool
}
