package builder

//Builder是一个生成器接口
type Builder interface {
    Part1()
    Part2()
    Part3()
}

type Director struct {
    builder Builder
}

//NewDirector ...
func NewDirector(builder Builder) *Director {
    return &Director{
        builder: builder,
    }
}

//Construct product
func (d *Director) Construct() {
    d.builder.Part1()
    d.builder.Part2()
    d.builder.Part3()
}

type Builder1 struct {
    result string
}

func (b1 *Builder1) Part1() {
    b1.result += "-->Part1"
}

func (b1 *Builder1) Part2() {
    b1.result += "-->Part2"
}

func (b1 *Builder1) Part3() {
    b1.result += "-->Part3"
}

func (b1 *Builder1) GetResult() string {
    return b1.result
}

type Builder2 struct {
    result int
}

func (b2 *Builder2) Part1() {
    b2.result += 1
}

func (b2 *Builder2) Part2() {
    b2.result += 2
}

func (b2 *Builder2) Part3() {
    b2.result += 3
}

func (b2 *Builder2) GetResult() int {
    return b2.result
}
