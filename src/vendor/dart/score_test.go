package dart

import (
	"fmt"
	"testing"
)

func TestScoreToString(t *testing.T)  {
	value := Score{"S180", 180}
	actual := value.String()
	expected := "180 = S180"
	if actual != expected {
		msg := fmt.Sprintf("expected value: '%s'  Actual value:'%s'", expected, actual)
		t.Error(msg)
	}
}
