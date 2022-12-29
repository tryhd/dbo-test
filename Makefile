FILENAME := $(filter-out $@,$(MAKECMDGOALS))

all:
	@mkdir -p app/types
	@mkdir -p app/controllers
	@go run app/helper/command/command.go $(FILENAME)
