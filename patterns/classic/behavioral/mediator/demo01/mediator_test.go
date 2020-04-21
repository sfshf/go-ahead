package mediator

import "testing"

func TestMediator(t *testing.T) {

    mediator := GetMediatorInstance()
    mediator.Cd = &CDDriver{}
    mediator.Cpu = &CPU{}
    mediator.Video = &VideoCard{}
    mediator.Sound = &SoundCard{}

    mediator.Cd.ReadData()

    if mediator.Cd.Data != "music,image" {
        t.Fatal("CD unexcept data %s", mediator.Cd.Data)
    }

    if mediator.Cpu.Sound != "music" {
		t.Fatalf("CPU unexpect sound data %s", mediator.Cpu.Sound)
	}

	if mediator.Cpu.Video != "image" {
		t.Fatalf("CPU unexpect video data %s", mediator.Cpu.Video)
	}

	if mediator.Video.Data != "image" {
		t.Fatalf("VidoeCard unexpect data %s", mediator.Video.Data)
	}

	if mediator.Sound.Data != "music" {
		t.Fatalf("SoundCard unexpect data %s", mediator.Sound.Data)
	}

}
