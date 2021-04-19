package main

import (
	"fmt"

	"training.go/findandreplaceingo/file"
)

func main() {
	//Première version les fichiers sont en durs et pas passés en paramètre de l'application.

	nbOcc, line, err := file.FindReplaceFile("go.txt", "go", "python")
	if err != nil {
		fmt.Printf("Erreur : %v \n", err)
		return
	}

	fmt.Println("***********Summary***********")
	fmt.Printf("Nombre d'occurence de go : %v \n Nombre de ligne : %v \n Lignes : %v \n ", nbOcc, len(line), line)
	fmt.Println("********end of Summary********")

}
