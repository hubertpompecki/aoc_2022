package input

import (
	"bufio"
	"log"
	"os"
)

func Read(fileName string) chan string {
    lines := make(chan string)

    readFile, err := os.Open(fileName)

    if err != nil {
        log.Fatalf("Couldn't read file: %v", err.Error())
    }

    fileScanner := bufio.NewScanner(readFile)

    go func(){
        for fileScanner.Scan() {
            lines <- fileScanner.Text()
        }
        close(lines)
    }()

    return lines
}
