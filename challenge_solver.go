/* Rot Cipher with GoLang

Author:: Abdulconsole

This Rotation is different from the traditional or usual rot cipher as it uses all the ascii characters, instead of just the 26 alphabet characters. 

sololearn is my university.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
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
	ws := NewWord{}

	var num int
	fmt.Println("Enter rot number: ")
	fmt.Scanln(&num)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your string: ")
	entry, _ := reader.ReadString('\n')

	entry = strings.TrimSuffix(entry, "\n")

	ws.rotate(num, entry)
	fmt.Println("Original: ", strings.Join(ws.Entry, ""))
	fmt.Println("Rotated: ", strings.Join(ws.Words, ""))
}
