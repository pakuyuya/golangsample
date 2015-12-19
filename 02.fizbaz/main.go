package main

import (
  "os"
  "flag"
  "strconv"
)

const (
  Fiz  = "FIZ"
  Baz  = "BAZ"
)


func main() {
  flag.Parse()
  
  for i := 0; i < flag.NArg(); i++ {
    var arg int
    arg, _ = strconv.Atoi(flag.Arg(i))
    
    if (arg % 3) == 0 {
      os.Stdout.WriteString(Fiz);
    }
    if (arg % 5) == 0 {
      os.Stdout.WriteString(Baz);
    }
     os.Stdout.WriteString(" ");
  }
}