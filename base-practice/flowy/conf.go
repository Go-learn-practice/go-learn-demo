package flowy

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var overrideUserData = ""

func OverrideUserData(ou string) {
	overrideUserData = ou
}

func GetUserDataDir() (string, error) {
	if overrideUserData != "" {
		return overrideUserData, nil
	}
	currentUser, err := user.Current()
	fmt.Printf("currentUser的值：%v\n", currentUser)

	if err != nil {
		log.Printf("user.Current Failed ,err = %s", err.Error())
		return "", err
	}

	var userDataDir string

	appDataPath := os.Getenv("LOCALAPPDATA")
	fmt.Printf("appDataPath的值：%s\n", appDataPath)

	if appDataPath == "" {

		userDataDir = filepath.Join(currentUser.HomeDir, "AppData", "Local")
	} else {
		userDataDir = appDataPath
	}

	const AppName = "flowy-aipc"

	var p = filepath.Join(userDataDir, AppName)

	// 使用 os.Stat 获取文件信息
	_, err = os.Stat(p)

	// os.Mkdir 用于创建单个目录 当文件不存在时
	if err != nil {
		return p, os.Mkdir(p, 0755)
	}

	return p, nil
}
