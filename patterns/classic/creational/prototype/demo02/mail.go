package mail

import (
    "fmt"
    "strings"
    "math/rand"
)

type Cloneable interface {
    Clone() Cloneable
}

type AdvTemplate struct {
    advSubject string
    advContext string
}

func NewAdvTemplate() *AdvTemplate {
    return &AdvTemplate{
        advSubject: "XX银行国庆信用卡抽奖活动",
        advContext: "国庆抽奖活动通知：只要刷卡就送你一百万！...",
    }
}

type Mail struct {
    receiver string
    subject string
    appellation string
    context string
    tail string
}

func NewMail(at *AdvTemplate) *Mail {
    return &Mail{
        subject: at.advSubject,
        context: at.advContext,
    }
}

func (mail *Mail) Clone() Cloneable {
    mailcopy := *mail
    return &mailcopy
}

func SendMail(mail *Mail) {
    fmt.Println("标题：", mail.subject)
    fmt.Println("收件人：", mail.receiver)
    fmt.Println("发送成功！")
}

func GetRandString(length int) string {
    var source string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var strBuilder strings.Builder
    for i := 0; i < length; i ++ {
        at := rand.Intn(51)
        strBuilder.WriteString(source[at:at+1])
    }
    return strBuilder.String()
}
