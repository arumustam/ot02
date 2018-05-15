package encoder

import (
  "bufio"
  "fmt"
  "io"
  "os"

  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)

func ToUTF8(inStream io.Reader) []string{
  var list []string
  scanner := bufio.NewScanner(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder()))
  for scanner.Scan() {
    list = append(list, scanner.Text())
  }
 return list
}

func main() {
  read_file, _ := os.OpenFile("./AddrBook.csv", os.O_RDONLY, 0600)
  //write_file, _ := os.OpenFile("./out", os.O_WRONLY|os.O_CREATE, 0600)
  list := ToUTF8(read_file)
  fmt.Println(list[8])
}
