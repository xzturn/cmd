// xzturn package, wrapper of command line shell
// 
// Author: tengming.sp
//
package cmd

import (
    "bytes"
    "io"
    "os"
    "os/exec"
)

////////////////////////////////////////////////////////////////////////////////

func RunCmd(name string, dir string, args ...string) error {
    cmd := exec.Command(name, args ...)
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func RunCmdEx(name string, dir string, prefix string, args ...string) error {
    cmd := exec.Command(name, args ...)
    cmd.Dir = dir
    var out bytes.Buffer
    out.WriteString(prefix)
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    out.WriteTo(os.Stdout)
    return err
}

func RunCmdByParams(name string, dir string, params []string) error {
    cmd := exec.Command(name)
    cmd.Args = append(cmd.Args, params...)
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func RunCmdSpecStdout(name string, dir string, out io.Writer, args ...string) error {
    cmd := exec.Command(name, args ...)
    cmd.Dir = dir
    cmd.Stdout = out
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

// NOTE: the usr should have SSH KEY on host machine for cmdstr execution without password
func RemoteRun(usr, host, cmdstr string) error {
    return RunCmd("ssh", ".", usr + "@" + host, cmdstr)
}
