package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "os/user"
  "code.google.com/p/go-sqlite/go1/sqlite3"
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
		log.Fatalf("error opening table file: %s", err)
	}
	defer file.Close()

  client, _ := sqlite3.Open(defaultSqlitePath())
  sql := "SELECT name FROM sqlite_master WHERE type='table' AND name='histories';"

  row := make(sqlite3.RowMap)
  for s, err := c.Query(sql); err == nil; err = s.Next() {
  	var rowid int64
	  s.Scan(&rowid, row)     // Assigns 1st column to rowid, the rest to row
	  fmt.Println(rowid, row) // Prints "1 map[a:1 b:demo c:<nil>]"
  }
}

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}

func defaultBashHistoryPath() string {
  user, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }
  return fmt.Sprintf("%s/.bash_history", user.HomeDir)
}

func defaultSqlitePath() string {
  user, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }
  return fmt.Sprintf("%s/.hh/history.db", user.HomeDir)
}
