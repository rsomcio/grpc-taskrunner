package cmd

import (
	"fmt"
	"time"

	"os"
	"os/exec"
)

func Run(c string, a []string) {

	cmd := exec.Command(c, a...)
	output, err := cmd.Output()
	if err != nil {
        panic(err)
	}

    fn := fmt.Sprintf("%s-%s.txt", c, time.Now().Format("20060102_1545"))
    f, err := os.Create(fn)

    defer f.Close()


    f.Write(output)

}
