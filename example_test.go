package testt

import "testing"

func TestA(t *testing.T) {
	if !Res(true) {
		t.FailNow()
	}

	t.FailNow()

}
