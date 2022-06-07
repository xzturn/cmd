package cmd

import (
    "bytes"
    "testing"
)

func TestRunCmd(t *testing.T) {
    if err := RunCmd("ls", ".", "-l", "-t", "-a"); err != nil { t.Error(err) }
    if err := RunCmdByParams("ls", ".", []string{"-l", "-t", "-a"}); err != nil { t.Error(err) }
    var buf bytes.Buffer
    if err := RunCmdSpecStdout("ls", ".", &buf, "-l", "-t", "-a"); err != nil { t.Error(err) }
    t.Log(buf.String())
}
