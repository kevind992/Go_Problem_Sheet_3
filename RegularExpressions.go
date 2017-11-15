package main

import(

	"fmt"
	"math/rand"
	"time"
	"regexp"

)

func main(){

	rand.Seed(time.Now().UTC().UnixNano())
	

	var userInputs [5]string
	userInputs[0] = "People say I look like both my mother and father."
	userInputs[1] = "Father was a teacher."
	userInputs[2] = "I was my father's favourite."
	userInputs[3] = "I'm looking forward to the weekend."
	userInputs[4] = "My grandfather was French!"

	for i:=0;i<5;i++{
		response := ElizaResponce(userInputs[i])
		fmt.Println("Eliza: ",response)
		fmt.Println()
	}
}
func ElizaResponce(input string) string{
	
	fmt.Println("User: ",input)
	
	var response [3]string
	response[0] = "I'm not sure what you're trying to say. Could you explain it to me?"
	response[1] = "How does that make you feel?"
	response[2] = "Why do you say that?"

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don't you tell me more about your Father?"
	}

	ranNum := rand.Intn(3)

	return response[ranNum]
		
}