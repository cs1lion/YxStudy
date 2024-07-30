package main

import "fmt"
import "strings"

func countwords(text string)map[string]int{
    words:=strings.Fields(strings.ToLower(text))
    frequency:=make(map[string]int,len(words))
    for _,word:=range words{
        word = strings.Trim(word,`.,"-;`)
        frequency[word]++
    }
    return frequency
}

func main(){
    text:=`It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of Darkness, it was the spring of hope, it was the winter of despair, we had everything before us, we had nothing before us, we were all going direct to Heaven, we were all going the other way – in short, the period was so far like the present period, that some of its noisiest authorities insisted on its being received, for good or for evil, in the superlative degree of comparison only.`
    frequency:=countwords(text)
    for word,count:=range frequency{
        if count>1{
            fmt.Printf("%d %v\n",count,word)
        }
    }
}