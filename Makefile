todo/list:
	@go run ./cmd/todo -list

todo/add:
	@go run ./cmd/todo/ -add "$(title)"

todo/edit: 
	@go run ./cmd/todo/ -edit "$(edit)"

todo/toggle: 
	@go run ./cmd/todo/ -completeToggle $(id)

todo/delete: 
	@go run ./cmd/todo/ -delete $(id)

build/Todo:
	@go build ./cmd/todo
