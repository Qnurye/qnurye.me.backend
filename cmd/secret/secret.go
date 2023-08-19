package main

import (
	"flag"
	"fmt"
	"qnurye/qnurye.me/pkg/config"
	"qnurye/qnurye.me/pkg/secret"
)

func main() {
	save := flag.Bool("save", false, "Save the key to config file")
	flag.Parse()

	secrets := secret.Generate()
	fmt.Printf(`Jwt secret generated :%s
Aes key for database generated: %s
Aes iv for database generated: %s
Aes key for client generated: %s
Aes iv for client generated: %s
`, secrets.Jwt, secrets.Database.Key, secrets.Database.IV, secrets.Client.Key, secrets.Client.IV)

	if *save {
		config.Set("secret.jwt", secrets.Jwt)
		config.Set("secret.database_Key", secrets.Database.Key)
		config.Set("secret.database_Iv", secrets.Database.IV)
		config.Set("secret.client_Key", secrets.Client.Key)
		config.Set("secret.client_Iv", secrets.Client.IV)
		fp := config.GetFilePath()
		println("Saved secrets to", fp)
	}
}
