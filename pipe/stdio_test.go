package pipe_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/stretchr/testify/assert"
)

func TestStdioPipeInputAndOutput(t *testing.T) {
	//-- arrange
	const PROMPT = "Prompt: "
	const INPUT = "input"

	const PRE_INPUT = "prefix"
	const POST_INPUT = "postfix"

	io := pipe.OpenStdio(1, 3, false)
	defer io.Close()

	io.Queue(PROMPT, INPUT)
	io.EndQueue()

	//-- act
	fmt.Println(PRE_INPUT)

	fmt.Print(PROMPT)
	res := readStdin()

	fmt.Println(POST_INPUT)

	//-- assert
	assert.Equal(t, INPUT, res)

	assert.Equal(t, PRE_INPUT, io.ReadLine())
	assert.Contains(t, io.ReadLine(), PROMPT)
	assert.Equal(t, POST_INPUT, io.ReadLine())
}

func TestStdioPipeWithEcho(t *testing.T) {
	//-- arrange
	const PROMPT = "Prompt: "
	const INPUT = "input"

	io := pipe.OpenStdio(1, 1, true)
	defer io.Close()

	io.Queue(PROMPT, INPUT)
	io.EndQueue()

	//-- act
	fmt.Print(PROMPT)
	res := readStdin()

	//-- assert
	assert.Equal(t, INPUT, res)
	assert.Equal(t, fmt.Sprintf("%s%v", PROMPT, INPUT), io.ReadLine())
}
