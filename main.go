package main

import (
	"fmt"
)

func main() {

	// int int8 = -255, 255; int16 = -65535, 65535;
	// int32 = -4,294,967,295, 4,294,967,295; int64 =   -18,446,744,073,709,551,615,
	//                                                  18,446,744,073,709,551,615;
	// uint (unsigned int) tem a mesma notação (uint uint8, uint16, uint32, uint64), mas
	// trata de números inteiros módulo (sem sinal).

	// o tipo "byte" é um alias pra uint8 (unsigned), e o tipo "rune" é um int32 (signed).

	// float32 e float34 são os equivalentes para pontos flutuantes;

	// complex32 e complex64 tratam de números imaginários para usos específicos da matemática discreta.

	var numeroBananas uint8
	var numeroCidadesBrasileiras uint16
	var numeroCidadesNoMundo uint32
	var numeroEstrelasNoCeu uint64

	numeroBananas = 10
	numeroCidadesBrasileiras = 5565
	numeroCidadesNoMundo = 2500000
	numeroEstrelasNoCeu = 400000000000

	fmt.Println("possuo ", numeroBananas, " em meu bolso.\n",
		"tem", numeroCidadesBrasileiras, "cidades no brasil.\n",
		"já no mundo, existem", numeroCidadesNoMundo, "cidades.\n",
		"e estima-se que tenham sido descobertas aproximadamente", numeroEstrelasNoCeu, "estrelas na galáxia.")

	fmt.Println("\n")

	fmt.Println("os valores estão armazenados em ", &numeroBananas, &numeroCidadesBrasileiras,
		&numeroCidadesNoMundo, &numeroEstrelasNoCeu, "respectivamente.")

	fmt.Println("\n")

	fmt.Printf("o número de bananas está representado por uma variável %T,\n"+
		"o número de cidades brasileiras em %T,\n"+"o número de cidades do mundo está em %T,\n"+
		"e o número de estrelas no céu está em %T.",
		numeroBananas, numeroCidadesBrasileiras, numeroCidadesNoMundo, numeroEstrelasNoCeu)
}
