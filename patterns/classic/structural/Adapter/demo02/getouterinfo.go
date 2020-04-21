package getouterinfo

//a公司人员信息接口
type IUserInfo interface {
    GetUserName() string
    GetHomeAddress() string
    GetMobileNumber() string
    GetOfficeTelNumber() string
    GetJobPosition() string
    GetHomeTelNumber() string
}

//------------------------------------------------------------------------------

type aUser struct {
    name string
    address string
    mobilephone string
    officephone string
    job string
    homephone string
}

func NewAUser(infos [6]string) IUserInfo {
    return &aUser{
        name: infos[0],
        address: infos[1],
        mobilephone: infos[2],
        officephone: infos[3],
        job: infos[4],
        homephone: infos[5],
    }
}

func (a *aUser) GetUserName() string {
    return a.name
}

func (a *aUser) GetHomeAddress() string {
    return a.address
}

func (a *aUser) GetMobileNumber() string {
    return a.mobilephone
}

func (a *aUser) GetOfficeTelNumber() string {
    return a.officephone
}

func (a *aUser) GetJobPosition() string {
    return a.job
}

func (a *aUser) GetHomeTelNumber() string {
    return a.homephone
}

//------------------------------------------------------------------------------

//适配器不能获得远程对象的字段，只能调用其公开函数。
type Adapter struct {
    IOuterUser
}

func NewAdapter(outeruser IOuterUser) IUserInfo {
    return &Adapter{
        IOuterUser: outeruser,
    }
}

func (a *Adapter) GetUserName() string {
    baseinfos := a.IOuterUser.GetUserBaseInfo()
    return baseinfos["name"]
}

func (a *Adapter) GetHomeAddress() string {
    homeinfos := a.IOuterUser.GetUserHomeInfo()
    return homeinfos["address"]
}

func (a *Adapter) GetMobileNumber() string {
    baseinfos := a.IOuterUser.GetUserBaseInfo()
    return baseinfos["mobilephone"]
}

func (a *Adapter) GetOfficeTelNumber() string {
    officeinfos := a.IOuterUser.GetUserOfficeInfo()
    return officeinfos["officephone"]
}

func (a *Adapter) GetJobPosition() string {
    officeinfos := a.IOuterUser.GetUserOfficeInfo()
    return officeinfos["job"]
}

func (a *Adapter) GetHomeTelNumber() string {
    homeinfos := a.IOuterUser.GetUserHomeInfo()
    return homeinfos["homephone"]
}

//------------------------------------------------------------------------------

//b公司（外包公司）人员信息接口
type IOuterUser interface {
    GetUserBaseInfo() map[string]string
    GetUserOfficeInfo() map[string]string
    GetUserHomeInfo() map[string]string
}

func NewIOuterUser(infos [10]string) IOuterUser {
    bUser := &bUser{
        infos: make(map[string]string),
    }
    bUser.infos["name"] = infos[0]
    bUser.infos["age"] = infos[1]
    bUser.infos["married"] = infos[2]
    bUser.infos["mobilephone"] = infos[3]
    bUser.infos["officephone"] = infos[4]
    bUser.infos["job"] = infos[5]
    bUser.infos["salary"] = infos[6]
    bUser.infos["homephone"] = infos[7]
    bUser.infos["address"] = infos[8]
    bUser.infos["familynum"] = infos[9]
    return bUser
}

type bUser struct {
    infos map[string]string
}

func (b *bUser) GetUserBaseInfo() map[string]string {
    baseinfos := make(map[string]string)
    keys := make([]string, 0)
    keys = append(keys, "name")
    keys = append(keys, "age")
    keys = append(keys, "married")
    keys = append(keys, "mobilephone")
    for key, value := range b.infos {
        for _, target := range keys {
            if target == key {
                baseinfos[key] = value
            }
        }
    }
    return baseinfos
}

func (b *bUser) GetUserOfficeInfo() map[string]string {
    officeinfos := make(map[string]string)
    keys := make([]string, 0)
    keys = append(keys, "officephone")
    keys = append(keys, "job")
    keys = append(keys, "salary")
    for key, value := range b.infos {
        for _, target := range keys {
            if target == key {
                officeinfos[key] = value
            }
        }
    }
    return officeinfos
}

func (b *bUser) GetUserHomeInfo() map[string]string {
    homeinfos := make(map[string]string)
    keys := make([]string, 0)
    keys = append(keys, "homephone")
    keys = append(keys, "address")
    keys = append(keys, "familynum")
    for key, value := range b.infos {
        for _, target := range keys {
            if target == key {
                homeinfos[key] = value
            }
        }
    }
    return homeinfos
}
