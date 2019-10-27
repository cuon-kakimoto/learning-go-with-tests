package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd(){
	// HACK: なんやこの機能は..!godocに例として追加される。。。
	// ex. godoc -http=:6060
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}