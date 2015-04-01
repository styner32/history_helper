package main

import (
	"bufio"
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
	"log"
	"os"
	"os/user"
)

const usage = `Usage:
  hh import           # import your bash history`

func main() {
	numberOfArguments := len(os.Args)

	if numberOfArguments < 2 {
		log.Fatal(usage)
	}

	var bashHistoryPath string
	switch os.Args[1] {
	case "import":
		if len(os.Args) > 2 {
			bashHistoryPath = os.Args[2]
		} else {
			bashHistoryPath = defaultBashHistoryPath()
		}
		ImportBashHistory(bashHistoryPath)
	default:
		log.Fatalf("Could not find command '%s'", os.Args[1])
	}
}

func ImportBashHistory(filePath string) {
	fmt.Printf("Importing from %s\n", filePath)
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	if err != nil {
		log.Fatalf("error opening history file: %s", err)
	}

	defer file.Close()

	_, err = os.Stat(defaultWorkDirectory())
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(defaultWorkDirectory(), 0777)
	}

	client, err := sqlite3.Open(defaultSqlitePath())
	if err != nil {
		log.Fatalf("failed to open sqlite file: %s", err)
	}

	err = client.Exec("CREATE TABLE IF NOT EXISTS histories (command TEXT);")
	if err != nil {
		log.Fatalf("error creating history table: %s", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sql := fmt.Sprintf("INSERT INTO histories VALUES('%s');", scanner.Text())
		log.Println(sql)
		client.Exec(sql)
	}
}

func defaultBashHistoryPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.bash_history", user.HomeDir)
}

func defaultSqlitePath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.hh/history.db", user.HomeDir)
}

func defaultWorkDirectory() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.hh/", user.HomeDir)
}
