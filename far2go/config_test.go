package far2go

import "testing"
import (
	"gopkg.in/yaml.v2"
	"log"
	"fmt"
)

func TestConfig(t *testing.T) {
	var C Options = Options{}

	d, err := yaml.Marshal(&C)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	//err := viper.Unmarshal(&C)
	//if err != nil {
	//	t.Fatalf("unable to decode into struct, %v", err)
	//}
}

func TestReadWriteConfig(t *testing.T) {
	ReadConfig()
	SaveConfig()

}

func TestSaveConfig(t *testing.T) {
	//ReadConfig()
	SaveConfig()
}

