package gotour

import (
	"bytes"
	"io"
	"os"
)

// StdoutSpy -- enable spying of stdout
func StdoutSpy(testFunc func()) string {
	realStdout := os.Stdout
	readFile, writeFile, err := os.Pipe()
	if err != nil {
		return ""
	}
	
	os.Stdout = writeFile

	testFunc()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, readFile)
		outC <- buf.String()
	}()

	writeFile.Close()
	os.Stdout = realStdout // restoring the real stdout
	out := <-outC

	return out
}