package main

import (
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	capruk "github.com/satriaprayoga/capruk/framework"
)

func doModel(arg3 string) error {
	err := createModel(arg3)
	if err != nil {
		return err
	}
	err = createRepo(arg3)
	if err != nil {
		return err
	}
	err = createRepoStub(arg3)
	if err != nil {
		return err
	}
	err = createInitStub(arg3)
	if err != nil {
		return err
	}
	return nil
}

func createModel(arg3 string) error {
	data, err := templateFS.ReadFile("templates/models/model.go.txt")
	if err != nil {
		return err
	}
	model := string(data)

	var modelName = arg3
	fileName := capruk.RootPath + "/models/" + strings.ToLower(modelName) + ".go"
	model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
	model = strings.ReplaceAll(model, "$modelname$", strings.ToLower(modelName))
	err = copyDataToFile([]byte(model), fileName)
	if err != nil {
		return err
	}
	return nil

}

func createRepo(arg3 string) error {
	data, err := templateFS.ReadFile("templates/repositories/repo.go.txt")
	if err != nil {
		return err
	}
	repo := string(data)
	var app_name = capruk.Config.AppName
	var modelName = arg3
	fileName := capruk.RootPath + "/repositories/" + "i_" + strings.ToLower(modelName) + "_repo.go"
	repo = strings.ReplaceAll(repo, "$MODELNAME$", strcase.ToCamel(modelName))
	repo = strings.ReplaceAll(repo, "$modelname$", strings.ToLower(modelName))
	repo = strings.ReplaceAll(repo, "$APP_NAME$", strings.ToLower(app_name))
	err = copyDataToFile([]byte(repo), fileName)
	if err != nil {
		return err
	}

	return nil
}

func createRepoStub(arg3 string) error {
	data, err := templateFS.ReadFile("templates/repositories/repo_impl.go.txt")
	if err != nil {
		return err
	}
	repo := string(data)
	var app_name = capruk.Config.AppName
	var modelName = arg3
	fileName := capruk.RootPath + "/repositories/" + strings.ToLower(modelName) + "_repo.go"
	repo = strings.ReplaceAll(repo, "$MODELNAME$", strcase.ToCamel(modelName))
	repo = strings.ReplaceAll(repo, "$modelname$", strings.ToLower(modelName))
	repo = strings.ReplaceAll(repo, "$APP_NAME$", strings.ToLower(app_name))
	err = copyDataToFile([]byte(repo), fileName)
	if err != nil {
		return err
	}

	return nil
}

func createInitStub(arg3 string) error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := rootPath + "/repositories/repo.go"
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	initScript := string(data)
	var modelName = arg3
	decstr := "//DECREPO//\n" + "var " + "Repo" + strcase.ToCamel(modelName) + " *" + strcase.ToCamel(modelName) + "Repo"
	initstr := "//INITREPO//\n" + "\tRepo" + strcase.ToCamel(modelName) + "= New" + strcase.ToCamel(modelName) + "Repo()"
	initScript = strings.ReplaceAll(initScript, "//DECREPO//", decstr)
	initScript = strings.ReplaceAll(initScript, "//INITREPO//", initstr)
	err = copyDataToFile([]byte(initScript), filename)
	if err != nil {
		return err
	}
	//fmt.Println(initScript)

	return nil

}
