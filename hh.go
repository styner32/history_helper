package main

import (
  "os"
  "fmt"
  "log"
)

const defaultBashHistoryPath = "~/.bash_history"

func main() {
  numberOfArguments := len(os.Args)

  if numberOfArguments < 2 {
    usage := `Usage:
  hh import           # import your bash history`
    log.Fatal(usage)
  }

  bashHistoryPath := defaultBashHistoryPath
  switch os.Args[1] {
  case "import":
    if len(os.Args) > 2 {
      bashHistoryPath = os.Args[2]
    }
    ImportBashHistory(bashHistoryPath)
  default:
    log.Fatalf("Could not find command '%s'", os.Args[1])
  }
}

func ImportBashHistory(filePath string) {
  fmt.Printf("Importing from %s\n", filePath)
}
