package main

func numbersLearning() {
	// A diferença entre os tipos é se aceitam número negativo e a quantidade de bits das vars
	var temp int8 = -5  //vai de -128 a 127
	var Utemp uint8 = 5 //vai de 0 a 255
	// temp = Utemp        # precisa ser do mesmo tipo
	println(temp, Utemp)

	// Número de bits
	const justAtemp int = 10 //Depende da arquitetura do processador mas geralmente 32 ou 64 bits
	//ai tem int8, int16, int32, int64 e da pra colocar o u

	const runeTemp rune = -2147483648
	const byteTemp byte = 255
	//int8 também é chamado de byte, pode ser um alias pro tipo, assim como rune é outro nome pra int32

}
