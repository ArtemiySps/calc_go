# КАЛЬКУЛЯТОР НА ЯЗЫКЕ GOLANG

Данный калькулятор поможет в простых математических вычислениях с числами

Калькулятор полностью написан на языке Golang в рамках курса "Программирование на Go | 24" от Яндекс.Лицея

---
## Описание

В данном проекте использовался собственный код калькулятора из финального задания Спринта 0. Вполне возможно, и я более чем уверен, что существуют куда более компактные импликации данного кода)) Тем не менее, он был структурирован по файлам и загружен на Github

Для себя я выделил несколько этапов его модернизации:
1. Реализовать http-сервер
2. Проработать все возможные ошибки
3. Сделать логирование для http-сервера
4. Написать тесты как для самого калькулятора, так и для http-сервера
5. Написать качественный README   :)

## Как использовать



## Общая структура файлов

![[Pasted image 20241222122821.png]]

## Логика калькулятора

Все вычисления происходят в файле pkg/calculation/calculation.go

Основная функция, из которой вызываются все остальные:
`func Calc(expression string) (float64, error)`
В ней обрабатываются ошибки, возникшие внутри остальных функций

После отправки запроса пользователем, входное выражение проверяется внутри функции:
`func checkString(expression string) error`
Здесь идут проверки с помощью циклов на возможные ошибки ввода (подробнее см. Обработка ошибок)

Затем, происходит разбиение выражения на слайс, где элементами являются строки с числами (или операциями в скобках) с стоящими перед ними знаками. Это происходит в функции:
`func makeSlice(expression string) []string`
Примеры:

| исходное выражение | полученный слайс      |
| ------------------ | --------------------- |
| 2+2                | \[+2 +2]              |
| 2+3*(4+5)          | \[+2 +3 \*(4+5)]      |
| 2+3*(4+(5+6))      | \[+2 +3 *\(4+\(5+6))] |


