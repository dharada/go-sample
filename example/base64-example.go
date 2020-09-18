package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "b05Kn3QB2seGkJaqnxBp:o2T1kbWpQp28l-7zjMDsyw"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

}
