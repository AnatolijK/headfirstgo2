//пакет datafile предназначен для чтения из файлов
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки файла
func GetFloats(fileName string) ([]float64, error) { // функция возвращает ctuvtyn значенией float64
	var numbers []float64 // переменная поумолчанию содержит nil
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err //возвращаем nil вместо сегмента. (Переменная в этот момент в любом случае
		// равна nil)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err //возвращаем nil вместо сегмента
		}
		numbers = append(numbers, number) // новое число присоединяется к сегменту

	}
	err = file.Close()
	if err != nil {
		return nil, err //возвращаем nil вместо сегмента
	}
	if scanner.Err() != nil {
		return nil, scanner.Err() //возвращаем nil вместо сегмента
	}
	return numbers, nil
}
