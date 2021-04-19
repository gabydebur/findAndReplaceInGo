package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new, dest string) (occ int, lines []int, err error) {

	srcFile, err := os.Open(src)
	defer srcFile.Close()

	if err != nil {
		return 0, nil, errors.New("Problème à l'ouverture en lecture du fichier source")
	}

	dstFile, err := os.Create(dest)
	if err != nil {
		return 0, nil, errors.New("Problème à l'ouverture en écriture du fichier destination")
	}
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	numeroligne := 1
	for scanner.Scan() {
		ligne := scanner.Text()
		if strings.Contains(ligne, old) {
			nbOcc := strings.Count(ligne, old)
			occ += nbOcc
			lines = append(lines, numeroligne)
			fmt.Printf("%v\n", strings.Replace(ligne, old, new, nbOcc))
			fmt.Fprintln(writer, strings.Replace(ligne, old, new, nbOcc))
		} else {
			fmt.Printf("%v\n", ligne)
			fmt.Fprintln(writer, ligne)
		}
		numeroligne++
	}

	return occ, lines, nil
}

func main() {
	//Première version les fichiers sont en durs et pas passés en paramètre de l'application.

	nbOcc, line, err := FindReplaceFile("go.txt", "go", "python", "resultat.txt")
	if err != nil {
		fmt.Printf("Erreur : %v \n", err)
		return
	}

	fmt.Println("***********Summary***********")
	fmt.Printf("Nombre d'occurence de go : %v\nNombre de ligne : %v\nLignes : %v\n", nbOcc, len(line), line)
	fmt.Println("********end of Summary********")

}
