package main



import (

	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


var messageChan chan string = make(chan string)

func main() {
		
	//calculateGivenCipherText("2081 2182")
	
	

	cipherText := encryptMessageRSA("stop this madness", 43, 59, 13)
	fmt.Println(cipherText)
	//cipherText = strings.Replace(cipherText, " ", "", -1)
	//fmt.Println(cipherText)

	//return


	sliceOfPrimes := primeTester(100)
	
	fmt.Println(sliceOfPrimes)

	

	//this contains all possible pairs of all prime numbers up to 100 where there is no pair (x,y) where x = y
	primePairs := generateAllUniquePrimePairs(sliceOfPrimes)

	fmt.Println(primePairs)

	//return
	
	//this contains all possible letter pairings that can be encrypted into a byte for lowercase a-z
	letterPairs := generateAllLetterPairings()


	fmt.Println(letterPairs)
	
	
	for num, primeSlice := range primePairs {

		for i := 0; i < len(primeSlice); i++ {
			
			//modulus := num * primeSlice[i] 

			if num == 43 && primeSlice[i] == 59 {

			fmt.Println("made it correctly")

			

			for char, letterSlice := range letterPairs {
			
				for j := 0; j < len(letterSlice); j++ {
					
					byteString := char + letterSlice[j]
					
					gcdSlice := getGCDRelativelyPrimes(num, primeSlice[i])
					

					for s:= 0; s < len(gcdSlice); s++ {
					testCipher := encryptMessageRSA(byteString, num, primeSlice[i], gcdSlice[s])
					
					chunkFromCipher := string(cipherText[0]) + string(cipherText[1]) + string(cipherText[2]) + string(cipherText[3])
					
					//
					
					
					if testCipher == chunkFromCipher {
						fmt.Println("CIPHERS" + testCipher + chunkFromCipher)
						plaintext := make([]string, 1, 1)
						//keys are the order in which the byte appears in the cipher text starting with 0 for first one 				
						combosDecryption := make(map[int][]string)
						
						

						matchLetters := getMapOpposite()

						//fmt.Println(testCipher + " "  + strconv.Itoa(num) + " " + strconv.Itoa(primeSlice[i])+ " " + strconv.Itoa(gcdSlice[s]))

						if num < primeSlice[i] {
							//fmt.Println("OKCheck")
							
							factoredMap := factorize(gcdSlice[s], (num - 1) *(primeSlice[i] - 1))
							//factoredMap := factorize(13, 2436)
							fmt.Println(factoredMap)
							if factoredMap == nil {
								continue
							}
							//fmt.Println("OK2")
							inverseValue := factoredMap[gcdSlice[s]]
							//inverseValue := factoredMap[13]
							fmt.Println("INVERSE")
							fmt.Println(inverseValue)
							//fmt.Println("OK3")
							for c := 0; c < len(cipherText); c = c + 5 {
								fmt.Println("C" + strconv.Itoa(c))
								check1, _ := strconv.Atoi(string(cipherText[c]))
								check2, _ := strconv.Atoi(string(cipherText[c + 1]))
								check3, _ := strconv.Atoi(string(cipherText[c + 2]))
								check4, _ := strconv.Atoi(string(cipherText[c + 3]))

								//fmt.Println("OK3")

								byteFromCipherString := strconv.Itoa(check1) + strconv.Itoa(check2) + strconv.Itoa(check3) + strconv.Itoa(check4)

								fmt.Printf("%s", byteFromCipherString)

								if strconv.Itoa(check1) == "0" {
										
									byteFromCipherString = strconv.Itoa(check2) + strconv.Itoa(check3) + strconv.Itoa(check4)
							
								}
								fmt.Println(byteFromCipherString)
								//fmt.Println(byteFromCipherString)
								byteFromCipherInt, _ := strconv.Atoi(byteFromCipherString)
								
								fmt.Printf("%s11", byteFromCipherInt)
				
								byteFromCipherInt = modExpon(byteFromCipherInt, inverseValue,
num * primeSlice[i])

								
								
								hasBeenFixed := false
			

								/*
								if byteFromCipherInt == 0 {
									byteFromCipherInt = 00
									hasBeenFixed = true
								}
								*/
								
								
								fmt.Printf("%sCC", byteFromCipherInt)

								//fmt.Printf("%s", byteFromCipherInt)

								//byteFromCipherInt = byteFromCipherInt % (num * primeSlice[i])

								//byteFromCipherIntS := strconv.Itoa(byteFromCipherInt)
								//fmt.Println(byteFromCipherIntS)
								//byteFromCipherString := strconv.Itoa(byteFromCipherInt)
								var char1 string
								var char2 string

								byteFromCipherString = strconv.Itoa(byteFromCipherInt)

								combosDecryption[c/5] = allpossibleDecryptions(byteFromCipherString)
								
								
		
								fmt.Printf("%sCCCC", byteFromCipherString)

								if len(byteFromCipherString) == 2 && !hasBeenFixed {
									fmt.Println("check")
									//byteFromCipherString = "0" + byteFromCipherString
									fmt.Println(byteFromCipherString)
									hasBeenFixed = true
									
								}
								if len(byteFromCipherString) == 1 && !hasBeenFixed {
									fmt.Println("ERR")
									fmt.Println(byteFromCipherString)
									byteFromCipherString = "260" + byteFromCipherString
									os.Exit(1)
									
								}
								/*
								for len(byteFromCipherString) < 3 {

									byteFromCipherString = "0" + byteFromCipherString				

								}
								*/
								
								if len(byteFromCipherString) == 2 {
								
									char1 = string(byteFromCipherString[0])
									char2 = string(byteFromCipherString[1])

								}
		
								if len(byteFromCipherString) == 3 {
									
									char1 = string(byteFromCipherString[0])

									fmt.Printf("%sRR", char1)

									char2 = string(byteFromCipherString[1]) + string(byteFromCipherString[2])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									} 
								}else if len(byteFromCipherString) == 4{
									char1 = string(byteFromCipherString[0]) + string(byteFromCipherString[1])
									if string(char1[0]) == "0" {
										char1 = string(char1[1])
									} 
									char2 = string(byteFromCipherString[2]) + string(byteFromCipherString[3])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									} 
								}
								fmt.Println(char1)
								fmt.Println("BREAK")
								fmt.Println(char2)
								fmt.Println(matchLetters[char1])
								fmt.Println(matchLetters[char2])
								
								plaintext = append(plaintext, matchLetters[char1]) 
								plaintext = append(plaintext, matchLetters[char2])
								//fmt.Println(matchLetters[char1] + matchLetters[char2])

							}

								
							
								decodeFromAllCombos(combosDecryption , 0, "", messageChan)
							

								go recursionMessage(messageChan)
									
								

								fmt.Println(plaintext[1:])
								//decoded := &plaintext

								//shiftDecodedSliceEveryPosition(decoded)

								

							
						}else{
							fmt.Println("NEVER")
							factoredMap := factorize((primeSlice[i] - 1), (num-1))
							inverseValue := factoredMap[primeSlice[i] - 1]
							for c := 0; c < len(cipherText); c = c + 5 {
								
								if inverseValue == 937 {
									fmt.Println(cipherText)
								}
								check1, _ := strconv.Atoi(string(cipherText[c]))
								check2, _ := strconv.Atoi(string(cipherText[c + 1]))
								check3, _ := strconv.Atoi(string(cipherText[c + 2]))
								check4, _ := strconv.Atoi(string(cipherText[c + 3]))

								byteFromCipherString := strconv.Itoa(check1) + strconv.Itoa(check2) + strconv.Itoa(check3) + strconv.Itoa(check4)
								fmt.Println(byteFromCipherString)
								byteFromCipherInt, _ := strconv.Atoi(byteFromCipherString)
								
								byteFromCipherInt = myExponFunc(byteFromCipherInt, inverseValue)


								byteFromCipherInt = byteFromCipherInt % (num * primeSlice[i])

								//byteFromCipherString := strconv.Itoa(byteFromCipherInt)
								char1 := ""
								char2 := ""

								
								/*
								minusSign := string(byteFromCipherString[0])
								if minusSign == "-" {
									if len(byteFromCipherString) == 4 {
									
									char1 = string(byteFromCipherString[1])

									char2 = string(byteFromCipherString[2]) + string(byteFromCipherString[3])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									}
									}else{
										char1 = string(byteFromCipherString[1]) + string(byteFromCipherString[2])
									if string(char1[0]) == "0" {
										char1 = string(char1[1])
									} 
									char2 = string(byteFromCipherString[3]) + string(byteFromCipherString[4])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									} 
								}else 
								
								*/
								if len(byteFromCipherString) == 3 {
									char1 = string(byteFromCipherString[1])
									fmt.Println(char1)

									char2 = string(byteFromCipherString[2]) + string(byteFromCipherString[3])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									} 
								}else{
									char1 = string(byteFromCipherString[1])
									if string(char1[0]) == "0" {
										char1 = string(char1[1])
									} 
									char2 = string(byteFromCipherString[2]) + string(byteFromCipherString[3])
									if string(char2[0]) == "0" {
										char2 = string(char2[1])
									} 
								}
								plaintext = append(plaintext, matchLetters[char1] + matchLetters[char2]) 
								//fmt.Println(matchLetters[char1] + matchLetters[char2])
							
						}

								fmt.Println(plaintext[1:])
			
					}

					}

				}

			}
			

		}


		
	
				

	}
	
	
	
	
	}

	
	}

}

