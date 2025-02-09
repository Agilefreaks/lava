package processmanager

import (
	"os"
	"os/exec"
	"strings"

	lvutil "github.com/lavanet/lava/ecosystem/lavavisor/pkg/util"
	"github.com/lavanet/lava/utils"
)

func CreateLink(binaryPath string) {
	dest, err := findLavaProtocolPath(binaryPath)
	if err != nil {
		utils.LavaFormatFatal("", err)
	}

	createAndVerifySymlink(binaryPath, dest)
	utils.LavaFormatInfo("Symbolic link created successfully.", utils.Attribute{Key: "Linked binary path", Value: dest})
}

func findLavaProtocolPath(binaryPath string) (string, error) {
	out, err := exec.LookPath("lavap")
	if err == nil {
		return strings.TrimSpace(out), nil
	}
	return copyBinaryToSystemPath(binaryPath)
}

func copyBinaryToSystemPath(binaryPath string) (string, error) {
	gobin, err := exec.Command("go", "env", "GOPATH").Output()
	if err != nil {
		return "", utils.LavaFormatError("couldn't determine Go binary path", err)
	}

	goBinPath := strings.TrimSpace(string(gobin)) + "/bin/"
	validateBinaryExecutable(binaryPath)
	removeExistingLink(goBinPath + "lavap")

	err = lvutil.Copy(binaryPath, goBinPath+"lavap")
	if err != nil {
		return "", utils.LavaFormatError("couldn't copy binary to system path", err)
	}

	out, err := exec.LookPath("lavap")
	if err != nil {
		return "", utils.LavaFormatError("couldn't find the binary in the system path", err)
	}
	return strings.TrimSpace(out), nil
}

func validateBinaryExecutable(path string) {
	version, err := exec.Command(path, "version").Output()
	if err != nil {
		utils.LavaFormatFatal("binary is not a valid executable: ", err)
	}
	utils.LavaFormatInfo("Executable binary validated.", utils.Attribute{Key: "version", Value: strings.TrimSpace(string(version))})
}

func removeExistingLink(linkPath string) {
	if _, err := os.Lstat(linkPath); err == nil {
		utils.LavaFormatInfo("Discovered an existing link. Attempting to refresh.")
		if err := os.Remove(linkPath); err != nil {
			utils.LavaFormatFatal("couldn't remove existing link", err)
		}
	} else if !os.IsNotExist(err) {
		utils.LavaFormatFatal("unexpected error when checking for existing link", err)
	}
}

func createAndVerifySymlink(binaryPath, dest string) {
	if _, err := os.Lstat(dest); err == nil {
		if err := os.Remove(dest); err != nil {
			utils.LavaFormatFatal("couldn't remove existing link", err)
		}
	}

	if err := os.Symlink(binaryPath, dest); err != nil {
		utils.LavaFormatFatal("couldn't create symbolic link", err)
	}

	link, err := os.Readlink(dest)
	if err != nil || link != binaryPath {
		utils.LavaFormatFatal("failed to verify symbolic link", err)
	}
}
