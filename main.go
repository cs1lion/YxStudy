package main

import "fmt"
//import "math"
//import "strconv"
func main(){
    ciperText:= "CSOITEUIWUIZNSROCNKFD"
    keyword:="GOLANG"
    message:=""
    keyIndex:=0
    for i:=0;i<len(ciperText);i++{
        //A=0,B=1....Z=25
        c:=ciperText[i]-'A'
        k:=keyword[keyIndex]-'A'
        //加密字母-关键字(密码)字母
        c = (c-k+26)%26+'A'
        message+=string(c)
        keyIndex++
        keyIndex%=len(keyword)
    }
    fmt.Println(message)
}