func recursionMessage(messageChan chan string){ 
	
	message := <- messageChan 
	if message == "done" {
		return 
	}else{
		fmt.Println(message)
		go recursionMessage(messageChan)
	}
}



func generateAllLetterPairings() map[string][]string {
	
	letterPairs := make(map[string][]string)

	alphabet := "abcdefghijklmnopqrstuvwxyz "

	for i := 0; i < len(alphabet); i++ {
	
		char := string(alphabet[i])

		sliceToAdd := make([]string, 1, 1)
	
		for j := 0; j < len(alphabet); j++ {
			charAdd := string(alphabet[j])
			sliceToAdd = append(sliceToAdd, charAdd)

		}

		letterPairs[char] = sliceToAdd[1:]

	}

	

	


	return letterPairs
	


}



func calculateGivenCipherText(cipherText string) {


	//get the number of bytes we are dealing with
	//ignore spaces and count bytes 

	charCount := 0

	byteCount := 0

	for i := 0; i < len(cipherText); i++ {
	
		char := string(cipherText[i])

		if char != " " {
			charCount++ 
		}
		
		if charCount == 4 {
		
			byteCount++ 
			charCount = 0

		}

	}

}

func getMapOpposite() map[string]string{

	alphabet := make(map[string]string)

	alphabet["0"] = "a"
	
	alphabet["1"] = "b"

	alphabet["2"] = "c"
	
	alphabet["3"] = "d"

	alphabet["4"] = "e"

	alphabet["5"] = "f"
	
	alphabet["6"] = "g"

	alphabet["7"] = "h"
	
	alphabet["8"] = "i"

	alphabet["9"] = "j"

	alphabet["10"] = "k"

	alphabet["11"] = "l"

	alphabet["12"] = "m"

	alphabet["13"] = "n"

	alphabet["14"] = "o"
	
	alphabet["15"] = "p"

	alphabet["16"] = "q"

	alphabet["17"] = "r"
	
	alphabet["18"] = "s"

	alphabet["19"] = "t"
	
	alphabet["20"] = "u"

	alphabet["21"] = "v"

	alphabet["22"] = "w"

	alphabet["23"] = "x"

	alphabet["24"] = "y"

	alphabet["25"] = "z"
	
	alphabet["26"] = " "

	return alphabet


}


