package reportgrade

import "testing"

func TestDecorator(t *testing.T) {

    var report SchoolReport = &SortDecorator{
        SchoolReport: &HighScoreDecorator{
            SchoolReport: &FouthGradeSchoolReport{},
        },
    }

    report.Report()
    report.Sign("张三")

}
