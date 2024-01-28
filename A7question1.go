//Sharanya Datrange

package main

import "fmt"

func main() {
	aliceChannel:=make(chan string)
	bobChannel:=make(chan string)
	channelSignal:=make(chan struct{})
	
	go func() {
		string1:="helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

		string2 :=""
		for _,symbol :=range string1 {

			if symbol=='$' {
				aliceChannel<-string2
				string2=""
			} else if symbol=='#' {
				bobChannel<-string2
				string2=""
			} else if symbol=='^' {
				channelSignal<-struct{}{}
			} else {
				string2=string2+string(symbol)
			}
		}

	}()
	
	output:=""
	for {
		select {
		case test:=<-aliceChannel:
			output=output+"alice: "+test+","
		case test:=<-bobChannel:
			output=output+"bob:"+test+","
		case <-channelSignal:
			fmt.Println(output)
		return
		}
	}

}
