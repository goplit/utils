package utilerror

const NotFoundError string = "not found"
const EmptyError string = "empty"

type AsError struct {
	err        error
	stackIndex int
}
