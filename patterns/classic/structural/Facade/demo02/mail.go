package mail

import "fmt"

func NewIPostOffice() IPostOffice {
    return &CityZPostOffice{
        ILetterProcess: &Letter{},
        IPolice: &CityZPolice{},
    }
}

//邮局接口是个门面，它是用来发邮件的
type IPostOffice interface {
    SendLetter(context string, address string)
}

type CityZPostOffice struct {
    ILetterProcess
    IPolice
}

func (office *CityZPostOffice) SendLetter(context string, address string) {
    office.ILetterProcess.WriteContext(context)
    office.ILetterProcess.FillEnvelope(address)
    office.IPolice.CheckLetter(office.ILetterProcess)
    office.ILetterProcess.LetterIntoEnvelope()
    office.ILetterProcess.Send()
}

//子系统内部模块-邮件发送处理接口
type ILetterProcess interface {
    WriteContext(context string)
    FillEnvelope(address string)
    LetterIntoEnvelope()
    Send()
}

type Letter struct {}

func (letter *Letter) WriteContext(context string) {
    fmt.Printf("信件内容为：\n%s\n", context)
}

func (letter *Letter) FillEnvelope(address string) {
    fmt.Printf("信件接收地址为：\n%s\n", address)
}

func (letter *Letter) LetterIntoEnvelope() {
    fmt.Printf("将信件封装进信封里...\n")
}

func (letter *Letter) Send() {
    fmt.Printf("将邮件发送出去...\n")
}

//子系统内部模块-警察局安检处理接口
type IPolice interface {
    CheckLetter(ILetterProcess)
}

type CityZPolice struct {}

func (*CityZPolice) CheckLetter(letter ILetterProcess) {
    fmt.Printf("Z城市警察局已经检查过邮件...\n")
}