func getMap() map[string]int{

	alphabet := make(map[string]int)

	alphabet["a"] = 0
	
	alphabet["b"] = 1

	alphabet["c"] = 2
	
	alphabet["d"] = 3

	alphabet["e"] = 4

	alphabet["f"] = 5
	
	alphabet["g"] = 6

	alphabet["h"] = 7
	
	alphabet["i"] = 8

	alphabet["j"] = 9

	alphabet["k"] = 10

	alphabet["l"] = 11

	alphabet["m"] = 12

	alphabet["n"] = 13

	alphabet["o"] = 14
	
	alphabet["p"] = 15

	alphabet["q"] = 16

	alphabet["r"] = 17
	
	alphabet["s"] = 18

	alphabet["t"] = 19
	
	alphabet["u"] = 20

	alphabet["v"] = 21

	alphabet["w"] = 22

	alphabet["x"] = 23

	alphabet["y"] = 24

	alphabet["z"] = 25
	
	alphabet[" "] = 26

	return alphabet
}



func encryptMessageRSA(message string, p int, q int, e int) string{

	letterMap := getMap() 
	
	evenOdd := len(message) % 2
	

	if evenOdd == 0 {

	}else{
	
		message = message + " "
		
	}

	cipherText := ""
	
	
	
	
	for i := 0; i < (len(message)/2); i++ {
		
		char1 := string(message[i*2])
		char2 := string(message[(i*2) + 1])

		//fmt.Println(char1)
		//fmt.Println(char2)
		
		char1Int := letterMap[char1]
		
		char2Int := letterMap[char2]

		
		
		byteString := strconv.Itoa(char1Int) + strconv.Itoa(char2Int)
		
		

		byteStringInt, _ := strconv.Atoi(byteString)

		

		cipherByte := modExpon(byteStringInt, e, (p * q))
		
		cipherByteString := strconv.Itoa(cipherByte)
		
		if len(cipherByteString) == 3 {
		
			cipherByteString = "0" + cipherByteString 

		}
		
		if len(cipherText) == 0 {
			cipherText = cipherByteString
		}else{
			cipherText = cipherText + " " + cipherByteString
		}

		

	}
		
	
	return cipherText




}


