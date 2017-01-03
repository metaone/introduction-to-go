package count

import "testing"

func TestCountLines(t *testing.T) {
	got := CountLines("./13_counting_the_number_of_lines_in_a_file/testdata/alice.txt")
	want := 3736
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
