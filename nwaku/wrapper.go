package nwaku

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func StopNode() {
	// Since we have reference to same process we can also use cmd.Process.Kill()
	strb, _ := ioutil.ReadFile("wakunode2.lock")
	command := exec.Command("kill", string(strb))
	command.Start()
	log.Printf("stopping wakunode2 process %s", string(strb))
}

func StartNode(nodeStopped chan bool) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for signal
	go func() {
		sig := <-sigs
		log.Printf("received %s", sig)
		StopNode()
		done <- true
	}()

	// TODO Use PATH or so here instead...
	// XXX Hardcode store node here, build up string instead
	// TODO Build up command here
	cmd := exec.Command("./bin/wakunode2", "--storenode=/ip4/188.166.135.145/tcp/30303/p2p/16Uiu2HAmL5okWopX7NqZWBUKVqW8iUxCEmd5GMHLVPwCgzYzQv3e")

    outfile, err := os.Create("./wakunode2.log")
	if err != nil {
		panic(err)
    }
    defer outfile.Close()

	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

    log.Printf("wakunode2 start, [PID] %d running...\n", cmd.Process.Pid)
    ioutil.WriteFile("wakunode2.lock", []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0666)

    <-done
    log.Printf("exiting")

	// XXX Move to StopNode function? Flow a bit odd
	nodeStopped <- true
}