func myExponFunc(num int, exp int) int {
	
	sum := num
		
	if (exp == 0) {
	
		return 1

	}

	if exp < 0 {
		
		/*
		for i := 0; i > exp + 1; i = i - 1 {
			sum = sum * num
		}

		*/
		return 0

	}

	for i := 0; i < exp - 1; i++ {

		sum = sum * num 

	}
	
	return sum


}


func modExpon(base int, exponent int, mod int) int {

	
	if exponent < 0 {
		return 0 % mod
	}	

	binaryExpon := fmt.Sprintf("%b", exponent)

	result := 1

	base = base % mod
	
	for i := 0; i < len(binaryExpon); i++	{
		
		bit, err := strconv.Atoi(binaryExpon[((len(binaryExpon) - 1 - i)):(len(binaryExpon) - i)])
		
		if err != nil {
			fmt.Println("ERROR PARSING BINARY STRING PROGRAM TERMINATING")
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		
		
		if bit == 1 {
		
			result = (result * base) % mod

			base = myExponFunc(base, 2) % mod
		
		}else if bit == 0{
			
			base = myExponFunc(base, 2) % mod

		}


	}



	return result
}






func primeTester(maxPrime int) []int{

	primeSlice := make([]int, 1, 1)

	for i := 2; i < maxPrime + 1; i++ {
		
		currentNum := float64(i) 
		
		minimumToCheck := math.Sqrt(currentNum)

		minimumToCheck = math.Round(minimumToCheck)

		castToInt := int(minimumToCheck)


		isPrime := checkIfPrime(castToInt, i)

		if isPrime {
			
			primeSlice = append(primeSlice[0:], i)
	
		}
		
			
	
	}

	return primeSlice[1:]
	


}



func checkIfPrime(minimumValue int, checkAgainst int) bool {

	for i := 2; i < minimumValue + 1; i++ {

		
		remainder := checkAgainst % i

		if remainder == 0 {
		
			return false

		}
		
	}

	return true
	
}


func generateAllUniquePrimePairs(primeCandidates []int) map[int][]int {

	
	allPairs := make(map[int][]int)

	var sliceWithoutCheckVal []int

	
	for i := 0; i < len(primeCandidates); i++ {
	
		

		//get the prime number to pair with others 
		checkPrime := primeCandidates[i]
		
		sliceWithoutCheckVal = nil

		
		sliceWithoutCheckVal = primeCandidates[i+1:]
		

		
		sliceAppendTo := make([]int, 1, 1)			
			

		for j := 0; j < len(sliceWithoutCheckVal); j++ {
			
			sliceAppendTo = append(sliceAppendTo[0:], sliceWithoutCheckVal[j])		
	
		}	

		allPairs[checkPrime] = 	sliceAppendTo[1:]

	}


	return allPairs
	

}


func getGCDRelativelyPrimes(p int, q int) []int {

	pMinus1 := p - 1 

	qMinus1 := q - 1


	product := pMinus1 * qMinus1

	//fmt.Println(product)

	contendersForE := make([]int, product, product)


	//start at 2 because 1 is obviously already a contender

	for i := 2; i < product; i++ {
		addBool, num := calcGCD(i, product)
		if addBool {
			
			contendersForE = append(contendersForE, num)	
		
		}

	}

	return contendersForE[product:]

}


func calcGCD(e int, n int) (bool, int) {

	remainder := 1

	divisor := e

	main := n 

	initialE := e

	
	
	// main = 2436 divisor = 3 remainder = 1 quotient = 0 

	for remainder != 0 {
	
		remainder = main % divisor

		if remainder == 0 && divisor == 1 {
		
			return true, initialE

		}

		main = divisor

		divisor = remainder 
			
	}

	return false, -1 

}



func factorize(numSmall int, numLarge int) map[int]int{

	remainderMap := make(map[int][]int)

	valuesMapL := make(map[int]int)

	valuesMapR := make(map[int]int)

	valuesMapL[numLarge] = 0

 	valuesMapL[numSmall] = 0

	valuesMapR[numLarge] = 0

 	valuesMapR[numSmall] = 0

	numToDivide := numLarge

	numDivisor := numSmall 

	timesItFits := 0 

	remainder := 1

	row := 1	

	for remainder != 0 {

		timesItFits = numToDivide / numDivisor 

		valuesMapL[timesItFits] = 0

		valuesMapL[numToDivide % numDivisor] = 0

		valuesMapR[timesItFits] = 0

		valuesMapR[numToDivide % numDivisor] = 0

		remainderMap[row] = []int{numToDivide, timesItFits, numDivisor, (numToDivide % numDivisor)}

		remainder = numToDivide % numDivisor

		numToDivide = numDivisor 

		numDivisor = remainder

		row++ 

	}

	row = row - 1
	
	totalRows := row 

	//fmt.Println(valuesMapL)
	//fmt.Println(valuesMapR)
	//fmt.Println(remainderMap)
	
	

	leftSideDone := false

	

	//second to last row here 
	row = row - 1
	
	//fmt.Println(row)
	
	//set initial left side 
	valuesMapL[remainderMap[row][0]] = 1

	valuesMapL[remainderMap[row][2]] = -1

	//fmt.Println("initialVals")
	//fmt.Println(valuesMapL)

	bool1L := false

	//focusKey := false

	key := 0

	factorToTrack := 0 

	valuesMapLCopy := make(map[int]int)

	affected := make([]int, 0, 0)

	
	count := 0 


	//1 is the special case because if 

	onceOnly := true


	for !leftSideDone {
		row = row - 2

		//fmt.Println(row)
		//fmt.Printf("%s", len(remainderMap))
		//fmt.Println(remainderMap)

		if len(remainderMap) <= 1 {
			fmt.Println("major error")			
			return nil
		}

		if len(remainderMap) <= 3 {
			
			fmt.Println("specialCase")

			specialCase := make(map[int]int)

			smallNumFrequency := remainderMap[1][1] + remainderMap[2][2]
			fmt.Println(smallNumFrequency)
			specialCase[numSmall] = smallNumFrequency
	
			specialCase[numLarge] = 0
			
			fmt.Println("specialCase")

			return specialCase

		}
		

		if onceOnly {
			valuesMapL[remainderMap[row][0]] = 1 

			valuesMapL[remainderMap[row][2]] = -(remainderMap[row][1])	

			valuesMapL[remainderMap[row + 2][0]] = 0

			//fmt.Println("initialVals")
			//fmt.Println(valuesMapL)

			//fmt.Printf("%s", row)

			onceOnly = false
		}




		for keyMap, val := range valuesMapL {
	
			//fmt.Println("TEST")
			//fmt.Println(valuesMapL)

			
			if (keyMap != numSmall && keyMap != numLarge) && (val > 0 || val < 0)  {
				for key4, val4 := range remainderMap {

					if val4[3] == keyMap {
						row = key4
					}
					

				}

				
				count ++ 
				//fmt.Println(valuesMapL)
				bool1L = true
				//focusKey = true	
				//fmt.Println(key)
				key = keyMap 
				factorToTrack = valuesMapL[key]	
				//fmt.Println(factorToTrack)
				//row = row + 1	
				valuesMapLCopy = nil
				valuesMapLCopy = make(map[int]int)
				for keyMap2, val2 := range valuesMapL {
			
					valuesMapLCopy[keyMap2] = val2
	
				}
				valuesMapL[keyMap] = 0
				
				valuesMapLCopy, affected = factorUpValOfFocus(numSmall, numLarge, factorToTrack, valuesMapLCopy, remainderMap,
				row, keyMap)
				
				//fmt.Println("BEFORE2")

			//	fmt.Println(valuesMapL)
			//	fmt.Println(len(affected))


				for i := 0; i < len(affected); i++ {
				
					valuesMapL[affected[i]] = valuesMapL[affected[i]] + valuesMapLCopy[affected[i]]

				}
				
				if checkReturn(valuesMapL, numSmall, numLarge) {

					return valuesMapL			

				}
				
		
				/*
				for key3, _ := range valuesMapLCopy {

					
					
					
						
						valuesMapL[key3] = valuesMapL[key3] + valuesMapLCopy[key3]

					

				}
			
				*/
				
				//fmt.Println("AFTER2")

				//fmt.Println(valuesMapL)
				//fmt.Println(valuesMapL)
				if count == 1000 {
					return valuesMapL
				}
				break
			}
		


		}

		
		if bool1L {
			bool1L = false
			continue 

		}else{
			leftSideDone = true
		}
		

	}

	row = totalRows

	return valuesMapL

	
	
}




func factorUpValOfFocus(numSmall int, numLarge int, factorToTrack int, copyMap map[int]int, remainderMap map[int][]int,
			row int, valOfFocus int) (map[int]int, []int) {

	restart := true

	for key, val := range remainderMap {

		if val[3] == valOfFocus {
			row = key
		}

	}

	

	for key0, _ := range copyMap {

		if key0 != valOfFocus {
			copyMap[key0] = 0
		}

	}
	
	//fmt.Printf("%s", factorToTrack)
	

	//fmt.Println("BEFORE")
	//fmt.Println(copyMap)


	keysAffected := make([]int, 1, 1)
	
	
	for restart {
		
		//fmt.Println("RESTART")
		//fmt.Println(remainderMap[row][0])
		//fmt.Println(remainderMap[row][2])
		//fmt.Println(remainderMap[row + 2][0])
		keysAffected = append(keysAffected, remainderMap[row][0])

		copyMap[remainderMap[row][0]] = copyMap[remainderMap[row][0]] + 1 

		keysAffected = append(keysAffected, remainderMap[row][2])

		copyMap[remainderMap[row][2]] = copyMap[remainderMap[row][2]] -(remainderMap[row][1])	
		
		keysAffected = append(keysAffected, remainderMap[row + 2][0])

		copyMap[remainderMap[row + 2][0]] = 0	

		

		onlyTwoOfInterest := false

		for key, val := range copyMap {
			

			if key != numSmall && key != numLarge && (val > 0 || val < 0) && key != 2 && false {
										
					onlyTwoOfInterest = false
										
					break
				}
				onlyTwoOfInterest = true			
			

			if onlyTwoOfInterest {
				
				//fmt.Println("After")

				//fmt.Println(copyMap)
				
				for k, _ := range copyMap {
			
					copyMap[k] = copyMap[k] * factorToTrack
				
				}
				
				return copyMap, keysAffected[1:]

			}

		}

	}
	return nil, nil

}



func checkReturn(checkMap map[int]int, numSmall int, numLarge int ) bool {
	
	canReturn := true

	for key, val := range checkMap {
	
	
		if key != numSmall && key != numLarge && val != 0 {

			canReturn = false

		}


	}

	return canReturn



}


func shiftDecodedSliceEveryPosition(slice *[]string) {
	

	copySlice := *slice

	checkForPlainText(copySlice)	

	copySlice = copySlice[1:]

	letterToNum := getMap() 
	
	numToLetter := getMapOpposite()


	for j := 0; j < 27; j++ {

	for i := 0; i < len(copySlice); i++ {


		letter := copySlice[i]

		valueOfLetterInt := letterToNum[letter]
		
		//valueOfLetterInt, _ := strconv.Atoi(valueOfLetter)

		valueOfLetterInt++

		if valueOfLetterInt == 27 {
			valueOfLetterInt = 0 
		}

		stringLetterNum := strconv.Itoa(valueOfLetterInt)
		copySlice[i] = numToLetter[stringLetterNum]



	}	
		checkForPlainText(copySlice)
		fmt.Print(copySlice)
		
	
	}
	
}


func checkForPlainText(slice []string) {


	copySlice := slice

	wholeSLice := ""

	for i := 0; i < len(copySlice); i++ {
	
		
		wholeSLice = wholeSLice + copySlice[i]



	}

	

	if strings.Contains(wholeSLice, "s") && strings.Contains(wholeSLice, "t") && strings.Contains(wholeSLice, "o") &&
		strings.Contains(wholeSLice, "p") && strings.Contains(wholeSLice, "t") && strings.Contains(wholeSLice, "h") &&
		strings.Contains(wholeSLice, "i") && strings.Contains(wholeSLice, "s") && strings.Contains(wholeSLice, "m") &&
		strings.Contains(wholeSLice, "a") && strings.Contains(wholeSLice, "d") && strings.Contains(wholeSLice, "n") &&
		strings.Contains(wholeSLice, "e") && strings.Contains(wholeSLice, "s") && strings.Contains(wholeSLice, "s") {

			fmt.Println(wholeSLice)
			fmt.Println("Message Decrypted")
			os.Exit(1)	


	}


}



func allpossibleDecryptions(decryptedByte string) []string {


	//allCombos := make([]string, 1, 1)
	
	
	

	if len(decryptedByte) == 1 {
		
		combo1 := "0" + decryptedByte
		combo2 := decryptedByte + "0" 
		combo3 := decryptedByte + "0" + "0" 
		combo4 := "0" + decryptedByte + "0"
		combo5 := "0" + "0" + decryptedByte 	
		combo6 := decryptedByte + "0" + "0" + "0"
 		combo7 := "0" + decryptedByte + "0" + "0"
		combo8 := "0"  +  "0" + decryptedByte + "0"
		combo9 := "0" + "0" + "0" + decryptedByte

		return []string{combo1, combo2, combo3, combo4, combo5, combo6, combo7, combo8, combo9}
		

	}else if len(decryptedByte) == 2{
		
		char1 := string(decryptedByte[0])
		char2 := string(decryptedByte[1])

		combo1 := char1 + char2 	
		combo2 := "0" + char1 + char2 
		combo3 := char1 + "0" + char2
		combo4 := char1 + char2 + "0"
		combo5 := "0" + "0" + char1 + char2 
		combo6 := "0" + char1 + "0" + char2
		combo7 := "0" + char1 + char2 + "0"
		combo8 := char1 + "0" + char2 + "0"
		combo9 := char1 + "0" + "0" + char2
		combo10 := char1 + char2 + "0" + "0"	

		return []string{combo1, combo2, combo3, combo4, combo5, combo6, combo7, combo8, combo9, combo10}	


	}else if len(decryptedByte) == 3 {
		
		char1 := string(decryptedByte[0])
		char2 := string(decryptedByte[1])
		char3 := string(decryptedByte[2])
		
		combo1 := char1 + char2 + char3 
		combo2 := "0" + char1 + char2 + char3 
		combo3 := char1 + "0" + char2 + char3 
		combo4 := char1 + char2 + "0" + char3
		combo5 := char1 + char2 + char3 + "0"

		return []string{combo1, combo2, combo3, combo4, combo5}		
		


	}else if len(decryptedByte) == 4 {
		return []string{decryptedByte}
	}

	return nil

}

func decodeFromAllCombos(allCombos map[int][]string, timesAdded int, finalString string, messageChan chan string) {
	
	fmt.Printf("%slength", len(allCombos))

	if timesAdded == len(allCombos) {
	
		messageChan <- decrypt(finalString)
		messageChan <- "done"
		return

	}
	
	timesAddedPlus := timesAdded + 1

	sliceToExpand := allCombos[timesAdded] 

	for i := 0; i < len(sliceToExpand); i++{

		go decodeFromAllCombos(allCombos, timesAddedPlus, finalString + sliceToExpand[i], messageChan)

	}
	



}


func decrypt(byteFromCipherString string) string{

	plaintext := ""

	matchLetters := getMapOpposite()

	for c := 0; c < len(byteFromCipherString); c++ {

		char1 := string(byteFromCipherString[0]) + string(byteFromCipherString[1])
		if string(char1[0]) == "0" {
			char1 = string(char1[1])
		} 
		char2 := string(byteFromCipherString[2]) + string(byteFromCipherString[3])
		if string(char2[0]) == "0" {
			char2 = string(char2[1])
		} 

		
		char1 = matchLetters[char1] 	
		char2 = matchLetters[char2]

		plaintext = plaintext + char1 + char2
	
	

	}

	return plaintext

	


}








