package exiftool

var readyToken = []byte("{ready}\n")

const writeMetadataSuccessToken = "image files updated\n"

var exiftoolBinary = "exiftool"

func hideConsoleWindow(cmd *exec.Cmd) {
	//cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
}
