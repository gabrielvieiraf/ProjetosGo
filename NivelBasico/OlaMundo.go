//Um Package (Pacote) é uma unidade composta por arquivos fonte que juntos servem a um propósito comum. 
package main //Pacote Executável
import "fmt" //fmt é um pacote que pode ser reutilizado dentro de outros pacotes, também conhecido como biblioteca.
//Função main onde darei um 'oi' através da função Println presente no pacote 'fmt'.
func main(){
  fmt.Println("Ola Mundo!") 
}
