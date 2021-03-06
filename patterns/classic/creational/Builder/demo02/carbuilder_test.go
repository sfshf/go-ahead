package carbuilder

func ExampleBenzModel() {
    var builder CarBuilder = &BenzBuilder{}
    var funcs []string = []string{"start", "alarm", "engineBoom", "stop"}
    var model *CarModel = builder.GetCarModel(funcs)
    model.Run()
    //Output:
    //benz is starting ...
    //benz is biu biu biu ...
    //benz is booming ...
    //benz is stopping ...
}

func ExampleBWMModel() {
    var builder CarBuilder = &BWMBuilder{}
    var funcs []string = []string{"start", "engineBoom", "alarm", "stop"}
    var model *CarModel = builder.GetCarModel(funcs)
    model.Run()
    //Output:
    //bmw is starting ...
    //bmw is booming ...
    //bmw is biu biu biu ...
    //bmw is stopping ...
}

func ExampleDirector() {
    var builder CarBuilder = &BWMBuilder{}
    var funcs []string = []string{"start", "engineBoom", "alarm", "stop"}
    director := NewDirector(builder)
    var model *CarModel = director.CarBuilder.GetCarModel(funcs)
    model.Run()
    //Output:
    //bmw is starting ...
    //bmw is booming ...
    //bmw is biu biu biu ...
    //bmw is stopping ...
}
