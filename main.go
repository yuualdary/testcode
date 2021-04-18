package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


 
func bahasa(str int)string{

		Basicnumber := []string{ "Satu","Dua","Tiga","Empat","Lima","Enam","Tujuh","Delapan","Sembilan","Sepuluh",
								"Sebelas","Duabelas","Tigabelas","Empatbelas","Limabelas","Enambelas","Tujuhbelas","Delapanbelas","Sembilanbelas",
							}
		Multipleoftens := []string{"Dua Puluh","Tiga Puluh","Empat Puluh","Lima Puluh","Enam Puluh","Tujuh Puluh","Delapan Puluh","Sembilan Puluh",}

		// Multipleofhundred := []string{"Ribu","Juta"}
						
		// fmt.Println(str)
		test := str%10

		
		if str == 0 {
			return ""
		}
		if str < 20 {
			fmt.Println(Basicnumber[str-1]) //for get which array 0-18
			return Basicnumber[str-1]
		}
		if str < 100{
			// if test == 0{
			// 	fmt.Println(Multipleoftens[str/10-2])//51/10 = 5,1-2 = 3,1 =[3] lima puluh, + 51
			// 	return Multipleoftens[str/10-2]
			// }else{
				fmt.Println(Multipleoftens[str/10-2]+" " + Basicnumber[test-1] )//51/10 = 5,1-2 = 3,1 =[3] lima puluh, + 51
				return Multipleoftens[str/10-2] + " " + Basicnumber[test-1] 
			// }
		}
		if str < 120{		
			// fmt.Println("Seratus" + " " + Basicnumber[str%100-1] )//51/10 = 5,1-2 = 3,1 =[3] lima puluh, + 51
			fmt.Println("Seratus" + " " + bahasa(str%100) )//51/10 = 5,1-2 = 3,1 =[3] lima puluh, + 51
			return  bahasa(str%100)
			// return ("Seratus" + " " + Basicnumber[str%100-1] )
		}	
		if str < 1000{
			// test2 := str%10//110 %10 = 10
			// if test2 == 0{
				// fmt.Println(Basicnumber[str/100-1]+" Ratus " + Multipleoftens[str%100/10-2])//Get First = 1-19 (ex: 110 : 210/100-1 = [1] = Dua + Ratus + 
				// 																		//Get Middle = >20 : 210/100-2 = [0] = Dua Puluh)
				// return Basicnumber[str/100-1]+" Ratus " + Multipleoftens[str%100/10-2]
				fmt.Println(Basicnumber[str/100-1]+" Ratus " + bahasa(str%100))
				return Basicnumber[str/100-1]
			// }else{

			// 	fmt.Println(Basicnumber[str/100-1]+" Ratus " + Multipleoftens[str/100-2] +  Basicnumber[test2-1])//Get First = 1-19 (ex: 110 : 210/100-1 = [1] = Dua + Ratus + 
			// 		//Get Middle = >20 : 210/100-2 = [0] = Dua Puluh)	
			// 	return Basicnumber[str/100-1]+" Ratus " + Multipleoftens[str/100-2] +  Basicnumber[str/100-2]
			// }
		}	

		if str < 100000{
		
			if str%1000 == 0{
				fmt.Println(Multipleoftens[str/1000-2]+" Ribu " + bahasa(str%1000) + bahasa(str%1000/1000) )
				return Multipleoftens[str/1000-2]+" Ribu " + bahasa(str%1000) + bahasa(str%1000/1000) 
			}else{
				fmt.Println(Multipleoftens[str/10000-2]+" Ribu " + bahasa(str%1000) + bahasa(str%1000/1000) )
				return Multipleoftens[str/10000-2]+" Ribu " + bahasa(str%1000) + bahasa(str%1000/1000) 
			}

		}
		
		if str < 1000000{
		
			if str%100000 == 0{
				fmt.Println(Basicnumber[str/100000-1]+" Ratus " + bahasa(str/100000-1) + bahasa(str%1000/1000) )
				return Basicnumber[str/100000-1]+" Ratus " + bahasa(str/100000-1) + bahasa(str%1000/1000)
			}else{
				fmt.Println(Basicnumber[str/100000-1]+" Ratus " + Multipleoftens[str%10000/1000-2]+  Basicnumber[(str%10000/1000-1)] )
				return Basicnumber[str/100000-1]+bahasa(str%1000/1000)
			}

		}






	return "Error"
}

func main(){

	text := bufio.NewReader(os.Stdin)

	fmt.Println("Input Text :.....")

	value,_ := text.ReadString('\n')

	str := strings.Replace(value, "\r\n","",-1)
	
	number,_ := strconv.Atoi(str)
	// fmt.Println(number)

	bahasa(number)
	
	// str := os.Args[1] //if 0 they will show path of file/ 1 mean first index

	// for conv, _ := strconv.Atoi(str) ; conv <=10; conv++{

	// 	fmt.Print(conv)

	// } 

}