package main

import (
	"flag"
	"log"
	"os/exec"
	"path"
	"time"
)

func main() {

	patchInput := flag.String("patch", "", "The path to the actual patch folder, ie /home/user/WSO2patch/patch1234")
	installInput := flag.String("install", "", "The path to the WSO2 install, ie /opt/wso2")
	noRestartFlag := flag.Bool("norestart", false, "Used if you do not want to restart WSO2 following the patch")
	flag.Parse()

	if *patchInput == "" || *installInput == "" {
		log.Fatal("Both the patch directory and the install directory must be filled")
	}

	patchDir := path.Clean(*patchInput)
	installDir := path.Clean(*installInput)
	patchParentDir := path.Dir(patchDir)
	patchName := path.Base(patchDir)

	err := exec.Command("cp", patchParentDir+"/wso2carbon-version.txt", installDir+"/bin").Run()
	if err != nil {
		log.Fatal("Failed to copy version file...")
	}

	err = exec.Command("cp", "-r", patchDir, installDir+"/repository/components/patches").Run()
	if err != nil {
		log.Fatal("Failed to copy the patch folder...")
	}

	log.Println("Patch copied to install directory")

	if *noRestartFlag {
		log.Print("Restart was skipped")
		log.Print("In order to fully apply the patch please restart your install at your convenience")
		return
	}

	log.Print("Restarting WSO2.....")

	err = exec.Command("sh", installDir+"/bin/wso2server.sh", "restart").Run()
	if err != nil {
		log.Fatal("Failed to restart the WSO2 install")
	}

	log.Print("Restarted!!!")

	// Waiting 5 seconds is completely arbitrary. It just gives some time for WSO2 to actually log the patch
	// Sometimes the patch takes longer than that to show up (sometimes taking 15+ seconds),
	// so the script may occasionally fail just due to lag
	log.Println("Waiting 5 seconds for server to finish patching")
	time.Sleep(5 * time.Second)

	err = exec.Command("grep", patchName, installDir+"/repository/logs/patches.log").Run()
	if err != nil {
		log.Fatal("WARNING: Does not seem to have been applied, you may manually check the patch log to verify")
	}

	log.Println("Success! The patch has been properly applied!")

}
