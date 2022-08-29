package main

import (
  "fmt"
  "os"
  "net/http"
  "time"
  "io"
  "io/ioutil"
  "bufio"
  "strings"
  "strconv"
)

const numCheck = 3
const delay = 5

func main()  {
  displaySaudation()

  for {
    displayMenu()
    fmt.Println("")
    opt := getOption()

    switch opt {
      case 1:
        fmt.Println("")
        fmt.Println("Monitoring...")
        checkUrl()

      case 2:
        fmt.Println("")
        fmt.Println("Showing Logs...")
        showLogs()

      case 0:
        fmt.Println("")
        fmt.Println("Exiting...")
        os.Exit(0)

      default:
        fmt.Println("Invalid Option!")
        os.Exit(-1)
    }
  }
}

func displaySaudation() {
  ver := 1.1

  fmt.Println("Version:", ver)
  fmt.Println("Hello!")
}

func displayMenu() {
  fmt.Println("1 - Start Monitoring")
  fmt.Println("2 - Show Logs")
  fmt.Println("0 - Exit")
}

func getOption() int {
  var opt int

  fmt.Scan(&opt)

  return opt
}

func checkUrl() {
  urls := readUrlFile()

  for i := 0; i < numCheck; i++ {
    for _, url := range urls {
      fmt.Println("Checking:", url)
      checkStatusCode(url)
      fmt.Println("")
    }
    time.Sleep(delay * time.Second)
  }
}

func readUrlFile() []string {
  var urls []string

  file, err := os.Open("./url.txt")
  if err != nil {
    fmt.Println("Error:", err)
  }

  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadString('\n')
    line = strings.TrimSpace(line)

    if err == io.EOF {
        break
    }

    urls = append(urls, line)
  }

  file.Close()

  return urls
}

func checkStatusCode(url string) {
  res, err := http.Get(url)

  if err != nil {
    fmt.Println("Error:", err)
  }

  if res.StatusCode == 200 {
    fmt.Println("Up and Running!")
    registerLogs(url, true)
  } else {
      fmt.Println("Error - Status Code:", res.StatusCode)
      registerLogs(url, false)
  } 
}

func registerLogs(url string, status bool) {
  file, err := os.OpenFile("./log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND , 0666)
  if err != nil {
    fmt.Println(err)
  }

  file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")
  
  file.Close()
}

func showLogs() {
  
  file, err := ioutil.ReadFile("log.txt")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(file))
}
