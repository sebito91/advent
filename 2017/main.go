package main

import (
	"fmt"

	"github.com/sebito91/advent/2017/four"
)

// 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
// 1111 produces 4 because each digit (all 1) matches the next.
// 1234 produces 0 because no digit matches the next.
// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

// main
func main() {
	fmt.Printf("Running each of the 2017 Advent of Code challenges...\n")

	//	for _, x := range []string{"1122", "1111", "1234", "91212129", str} {
	//		if err := one.DayOne(x); err != nil {
	//			fmt.Printf("failed DayOne (%s)! %+v\n", x, err)
	//		}
	//	}
	//
	//fmt.Printf("*********************************************\n")
	//
	//	for _, x := range []string{"1212", "1221", "123425", "123123", "12131415", str} {
	//		if err := one.DayOnePartTwo(x); err != nil {
	//			fmt.Printf("failed DayOnePartTwo (%s)! %+v\n", x, err)
	//		}
	//	}
	//
	//fmt.Printf("*********************************************\n")
	//
	//	if err := two.DayTwoPartOne(keys); err != nil {
	//		fmt.Printf("error in DayTwoPartOne: %+v\n", err)
	//	}
	//
	//fmt.Printf("*********************************************\n")
	//
	//	if err := two.DayTwoPartTwo(keys); err != nil {
	//		fmt.Printf("error in DayTwoPartTwo: %+v\n", err)
	//	}
	//
	//fmt.Printf("*********************************************\n")
	//
	//	for _, x := range []int{1, 12, 23, 1024, 347991} {
	//		if err := three.DayThreePartOne(x); err != nil {
	//			fmt.Printf("error in DayThreePartOne: %+v\n", err)
	//		}
	//	}
	//
	//	fmt.Printf("*********************************************\n")
	//
	//	for _, x := range []int{1, 10, 25, 806, 1024, 347991} {
	//		if err := three.DayThreePartTwo(x); err != nil {
	//			fmt.Printf("error in DayThreePartTwo: %+v\n", err)
	//		}
	//	}

	if err := four.DayFourPartOne(); err != nil {
		fmt.Printf("error in DayFourPartOne: %+v\n", err)
	}
}

// DayThree key
//var key = 347991

// DayTwo key
//var keys = [][]int{
//	[]int{5806, 6444, 1281, 38, 267, 1835, 223, 4912, 5995, 230, 4395, 2986, 6048, 4719, 216, 1201},
//	[]int{74, 127, 226, 84, 174, 280, 94, 159, 198, 305, 124, 106, 205, 99, 177, 294},
//	[]int{1332, 52, 54, 655, 56, 170, 843, 707, 1273, 1163, 89, 23, 43, 1300, 1383, 1229},
//	[]int{5653, 236, 1944, 3807, 5356, 246, 222, 1999, 4872, 206, 5265, 5397, 5220, 5538, 286, 917},
//	[]int{3512, 3132, 2826, 3664, 2814, 549, 3408, 3384, 142, 120, 160, 114, 1395, 2074, 1816, 2357},
//	[]int{100, 2000, 112, 103, 2122, 113, 92, 522, 1650, 929, 1281, 2286, 2259, 1068, 1089, 651},
//	[]int{646, 490, 297, 60, 424, 234, 48, 491, 245, 523, 229, 189, 174, 627, 441, 598},
//	[]int{2321, 555, 2413, 2378, 157, 27, 194, 2512, 117, 140, 2287, 277, 2635, 1374, 1496, 1698},
//	[]int{101, 1177, 104, 89, 542, 2033, 1724, 1197, 474, 1041, 1803, 770, 87, 1869, 1183, 553},
//	[]int{1393, 92, 105, 1395, 1000, 85, 391, 1360, 1529, 1367, 1063, 688, 642, 102, 999, 638},
//	[]int{4627, 223, 188, 5529, 2406, 4980, 2384, 2024, 4610, 279, 249, 2331, 4660, 4350, 3264, 242},
//	[]int{769, 779, 502, 75, 1105, 53, 55, 931, 1056, 1195, 65, 292, 1234, 1164, 678, 1032},
//	[]int{2554, 75, 4406, 484, 2285, 226, 5666, 245, 4972, 3739, 5185, 1543, 230, 236, 3621, 5387},
//	[]int{826, 4028, 4274, 163, 5303, 4610, 145, 5779, 157, 4994, 5053, 186, 5060, 3082, 2186, 4882},
//	[]int{588, 345, 67, 286, 743, 54, 802, 776, 29, 44, 107, 63, 303, 372, 41, 810},
//	[]int{128, 2088, 3422, 111, 3312, 740, 3024, 1946, 920, 131, 112, 477, 3386, 2392, 1108, 2741},
//}
//
//// DayOne key
//var str = "3294199471327195994824832197564859876682638188889768298894243832665654681412886862234525991553276578641265589959178414218389329361496673991614673626344552179413995562266818138372393213966143124914469397692587251112663217862879233226763533911128893354536353213847122251463857894159819828724827969576432191847787772732881266875469721189331882228146576832921314638221317393256471998598117289632684663355273845983933845721713497811766995367795857965222183668765517454263354111134841334631345111596131682726196574763165187889337599583345634413436165539744188866156771585647718555182529936669683581662398618765391487164715724849894563314426959348119286955144439452731762666568741612153254469131724137699832984728937865956711925592628456617133695259554548719328229938621332325125972547181236812263887375866231118312954369432937359357266467383318326239572877314765121844831126178173988799765218913178825966268816476559792947359956859989228917136267178571776316345292573489873792149646548747995389669692188457724414468727192819919448275922166321158141365237545222633688372891451842434458527698774342111482498999383831492577615154591278719656798277377363284379468757998373193231795767644654155432692988651312845433511879457921638934877557575241394363721667237778962455961493559848522582413748218971212486373232795878362964873855994697149692824917183375545192119453587398199912564474614219929345185468661129966379693813498542474732198176496694746111576925715493967296487258237854152382365579876894391815759815373319159213475555251488754279888245492373595471189191353244684697662848376529881512529221627313527441221459672786923145165989611223372241149929436247374818467481641931872972582295425936998535194423916544367799522276914445231582272368388831834437562752119325286474352863554693373718848649568451797751926315617575295381964426843625282819524747119726872193569785611959896776143539915299968276374712996485367853494734376257511273443736433464496287219615697341973131715166768916149828396454638596713572963686159214116763"
