package main

import (
	"fmt"
	"time"
)

// Задание 1: Вывод текущего времени и даты
func task1() {
	fmt.Println("=== Задание 1: Текущее время и дата ===")
	currentTime := time.Now()
	fmt.Printf("Текущее время: %s\n", currentTime.Format("15:04:05"))
	fmt.Printf("Текущая дата: %s\n", currentTime.Format("02.01.2006"))
	fmt.Printf("Полная дата и время: %s\n\n", currentTime.Format("02.01.2006 15:04:05"))
}

// Задание 2: Переменные различных типов
func task2() {
	fmt.Println("=== Задание 2: Переменные различных типов ===")
	
	// Объявление переменных с явным указанием типа
	var integerVar int = 42
	var floatVar float64 = 3.14159
	var stringVar string = "Hello, Go!"
	var boolVar bool = true
	
	fmt.Printf("Целое число: %d (тип: %T)\n", integerVar, integerVar)
	fmt.Printf("Число с плавающей запятой: %.5f (тип: %T)\n", floatVar, floatVar)
	fmt.Printf("Строка: %s (тип: %T)\n", stringVar, stringVar)
	fmt.Printf("Логическое значение: %t (тип: %T)\n\n", boolVar, boolVar)
}

// Задание 3: Краткая форма объявления переменных
func task3() {
	fmt.Println("=== Задание 3: Краткая форма объявления переменных ===")
	
	// Краткая форма объявления (тип определяется автоматически)
	name := "Алексей"
	age := 25
	temperature := 23.7
	isStudent := true
	
	fmt.Printf("Имя: %s (тип: %T)\n", name, name)
	fmt.Printf("Возраст: %d (тип: %T)\n", age, age)
	fmt.Printf("Температура: %.1f°C (тип: %T)\n", temperature, temperature)
	fmt.Printf("Студент: %t (тип: %T)\n\n", isStudent, isStudent)
}

// Задание 4: Арифметические операции с целыми числами
func task4() {
	fmt.Println("=== Задание 4: Арифметические операции с целыми числами ===")
	
	var a, b int
	a = 15
	b = 4
	
	fmt.Printf("Число a = %d, число b = %d\n", a, b)
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Деление: %d / %d = %d (целочисленное деление)\n", a, b, a/b)
	fmt.Printf("Остаток от деления: %d %% %d = %d\n", a, b, a%b)
	fmt.Printf("Возведение в степень: %d^%d = %.0f\n\n", a, b, pow(float64(a), float64(b)))
}

// Задание 5: Функция для вычисления суммы и разности чисел с плавающей запятой
func calculateSumAndDiff(x, y float64) (float64, float64) {
	sum := x + y
	difference := x - y
	return sum, difference
}

func task5() {
	fmt.Println("=== Задание 5: Сумма и разность чисел с плавающей запятой ===")
	
	num1 := 7.5
	num2 := 3.2
	
	sum, diff := calculateSumAndDiff(num1, num2)
	
	fmt.Printf("Число 1: %.2f\n", num1)
	fmt.Printf("Число 2: %.2f\n", num2)
	fmt.Printf("Сумма: %.2f + %.2f = %.2f\n", num1, num2, sum)
	fmt.Printf("Разность: %.2f - %.2f = %.2f\n\n", num1, num2, diff)
}

// Задание 6: Вычисление среднего значения трех чисел
func task6() {
	fmt.Println("=== Задание 6: Среднее значение трех чисел ===")
	
	// Вариант 1: с целыми числами
	num1 := 10
	num2 := 20
	num3 := 30
	averageInt := float64(num1+num2+num3) / 3.0
	
	fmt.Printf("Целые числа: %d, %d, %d\n", num1, num2, num3)
	fmt.Printf("Среднее значение: (%.1f + %.1f + %.1f) / 3 = %.2f\n", 
		float64(num1), float64(num2), float64(num3), averageInt)
	
	// Вариант 2: с числами с плавающей запятой
	f1 := 8.5
	f2 := 12.7
	f3 := 9.3
	averageFloat := (f1 + f2 + f3) / 3.0
	
	fmt.Printf("\nЧисла с плавающей запятой: %.1f, %.1f, %.1f\n", f1, f2, f3)
	fmt.Printf("Среднее значение: (%.1f + %.1f + %.1f) / 3 = %.2f\n\n", f1, f2, f3, averageFloat)
}

// Вспомогательная функция для возведения в степень
func pow(x, y float64) float64 {
	result := 1.0
	for i := 0; i < int(y); i++ {
		result *= x
	}
	return result
}

func main() {
	fmt.Println("Лабораторная работа 1 на Go")
	fmt.Println("===========================\n")
	
	task1() // Задание 1
	task2() // Задание 2
	task3() // Задание 3
	task4() // Задание 4
	task5() // Задание 5
	task6() // Задание 6
	
	fmt.Println("Все задания выполнены!")
}