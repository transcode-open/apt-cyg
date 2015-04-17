package main
import (
  "compress/bzip2"
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)

func wget(url string) {
  tokens := strings.Split(url, "/")
  fileName := tokens[len(tokens)-1]
  fmt.Println("Downloading", url, "to", fileName)
  output, err := os.Create(fileName)
  if err != nil {
    return
  }
  response, err := http.Get(url)
  if err != nil {
    return
  }
  n, err := io.Copy(output, response.Body)
  if err != nil {
    return
  }
  fmt.Println(n, "bytes downloaded.")
}

func bunzip2(alpha string) {
  bravo, err := os.Open(alpha)
  if err != nil {
    return
  }
  charlie := bzip2.NewReader(bravo)
  delta, err := os.Create("setup.ini")
  if err != nil {
    return
  }
  io.Copy(delta, charlie)
}

func foxtrot() {
  // create release folder
  os.MkdirAll("mirror/x86_64/release", 0)  
  // cd
  os.Chdir("mirror/x86_64")
  // download
  wget("http://cygwin.osuosl.org/x86_64/setup.bz2")
  // extract
  bunzip2("setup.bz2")
}

func main() {  
  foxtrot()
  // parse ini
  // category: Base
}
