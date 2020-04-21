package mail

import "testing"

//非克隆方式，线程不安全
func TestMail(t *testing.T) {

    var MAX_COUNT int = 6
    mail := NewMail(NewAdvTemplate())
    mail.tail = "XX银行版权所有"
    for i := 0; i < MAX_COUNT; i ++ {
        mail.appellation = GetRandString(5) + "先生（女士）"
        mail.receiver = GetRandString(5) + "@" + GetRandString(8) + ".com"
        SendMail(mail)
    }

}

func TestCloneMail(t *testing.T) {

    var MAX_COUNT int = 6
    mail := NewMail(NewAdvTemplate())
    mail.tail = "XX银行版权所有"
    for i := 0; i < MAX_COUNT; i ++ {
        clonemail := mail.Clone().(*Mail)
        clonemail.appellation = GetRandString(5) + "先生（女士）"
        clonemail.receiver = GetRandString(5) + "@" + GetRandString(8) + ".com"
        SendMail(clonemail)
    }

}
