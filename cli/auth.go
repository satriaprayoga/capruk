package main

import (
	"strings"

	capruk "github.com/satriaprayoga/capruk/framework"
)

func doAuth() error {
	data, err := templateFS.ReadFile("templates/auth/auth.go.txt")
	if err != nil {
		return err
	}
	model := string(data)

	fileName := capruk.RootPath + "/models/" + "user" + ".go"
	err = copyDataToFile([]byte(model), fileName)
	if err != nil {
		return err
	}
	err = createAuthRepo("user")
	if err != nil {
		return err
	}
	err = createAuthRepoStub("user")
	if err != nil {
		return err
	}
	return nil
}

func createAuthRepo(arg3 string) error {
	data, err := templateFS.ReadFile("templates/auth/auth_repo.go.txt")
	if err != nil {
		return err
	}
	repo := string(data)
	var app_name = capruk.Config.AppName
	var modelName = "user"
	fileName := capruk.RootPath + "/repositories/" + "i_" + strings.ToLower(modelName) + "_repo.go"
	repo = strings.ReplaceAll(repo, "$APP_NAME$", strings.ToLower(app_name))
	err = copyDataToFile([]byte(repo), fileName)
	if err != nil {
		return err
	}

	return nil
}

func createAuthRepoStub(arg3 string) error {
	data, err := templateFS.ReadFile("templates/auth/auth_repo_impl.go.txt")
	if err != nil {
		return err
	}
	repo := string(data)
	var app_name = capruk.Config.AppName
	var modelName = arg3
	fileName := capruk.RootPath + "/repositories/" + strings.ToLower(modelName) + "_repo.go"
	repo = strings.ReplaceAll(repo, "$APP_NAME$", strings.ToLower(app_name))
	err = copyDataToFile([]byte(repo), fileName)
	if err != nil {
		return err
	}

	return nil
}
