package memento

func ExampleGameMemento() {

    game := &Game{
        hp: 100,
        mp: 100,
    }

    game.Status()
    progress := game.Save()

    game.Play(-80, -90)
    game.Status()

    game.Load(progress)
    game.Status()

    //Output:
    //Current HP: 100, MP: 100
    //Current HP: 20, MP: 10
    //Current HP: 100, MP: 100

}
