package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

    var b []byte = make([]byte, 1)
    for b[0] != rune('q') {
        os.Stdin.Read(b)

        switch rune(b) {
                case 97: //Key A {
                        diw++
                        //west player
                        }
                case 119: //Key W {
                        dih++
                        //up player
                        }
                case 115: //Key S {
                        dih--
                        //up player
                        }
                case 100: //Key D {
                        diw--
                        //up player
                        }
                case 101: //Key E {
                        // use
                        }
                case 104: //Key H {
                        //Attack
                        }
                default:
        }


        fmt.Println("I got the byte", b, "("+string(b)+")")



    }
    exec.Command("stty", "-F", "/dev/tty", "echo").Run()

}
