/*package main

import "fmt"
import "strings"

func main() {
    x := "chars@arefun"

    i := strings.Index(x, "@")
    fmt.Println("Index: ", i)
    chars := x[:i]
    arefun := x[i+1:]
    fmt.Println(chars)
    fmt.Println(arefun)

}

package main

import "fmt"

func main() {

   str := "Hello"
   for i, elem := range str {
       fmt.Println(i, str[i], elem)
   }

   for elem := range str {
       fmt.Println(elem)
   }
}
*/
package main
import "fmt"

   func main() {
        var word string = "ZbjTS"

       // P R I N T
       fmt.Println(word)
       yo := string([]rune(word)[0])
       fmt.Println(yo)

       //I N D E X
       x :=0
       for x < len(word){
           yo := string([]rune(word)[x])
           fmt.Println(yo)
           x+=1
       }

}
