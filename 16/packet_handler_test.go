package dec16

//
//import "testing"
//
//func TestGetPacketVersion(t *testing.T) {
//	result := GetPacketVersion([]int{0, 0, 1})
//
//	want := 1
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//	result = GetPacketVersion([]int{1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1})
//
//	want = 6
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//}
//
//func TestGetPacketTypeID(t *testing.T) {
//	result := GetPacketTypeID([]int{1, 1, 0, 0, 0, 1})
//
//	want := 1
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//	result = GetPacketTypeID([]int{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1})
//
//	want = 3
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//}
//
//func TestGetLiteralPacketData(t *testing.T) {
//	result := GetLiteralPacketValue([]int{0, 0, 0, 1, 1})
//
//	want := 3
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//	result = GetLiteralPacketValue([]int{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0})
//
//	want = 2021
//
//	if result != want {
//		t.Errorf("Incorrect parsing! Got %v, expected %v", result, want)
//	}
//
//}
//
