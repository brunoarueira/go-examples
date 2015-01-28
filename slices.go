package main

import "fmt"

func main() {
        s := make([]string, 3)
        fmt.Println("empty: ", s)

        s[0] = "a"
        s[1] = "b"
        s[2] = "c"
        fmt.Println("set:", s)
        fmt.Println("get:", s[2])

        fmt.Println("length: ", len(s))

        s = append(s, "d")
        s = append(s, "e", "f")
        fmt.Println("append:", s)

        c := make([]string, len(s))
        copy(c, s)
        fmt.Println("copy:", c)

        l := s[2:5]
        fmt.Println("s[2:5]:", l)

        l = s[:5]
        fmt.Println("s[:5]:", l)

        l = s[2:]
        fmt.Println("s[2:]:", l)

        t := []string{"g", "h", "i"}
        fmt.Println("declare:", t)

        twoD := make([][]int, 3)
        for i := 0; i < 3; i++ {
                innerLen := i + 1
                twoD[i] = make([]int, innerLen)

                for j := 0; j < innerLen; j++ {
                        twoD[i][j] = i + j
                }
        }
        fmt.Println("two dimensional:", twoD)
}
