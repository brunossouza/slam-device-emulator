package main

import (
	"flag"
	"fmt"
	"slam-device-emulator/models"
	"time"

	"slam-device-emulator/datastore"
)

// variaveis de configuração
var (
	configFlags *flag.FlagSet
	host        string
	equipamento string
	devices     []models.Device
)

func init() {
	configFlags = flag.NewFlagSet("config", flag.ExitOnError)
	configFlags.StringVar(&host, "h", "localhost:8080", "ip e porta ou dominio do servidor")
	configFlags.StringVar(&equipamento, "e", "não informado", "equipamento que o consumo será emulado, exemplo televisão")
}

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("expected 'config' or 'run' subcommands")
	// 	os.Exit(1)
	// }

	// switch os.Args[1] {
	// case "config":
	// 	configFlags.Parse(os.Args[2:])

	// 	break
	// }

	// devices = append(devices, controllers.GetDeviceConfig(equipamento, host))
	// for _, dev := range devices {
	// 	fmt.Println("Device:\t", dev)

	// }

	cfg := datastore.ReadConfigFile()

	fmt.Printf("%v\n", cfg)

	datastore.SaveTokensFile(time.Now().String())
	datastore.ReadTokens()
}
