package static

import "testing"

func TestLoadStatics(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics == nil {
		t.Fatalf("LoadStatics() is nil")
	}
}

func TestLoadStaticsDTDFile(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics.DTDFile == "" {
		t.Fatalf("PAAFile is empty")
	}
}

func TestLoadStaticsFS(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics.FS == nil {
		t.Fatalf("FS is nil")
	}
}

func TestLoadStaticsFSDTD(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics.FS == nil {
		t.Fatalf("FS is nil")
	}

	dtd, err := statics.FS.Open("squad.dtd")
	if err != nil {
		t.Fatalf("LoadStatics() squad.dtd error: %s", err)
	}
	if dtd == nil {
		t.Fatalf("squad.dtd is nil")
	}
}

func TestLoadStaticsFSPAA(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics.FS == nil {
		t.Fatalf("FS is nil")
	}

	paa, err := statics.FS.Open("squad.paa")
	if err != nil {
		t.Fatalf("LoadStatics() squad.paa error: %s", err)
	}
	if paa == nil {
		t.Fatalf("squad.paa is nil")
	}
}

func TestLoadStaticsPAAFile(t *testing.T) {
	statics, err := LoadStatics()
	if err != nil {
		t.Fatalf("LoadStatics() error: %s", err)
	}
	if statics.PAAFile == "" {
		t.Fatalf("PAAFile is empty")
	}
}
