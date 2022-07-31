//Exemplo 2: Tipos de variável
//Ao escolher um tipo para nossas variáveis, devemos nos atentar
//em suas características e também à quantidade de memória necessária para armazená-la.
package main

import (
    "fmt"
    "strings"
)

func main() {
       
        //Declaração de variáveis e inicializando
        
        //Inteiros
        var a int = 3    //Variável inteira   - Tamanho: 4 bytes/ 32 bits
        //Varias formas de inteiro: int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
        var b byte       //Variável byte      - Tamanho: 1 byte / 8  bits -Um alias para uint8
        var c rune       //É um alias int32 e é usada para armazenar Unicode
        //Ponto Flutuante
        var d float32 = 2.14   //Variável flutuante - Tamanho: 4 bytes/ 32 bits ou 64 bits
        
        var nome string = "Gabriel"
        
        var flag boolean; //Variável booleana (Variável lógica (verdadeiro/falso)))
               
       //Inicializando Variáveis
        
       fmt.Println("Resultado da soma: ")
       fmt.Println(a+d)
        
       //A precisão que o ponto flutuante dá para a equação
       //é desnecessária no exemplo da soma
      
       //Agora veja a diferença entre os tipos em uma divisão não exata  
       //Ponto Flutuante
       fmt.println("Resultado da divisao: ")
       fmt.println(a/d)
       variavel := a/2
       fmt.println("Resultado da divisao:\n ")
       fmt.println(variavel)
       //No caso do inteiro, todas as casas após a vírgula são desconsideradas.
    }   
}
