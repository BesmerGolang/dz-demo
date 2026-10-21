package storage

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var (
	mu          sync.Mutex
	storageFile = "tokens.json"
)

// функция генерации токена
func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
func SaveToken(token, email string) error {
	mu.Lock()
	defer mu.Unlock()
	data := make(map[string]string)
	/*ниже условие, если нет ошибки и файл не пустой - превращаем байты
	полученные от os.ReadFile в пригодные для чтения и записываем их в мапу*/
	fileBytes, err := os.ReadFile(storageFile)
	// Если ошибка есть И это НЕ ошибка "файл не существует", прерываем работу
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	// Если данные есть, парсим их
	if len(fileBytes) > 0 {
		if err := json.Unmarshal(fileBytes, &data); err != nil {
			return err
		}
	}
	data[token] = email //записали новые данные

	//маршал интент позволяет форматировать данные перед чтением
	updatedBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	//берем отформатированные байты и имя файла и делаем в него запись 0644 права доступа к файлу из систем Unix/Linux
	return os.WriteFile(storageFile, updatedBytes, 0644)
}
func CheckToken(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	data := make(map[string]string)
	fileBytes, err := os.ReadFile(storageFile)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return false
	}
	if len(fileBytes) > 0 {
		if err := json.Unmarshal(fileBytes, &data); err != nil {
			return false
		}
	}
	_, exist := data[token]
	if !exist {
		return false
	}
	delete(data, token)
	updateBytes, err := json.Marshal(data)
	if err != nil {
		return false
	}
	err = os.WriteFile(storageFile, updateBytes, 0644)
	return true
}
