package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e coloca um hash na string
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}
//VerificarSenha compara senha e com hash e retorna se são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}

