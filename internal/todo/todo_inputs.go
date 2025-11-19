package todo

type PatchTodoInput struct {
	ID          int
	Name        *string
	Description *string
	Completed   *bool
}
