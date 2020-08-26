package main

import (
	"fmt"
	"slam-device-emulator/controllers"
	"slam-device-emulator/datastore"
	"time"
)

func main() {

	cfg := datastore.ReadConfigFile()

	for idx, equipment := range cfg.Equipments {
		device := controllers.GetDeviceConfig(equipment, cfg.Server.Host)
		fmt.Println("Device:\t", idx, "\n", device)
		datastore.SaveTokensFile(device.Token)
	}

	tokens := datastore.ReadTokens()

	for i := 0; i < 10; i++ {
		for _, token := range tokens.Tokens {
			fmt.Println(token)
			controllers.RegistryMeasure(token, cfg.Server.Host)
		}
		time.Sleep(1 * time.Second)
	}
}
