package main

import (
  "bufio"
  "fmt"
  "io"
  "os"

  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)

func To_utf8(inStream io.Reader, outStream io.Writer) error {
  scanner := bufio.NewScanner(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder()))
  list := make([]string, 0)
  for scanner.Scan() {
    list = append(list, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    return err
  }
  writer := bufio.NewWriter(outStream)
  for _, line := range list {
    var err error
    _, err = fmt.Fprintln(writer, line)
    if err != nil {
      return err
    }
  }
  return writer.Flush()
}

func main() {
  read_file, _ := os.OpenFile("./AddrBook.csv", os.O_RDONLY, 0600)
  write_file, _ := os.OpenFile("./out", os.O_WRONLY|os.O_CREATE, 0600)
  To_utf8(read_file, write_file)
}
