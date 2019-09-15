package main

import "fmt"

func m1(d []string) []string {
  
  fmt.Println(d)

  return d[2:4]
}
// START OMIT
func main() {
    s := make([]string, 3) // slice of strings of size 3
    s[0] = "zero" // set just like an array
    s[1] = "one"
    s[2] = "two"
    s = append(s, "three") // HL
    s = append(s, "four", "five") // HL
   
    var x=m1(s)
    fmt.Println(x)
    
    x[0]="yo"
    fmt.Println(s)
}
