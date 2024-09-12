package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Библиотека golang.org/x/crypto/bcrypt предоставляет реализацию алгоритма шифрования паролей с использованием метода bcrypt.
// Она используется для безопасного хеширования паролей, что делает их более устойчивыми к атакам методом грубой силы или к подбору паролей.
// Метод bcrypt.CompareHashAndPassword сравнивает хешированный пароль с введённым пользователем паролем, проверяя их совпадение.
// Этот алгоритм считается одним из самых безопасных для хранения паролей в базе данных.

// GenerateHash создаёт SHA-256 хеш строки...
func GenerateHash(input string) string {
	hash := sha256.New()                   // Создаем новый SHA-256 хеш.
	hash.Write([]byte(input))              // Добавляем строку в хеш.
	hashedBytes := hash.Sum(nil)           // Получаем итоговый хеш в виде байтового среза.
	return hex.EncodeToString(hashedBytes) // Конвертируем байты в строку в формате hex.
}

// // CheckPasswordHash сравнивает хешированный пароль с введённым пользователем паролем...
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// CheckPasswordHash сравнивает хешированный пароль с введённым пользователем паролем...
func CheckPasswordHash(password, hash string) bool {
	return GenerateHash(password) == hash
}
