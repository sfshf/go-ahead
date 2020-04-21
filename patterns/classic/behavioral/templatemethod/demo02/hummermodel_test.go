package hummermodel

func ExampleHummerH1Model() {

    var model HummerModel = NewHummerH1Model()
    model.Run()
    //Output:
    //  h1 is starting ...
    //  h1's engine boom boom boom ...
    //  h1's alarm bell ringing ...
    //  h1 is stopping ...

}

func ExampleHummerH2Model() {

    var model HummerModel = NewHummerH2Model()
    model.Run()
    //Output:
    //  h2 is starting ...
    //  h2's engine boom boom boom ...
    //  h2's alarm bell ringing ...
    //  h2 is stopping ...

}