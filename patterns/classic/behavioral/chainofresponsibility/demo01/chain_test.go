package chain

func ExampleChain1() {

    c1 := NewChainNode(&GeneralHandler{})
    c2 := NewChainNode(&ProjectHandler{})
    c3 := NewChainNode(&DepHandler{})

    //串成链
    c1.SetNext(c2)
    c2.SetNext(c3)

    var c Handler = c1

    c.HandleFeeRequest("bob", 400)
    c.HandleFeeRequest("tom", 1400)
    c.HandleFeeRequest("ada", 10000)
    c.HandleFeeRequest("floar", 400)
    //Output:
    // Project handler permit bob 400 fee request
	// Dep handler permit tom 1400 fee request
	// General handler permit ada 10000 fee request
	// Project handler don't permit floar 400 fee request

}

func ExampleChain2() {
    c1 := NewChainNode(&DepHandler{})
    c2 := NewChainNode(&ProjectHandler{})
    c3 := NewChainNode(&GeneralHandler{})

	c1.SetNext(c2)
	c2.SetNext(c3)

	var c Handler = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
	// Output:
	// Project handler permit bob 400 fee request
	// Dep handler permit tom 1400 fee request
	// General handler permit ada 10000 fee request
	// Project handler don't permit floar 400 fee request

}
