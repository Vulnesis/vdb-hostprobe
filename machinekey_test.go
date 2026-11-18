package hostprobe

import "testing"

// A machine key must not change with the brand: the platform stores keys, and
// the CLI and Belai of both brands must agree on one key per machine. The
// expected value was computed with the original prefix.
func TestMachineKeyStableAcrossBrands(t *testing.T) {
	const id = "4C4C4544-0039-1056-8038-B7C04F4B5132"
	const want = "6f574212e4258c7dfde1a1b2d6411478b54e4fb70489798c85499c4c7085f131"
	if got := HashMachineID("4c4c4544003910568038b7c04f4b5132"); got != want {
		t.Fatalf("HashMachineID = %s, want %s", got, want)
	}
	if HashMachineID(id) == "" {
		t.Fatal("a real id must produce a key")
	}
}
