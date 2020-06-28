package main

import (
	"encoding/json"
    // "errors"
	"fmt"
    "log"
	// "os"
    "strings"
    "time"
)

type User struct {
	ID        int    `json:"id,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	RealName  string `json:"real_name,omitempty"`
	IdCardNum string `json:"identity_card_number,omitempty"`
	Ethnic    string `json:"ethnic,omitempty"`
	Profile   struct {
		Motto           string `json:"motto,omitempty"`
		Company         string `json:"company,omitempty"`
		Location        string `json:"location,omitempty"`
		Website         string `json:"website,omitempty"`
		TwitterUsername string `json:"twitter_username,omitempty"`
	} `json:"profile,omitempty"`
	StartAt   time.Time `json:"start_at,omitempty"`
	EndAt     time.Time `json:"end_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (user *User) UnmarshalJSON(data []byte) error {

    return nil

}

func (user User) MarshalJSON() ([]byte, error) {

    data := map[string]interface{}{
        "nick_name": user.NickName,
        "ethnic": user.Ethnic,
        "profile": user.Profile,
        "start_at": user.StartAt,
        "end_at": user.EndAt,
    }
    return json.Marshal(data)

}

func main() {

    createdAt, err := time.Parse("2006 年 01 月 02 日", "1997 年 07 月 19 日")
    if err != nil { log.Fatal(err) }
    model := User{
        ID: 1,
        NickName: "玩笑别开大",
        RealName: "庸嘉明",
        IdCardNum: "311132200001011234",
        Ethnic: "汉",
        StartAt: time.Now(),
        EndAt: time.Now().Add(5*time.Minute),
        CreatedAt: createdAt,
        UpdatedAt: createdAt.AddDate(0, 6, 0),
    }
    model.Profile.Motto = "生命不息，奋斗不止"
    model.Profile.Company = "心世界公司"
    model.Profile.Location = "地球村"
    model.Profile.Website = "https://www.example.com"
    model.Profile.TwitterUsername = "贾政金"
    // Marshal -- 假设从数据库取出的数据结构，进行处理后，发送给前端
    data, err := json.MarshalIndent(model, "", "    ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))

    // Unmarshal -- 假设将前端发来的数据进行解包，创建处结构
    m := new(*User)
    err = json.Unmarshal(data, m)
    if err != nil { log.Fatal(err) }
    fmt.Printf("%+v\n", m)

}
