package exiftool

import (
	"os/exec"
	"syscall"
)

var readyToken = []byte("{ready}\r\n")

const writeMetadataSuccessToken = "image files updated\r\n"

var exiftoolBinary = "exiftool.exe"

func hideConsoleWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
}
