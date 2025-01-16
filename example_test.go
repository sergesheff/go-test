package testt

import "testing"

func ATest(t *testing.T) {
	if !Res(true) {
		t.FailNow()
	}

	t.FailNow()
}
