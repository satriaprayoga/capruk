package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	capruk "github.com/satriaprayoga/capruk/framework"
	"github.com/satriaprayoga/capruk/utils"
)

func doController(arg3 string) error {
	err := createController(arg3)
	if err != nil {
		return err
	}
	err = createSetup(arg3)
	if err != nil {
		return err
	}
	return nil
}

func createController(arg3 string) error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	var contDir = rootPath + "/controllers/" + strings.ToLower(arg3)
	err = utils.CreateDirIfNotExist(contDir)
	if err != nil {
		return err
	}

	data, err := templateFS.ReadFile("templates/controllers/controllers.go.txt")
	if err != nil {
		return err
	}
	cont := string(data)
	var contName = arg3
	fileName := contDir + "/" + strings.ToLower(contName) + ".go"
	cont = strings.ReplaceAll(cont, "$CONTNAME$", strcase.ToCamel(contName))
	cont = strings.ReplaceAll(cont, "$contname$", strings.ToLower(contName))
	err = copyDataToFile([]byte(cont), fileName)
	if err != nil {
		return err
	}
	return nil
}

func createSetup(arg3 string) error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := rootPath + "/routes/route.go"
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	initScript := string(data)
	var contName = arg3
	var app_name = capruk.Config.AppName
	pkgName := fmt.Sprintf("%v/controllers/%v", app_name, strings.ToLower(contName))
	pkgQ := fmt.Sprintf("%q", pkgName)
	impstr := "//IMPORTCONT//\n" + "import\t" + strings.ToLower(contName) + " " + pkgQ
	initstr := "//INITSETUP//\n" + "\t" + strings.ToLower(contName) + "." + "Setup()"
	initScript = strings.ReplaceAll(initScript, "//IMPORTCONT//", impstr)
	initScript = strings.ReplaceAll(initScript, "//INITSETUP//", initstr)
	//fmt.Println(initScript)
	err = copyDataToFile([]byte(initScript), filename)
	if err != nil {
		return err
	}
	return nil
}
