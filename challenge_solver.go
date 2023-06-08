/* Rot Cipher with GoLang

Author:: Abdulconsole

This Rotation is different from the traditional or usual rot cipher as it uses all the ascii characters, instead of just the 26 alphabet characters.

*/

package main

import (
    "fmt"
    "strings"
)

type NewWord struct {
    Words []string
    Entry []string
}

func (ptr *NewWord)rotate(num int, entry string) {
    for _, v := range entry {
        enc := (int(v) + num) % 128
        
        ptr.Entry = append(ptr.Entry, string(v))
        ptr.Words = append(ptr.Words, string(enc))
    }
}

func main() {
    var num int
    fmt.Println("Enter rot number: ")
    fmt.Scanln(&num)
    var entry string
    fmt.println("Enter your string: ")
    fmt.Scanln(&entry)
    
    ws := NewWord{}
    
    ws.rotate(num, entry)
    fmt.Println("Original: ",strings.Join(ws.Entry, ""))
    fmt.Println("Rotated: ",strings.Join(ws.Words, ""))
}
