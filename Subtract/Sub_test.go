package subtract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallSubhere(t *testing.T) {
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\nThis is message before calling SUBTRACT function the test go file")
	/*Calling max function this is a comment */
	var Subresultvalue int
	Subresultvalue = CallSub(5, 2)
	//To log any info while test is running the test has to be run in command prompt with go test -v command to show log
	t.Logf("\nThis is message from t.Logf and Subtraction returned is  %d", Subresultvalue)
	
	assert.Equal(t, 3, Subresultvalue)
	t.Log("Testpassed after the assert logic of max function in test go file.")

}
