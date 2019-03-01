package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/atotto/clipboard"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the string to be cut: ")
    text, _ := reader.ReadString('\n')
    
    text = strings.TrimSuffix(text, "\n")

    slicesOfFour := createSlices(text, 4)

    setupTtyToReadChars()

    var b []byte = make([]byte, 1)
    for _, chunk := range slicesOfFour {
        fmt.Println(chunk)
        fmt.Print("Would you like to (c)opy (s)kip or (e)xit [c/s/e]")
        os.Stdin.Read(b)
        s := string(b)

        fmt.Println("") 
        
        switch s {
          case "c":
            clipboard.WriteAll(chunk)
          case "s":
            continue
          case "e":
            os.Exit(0)
        }
    }
}

func createSlices (text string, n int) []string {
    inputRunes := []rune(text)
  
    slicesOfFour := []string{}
    substr := ""
    for i, r := range inputRunes {
        substr = substr + string(r)
        if (i + 1) % n == 0 {
          slicesOfFour = append(slicesOfFour, substr)
          substr = ""
        } else if (i + 1) == len(inputRunes) {
          slicesOfFour = append(slicesOfFour, substr)
        }
    }

    return slicesOfFour
}

func setupTtyToReadChars () {
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
