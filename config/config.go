package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configJSON struct {
	Token string `json:"Token"`
}

var (
	config configJSON
	Token string
)

func ReadConfig() error {
	fmt.Println("Reading config file ...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token

	return nil
}
