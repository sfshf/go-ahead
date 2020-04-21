package getouterinfo

import "fmt"

func ExampleAUser() {

    //模拟数据库取本公司员工数据
    infos := [6]string{"张三", "上海杨浦", "11111111111",
        "2222222", "Golang开发工程师", "3333333"}

    //生成对象
    var user IUserInfo = NewAUser(infos)

    //展示给领导
    fmt.Println(user.GetUserName())
    fmt.Println(user.GetHomeAddress())
    fmt.Println(user.GetMobileNumber())
    fmt.Println(user.GetOfficeTelNumber())
    fmt.Println(user.GetJobPosition())
    fmt.Println(user.GetHomeTelNumber())
    //Output:
    //张三
    //上海杨浦
    //11111111111
    //2222222
    //Golang开发工程师
    //3333333

}

func ExampleAdapter() {

    //模拟RMI取外包员工数据
    infos := [10]string{"李四", "35", "已婚",
        "74102589630", "3690258", "资深测试工程师",
        "18000", "7412584", "上海嘉定", "家里有四口人"}
    outeruser := NewIOuterUser(infos)

    //生成对象
    var user IUserInfo = NewAdapter(outeruser)

    //展示给领导
    fmt.Println(user.GetUserName())
    fmt.Println(user.GetHomeAddress())
    fmt.Println(user.GetMobileNumber())
    fmt.Println(user.GetOfficeTelNumber())
    fmt.Println(user.GetJobPosition())
    fmt.Println(user.GetHomeTelNumber())
    //Output:
    //李四
    //上海嘉定
    //74102589630
    //3690258
    //资深测试工程师
    //7412584

}
