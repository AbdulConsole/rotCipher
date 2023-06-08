/* Rot Cipher with GoLang

Author:: Abdulconsole

This Rotation Cipher implementation using the traditional or usual 26 alphabet characters. 

Range with: entry >= "a" && entry <= "z"

Formula: (num - num(a)) % 26

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

func (ptr *NewWord) rotate(num int, entry string) {
    for _, v := range entry {
        if v >= 'a' && v <= 'z' {
            enc := 'a' + ((int(v)-'a')+num)%26
            ptr.Entry = append(ptr.Entry, string(v))
            ptr.Words = append(ptr.Words, string(enc))
        } else {
            ptr.Entry = append(ptr.Entry, string(v))
            ptr.Words = append(ptr.Words, string(v))
        }
    }
}

func main() {
    ws := NewWord{}

    var num int
    fmt.Println("Enter rot number: ")
    fmt.Scanln(&num)
    var entry string
    fmt.Println("Enter your string: ")
    fmt.Scanln(&entry)
    
    ws.rotate(num, entry)
    fmt.Println("Original: ",strings.Join(ws.Entry, ""))
    fmt.Println("Rotated: ",strings.Join(ws.Words, ""))
}
