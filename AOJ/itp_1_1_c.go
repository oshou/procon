 package main

 import (
	 "bufio"
	 "os"
	 "math"
	 "strconv"
	 "fmt"
 )

 var sc = bufio.NewScanner(os.Stdin)

 func init(){
	 sc.Split(bufio.ScanWords)
	 sc.Buffer([]byte{},math.MaxInt64)
 }

 func readInt()int{
	 sc.Scan()
	 r,_:=strconv.Atoi(sc.Text())
	 return r
 }

 func main(){
	 a,b := readInt(),readInt()
	 fmt.Println(a*b,2*a+2*b)
 }
