Загрузчик
Вы создаете программу для загрузки файлов.
Чтобы ускорить загрузку, вы решили использовать многопоточность. Каждая загрузка файла будет выполняться как отдельная горутина.

Для имитации загрузки файла функция download() должна принимать в качестве аргумента целое число (которое будет играть роль размера файла) и подсчитывать сумму всех целых чисел от 0 до этого числа (включительно).

Данная программа принимает три размера файлов в качестве входных данных и вызывает функцию download() как горутину для каждого файла.
Завершите программу, объявив функции download() и отправив результат их работы в main() с помощью каналов. Результаты работы трех функций должны быть сложены и выведены на экран.

Вы можете использовать каналы, чтобы получить данные из функций download() и сложить их вместе.

Sample Input:

10
100
42
Sample Output:

6008