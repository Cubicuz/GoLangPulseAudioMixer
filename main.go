package main

import (
	"fmt"

	"github.com/Cubicuz/GoLangPulseAudioMixer/gopainterface"
	"github.com/Cubicuz/GoLangPulseAudioMixer/gopamixer"
)

func main() {

	err := gopainterface.Initialize()

	if err != nil {
		fmt.Println(err)
		gopainterface.Deinitialize()
		return
	}

	success, err := gopainterface.Connect()

	if success {
		gopamixer.Somestuff()
	} else {
		if err != nil {
			fmt.Println(err)
		}
	}

	gopainterface.Deinitialize()

}
