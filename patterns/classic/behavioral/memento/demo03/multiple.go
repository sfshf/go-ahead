package multiple

import (
    "reflect"
    "fmt"
    "math/rand"
)

type IMemento interface {
    Backupid() string
}

type MementoManager struct {
    Mementos map[string]IMemento
}

func (mm *MementoManager) GetMemento(backupid string) IMemento {
    return mm.Mementos[backupid]
}

func (mm *MementoManager) AddMemento(backupid string, memento IMemento) {
    mm.Mementos[backupid] = memento
}

type Originator struct {
    State1 int
    State2 float64
    State3 string
}

func NewOriginator(state1 int, state2 float64, state3 string) *Originator {
    return &Originator{
        State1: state1,
        State2: state2,
        State3: state3,
    }
}

func (o *Originator) Show() {
    fmt.Printf("[Originator] state1 = %#v, state2 = %#v, state3 = %#v\n",
        o.State1, o.State2, o.State3)
}

type OriginatorMemento struct {
    mm *MementoManager
    States map[string]interface{}
}

func (om *OriginatorMemento) Backupid() string {
    for id, m := range om.mm.Mementos {
        if m == om {
            return id
        }
    }
    return "[ERROR]NOT IN MEMENTO MAP."
}

type BackupUtils struct {
    mm *MementoManager
}

func NewBackUtils() *BackupUtils {
    return &BackupUtils{
        mm: &MementoManager{
            Mementos: make(map[string]IMemento),
        },
    }
}

func (bu *BackupUtils) BackupProp(obj interface{}) *OriginatorMemento {
    //创建一个具体的memento对象
    originatorMemento := &OriginatorMemento{
        mm: bu.mm,
        States: make(map[string]interface{}),
    }
    structType := reflect.TypeOf(obj).Elem()
    structValue := reflect.ValueOf(obj).Elem()
    numOfFields := structType.NumField()
    for i := 0; i < numOfFields; i ++ {
        fieldName := structType.Field(i).Name
        fieldValue := structValue.Field(i)
        switch  fieldValue.Kind() {
        case reflect.Int:
            // fmt.Println(fieldName)
            // fmt.Println(fieldValue.Int())
            //此处一个细节：注意Int()返回的是int64类型
            originatorMemento.States[fieldName] = fieldValue.Int()
        case reflect.Float64:
            originatorMemento.States[fieldName] = fieldValue.Float()
        case reflect.String:
            originatorMemento.States[fieldName] = fieldValue.String()
        }
    }
    //将生成好的memento对象加入到manager对象中
    backupid := string(rand.Intn(100)) //这里可以使用更合适的方式
    bu.mm.Mementos[backupid] = originatorMemento
    return originatorMemento
}

func (bu *BackupUtils) RestoreProp(obj interface{}, backupid string) {

    structType := reflect.TypeOf(obj).Elem()
    structValue := reflect.ValueOf(obj).Elem()
    numOfFields := structType.NumField()
    for id, m := range bu.mm.Mementos {
        if backupid == id {
            for i := 0; i < numOfFields; i ++ {
                fieldName := structType.Field(i).Name
                for key, value := range m.(*OriginatorMemento).States {
                    if key == fieldName {
                        fieldValue := reflect.ValueOf(value)
                        // fmt.Println(fieldValue.Kind())
                        switch fieldValue.Kind() {
                        case reflect.Int64:
                            structValue.FieldByName(fieldName).SetInt(fieldValue.Int())
                        case reflect.Float64:
                            structValue.FieldByName(fieldName).SetFloat(fieldValue.Float())
                        case reflect.String:
                            structValue.FieldByName(fieldName).SetString(fieldValue.String())
                        }
                    }
                }
            }
        }
    }
}
