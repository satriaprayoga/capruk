package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/satriaprayoga/capruk/utils"
)

var appUrl string

func doNew(appName string) error {
	if appName == "" {
		return errors.New("new requires an application names")
	}
	appName = strings.ToLower(appName)
	appUrl = appName
	// sanitize the application name (convert url to single name)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}
	log.Println("App Name is", appName)
	//git clone the skeleton application
	color.Green("\tCloning repository...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "https://github.com/satriaprayoga/capruk-app.git", //    "git@github.com:satriaprayoga/capruk-app.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		return err
	}

	//remove .git directory
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		return err
	}
	//create a config file
	color.Yellow("\tCreating config.json file...")
	data, err := templateFS.ReadFile("templates/config.json.txt")
	if err != nil {
		return err
	}

	config := string(data)
	config = strings.ReplaceAll(config, "${APP_NAME}", appName)
	config = strings.ReplaceAll(config, "${KEY}", utils.GenerateString(32))

	err = copyDataToFile([]byte(config), fmt.Sprintf("./%s/config.json", appName))
	if err != nil {
		return err
	}

	//create a makefile
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.win", appName))
		if err != nil {
			return err
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			return err
		}

	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			return err
		}
		defer source.Close()
		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			return err
		}
	}

	_ = os.Remove("./" + appName + "/Makefile.mac")
	_ = os.Remove("./" + appName + "/Makefile.win")

	//update the go.mod file
	color.Yellow("\tCreating go.mod file...")
	_ = os.Remove("./" + appName + "/go.mod")

	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		return err
	}

	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appUrl)

	err = copyDataToFile([]byte(mod), "./"+appName+"/go.mod")
	if err != nil {
		return err
	}

	//update existing .go files with correct name/imports
	color.Yellow("\tUpdating source files....")
	os.Chdir("./" + appName)
	err = updateSource()
	if err != nil {
		return err
	}

	//run go mod tidy in the project directory
	color.Yellow("\tRunning go mod tidy...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Start()
	if err != nil {
		return err
	}
	color.Green("Done building " + appUrl)
	return nil

}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	//error check first
	if err != nil {
		return err
	}
	//check if current file is a directory
	if fi.IsDir() {
		return nil
	}
	//check go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}
	if matched {
		read, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		newContents := strings.Replace(string(read), "myapp", appUrl, -1)
		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateSource() error {
	//recursive walk entire project folders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		return err
	}
	return nil
}
