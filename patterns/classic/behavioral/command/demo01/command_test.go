package command

func ExampleBoxButtons() {
    board := &MotherBoard{}
    start := NewStartCommand(board)
    reboot := NewRebootCommand(board)

    box1 := NewBox(start, reboot)
    box1.PressButton1()
    box1.PressButton2()

    box2 := NewBox(reboot, start)
    box2.PressButton1()
    box2.PressButton2()

    //Output:
    //System is startint ...
    //System is Rebooting ...
    //System is Rebooting ...
    //System is startint ...
}
