package handlers

import "fmt"

type IHandler interface {
    Verify(int) bool
    HandleRequest(*Woman) bool
}

type Woman struct {
    level int  // 1--未出嫁；2--出嫁；3--夫死
    request string
}

type InCharged struct {
    IHandler
    next *InCharged
}

func (ic *InCharged) SetNext(next *InCharged) {
    ic.next = next
}

func (ic *InCharged) Verify(level int) bool {
    return true
}

func (ic *InCharged) HandleRequest(woman *Woman) bool {
    if ic.IHandler.Verify(woman.level) {
        return ic.IHandler.HandleRequest(woman)
    }
    if ic.next != nil {
        return ic.next.HandleRequest(woman)
    }
    return false
}

type FatherHandler struct {}

func (fh *FatherHandler) Verify(level int) bool {
    return level == FATHER_LEVEL
}

func (fh *FatherHandler) HandleRequest(woman *Woman) bool {
    fmt.Printf("[Father] 同意了他女儿的【%s】请求！\n", woman.request)
    return true
}

type HusbandHandler struct {}

func (hh *HusbandHandler) Verify(level int) bool {
    return level == HUSBAND_LEVEL
}

func (hh *HusbandHandler) HandleRequest(woman *Woman) bool {
    fmt.Printf("[Husband] 同意了他老婆的【%s】请求！\n", woman.request)
    return true
}

type SonHandler struct {}

func (sh *SonHandler) Verify(level int) bool {
    return level == SON_LEVEL
}

func (sh *SonHandler) HandleRequest(woman *Woman) bool {
    fmt.Printf("[Son] 同意了他老娘的【%s】请求！\n", woman.request)
    return true
}

var (
    FATHER_LEVEL, HUSBAND_LEVEL, SON_LEVEL int
)

func init() {
    FATHER_LEVEL, HUSBAND_LEVEL, SON_LEVEL = 1, 2, 3
}
