package Seal

import (
	"fmt"
	"os"
	"os/exec"
)

//Sella :
func Sella(path string) string {
	//Capturamos directorio actual
	dirOrigin, err := os.Getwd()

	//Nos movemos al nuevo directorio
	os.Chdir("./" + path)

	//Ejecutamos en el nuevo directorio
	cmd, err := exec.Command("/bin/bash", "./Timbra.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}

	//Capturamos la salida del CMD
	output := string(cmd)

	//Regresamos al directorio original
	os.Chdir(dirOrigin)

	//Devolvemos la salida del cmd
	return output

}
