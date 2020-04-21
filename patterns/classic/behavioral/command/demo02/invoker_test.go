package invoker

import "testing"

func TestDeletePageCommand(t *testing.T) {
    var command ICommand = &DeletePageCommand{}
    invoker := &Invoker{
        ICommand: command,
    }
    invoker.Action()
}

func TestAddRequirementCommand(t *testing.T) {
    var command ICommand = &AddRequirementCommand{}
    invoker := &Invoker{
        ICommand: command,
    }
    invoker.Action()
}
