package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// read configuration file
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to read config file : %s", err))
	}
	ifname := viper.GetString("opendns.interface")
	user := viper.GetString("opendns.user")
	pass := viper.GetString("opendns.password")
	url := fmt.Sprintf("https://updates.opendns.com/nic/update?hostname=%s", ifname)
	//fmt.Printf("u:%s, p:%s, l:%s", user, pass, url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//fmt.Println("ok", err)
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	resp.Body.Close()
}
