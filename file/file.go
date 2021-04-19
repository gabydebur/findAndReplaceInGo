package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {

	srcFile, err := os.Open(src)
	defer srcFile.Close()

	if err != nil {
		return 0, nil, errors.New("Problème à l'ouverture en lecture du fichier source")
	}

	scanner := bufio.NewScanner(srcFile)
	numeroligne := 1
	for scanner.Scan() {
		ligne := scanner.Text()
		if strings.Contains(ligne, old) {
			nbOcc := strings.Count(ligne, old)
			occ += nbOcc
			lines = append(lines, numeroligne)
			fmt.Printf("%v\n", strings.Replace(ligne, old, new, nbOcc))
		} else {
			fmt.Printf("%v\n", ligne)
		}
		numeroligne++
	}

	return occ, lines, nil
}
