package mail

import "testing"

func TestIPostOffice(t *testing.T) {

    context := "明天白天一起看周星驰那部今晚打老虎的电影。切记！切记！"
    address := "H城J街道K巷123栋303室"
    office := NewIPostOffice()
    office.SendLetter(context, address)

}
