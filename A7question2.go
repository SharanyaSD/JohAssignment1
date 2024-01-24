//Sharanya Datrange

package main
import  (
	"fmt"
	"bufio"
	"os"
	"sync"
	"runtime"
)
func main() {
	fmt.Println("Enter a string: ")
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input:=scanner.Text()

	var wg sync.WaitGroup

	wg.Add(1)

	go reverseString(input,&wg)

	wg.Wait()

}
func reverseString(input string, wg *sync.WaitGroup) {
	defer wg.Done()
	reversedStr:=""
	for i:=len(input)-1;i>=0;i-- {
		reversedStr+=string(input[i])
	}
	fmt.Println("Reversed string: ",reversedStr)
	fmt.Println("No. of Go routines launched: ", runtime.NumGoroutine())
}