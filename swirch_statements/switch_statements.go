package main

// Uncoment and run one by one top to bottom
func main() {

	// // No switch used
	// argument := os.Args[1] //Collects argument from CLI input
	// if argument == "agua" {
	// 	fmt.Println("mole, pedra dura, tanto bate até que fura")
	// } else if argument == "cavalo" {
	// 	fmt.Println("dado não se olha os dentes")
	// } else if argument == "pimenta" {
	// 	fmt.Println("no * dos outros é refresco")
	// } else {
	// 	fmt.Println("Sem palavras")
	// }

	// // Converting to switch statements
	// argument := os.Args[1]
	// switch argument {
	// case "agua":
	// 	fmt.Println("mole, pedra dura, tanto bate até que fura")
	// case "cavalo":
	// 	fmt.Println("dado não se olha os dentes")
	// case "pimenta":
	// 	fmt.Println("no * dos outros é refresco")
	// default:
	// 	fmt.Println("Sem palavras")
	// }

	// // Using fallthrough and other
	// president := os.Args[1]
	// switch president {
	// case "Trump":
	// 	fmt.Println("Incompetente")
	// 	fallthrough // Will execute the case below as well
	// case "Bolsonaro", "Olavo": // Will execute if either Bolsonaro or Olavo are passed as arguments
	// 	fmt.Println("Burro")
	// case "Obama":
	// 	fmt.Println("honorable & able")
	// case "Michele": // Will run and nothing will happen
	// }
	// // Output:
	// // go run swirch_statements/switch_statements.go Trump
	// // Incompetente
	// // Burro

	// // Using variables inside case
	// greatCoder := os.Args[1]
	// lovecoder := "Ada"
	// toughtcoder := "Grace"
	// funcoder := "Carol"
	// switch greatCoder {
	// case lovecoder:
	// 	fmt.Println("Lovelace")
	// case toughtcoder:
	// 	fmt.Println("Hoper")
	// case funcoder:
	// 	fmt.Println("Shaw")
	// }

	// // Output:
	// // go run swirch_statements/switch_statements.go Carol
	// // Shaw

	// // We can initialize switch by declaring a variable in the switch line
	// // and use functions at the same time ex: lenght
	// givenWord := os.Args[1]
	// secretWord := "paralelepipedo"
	// switch wordLenght := len(givenWord); { //Declaring wordLenght here, together with switch
	// case wordLenght != 14:
	// 	fmt.Println("Wrong, the secret word has 14 characters")
	// case wordLenght == 14 && givenWord != secretWord: //Lenght is correct, but word is wrong
	// 	fmt.Println("Good Try")
	// case givenWord == secretWord:
	// 	fmt.Println("Correct!")
	// }
}
