package reportgrade

import "fmt"

type SchoolReport interface {
    Report()
    Sign(string)
}

type FouthGradeSchoolReport struct {}

func (*FouthGradeSchoolReport) Report() {
    fmt.Println("尊敬的XXX家长：")
    fmt.Println("... ...")
    fmt.Println("语文：62，数学：65，体育：98，自然63")
    fmt.Println("... ...")
    fmt.Println("家长签名：")
}

func (*FouthGradeSchoolReport) Sign(name string) {
    fmt.Printf("家长签名为：%s\n", name)
}

type HighScoreDecorator struct {
    SchoolReport
}

func (high *HighScoreDecorator) Report() {
    high.ReportHightScore()
    high.SchoolReport.Report()
}

func (*HighScoreDecorator) ReportHightScore() {
    fmt.Println("这次考试，语文最高分：75，数学最高分：78，英文最高分：80")
}

type SortDecorator struct {
    SchoolReport
}

func (sort *SortDecorator) Report() {
    sort.SchoolReport.Report()
    sort.ReportSort()
}

func (*SortDecorator) ReportSort() {
    fmt.Println("这次考试，本人排名第38名")
}
