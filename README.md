# JumpingRabbit
В нашем зоопарке появился заяц. Его поместили в клетку, и чтобы ему не было скучно, директор зоопарка распорядился поставить в его клетке лесенку. Теперь наш зайчик может прыгать по лесенке вверх, перепрыгивая через ступеньки. Лестница имеет определенное количество ступенек N. Заяц может одним прыжком преодолеть не более К ступенек. Для разнообразия зайчик пытается каждый раз найти новый путь к вершине лестницы. Директору любопытно, сколько различных способов есть у зайца добраться до вершины лестницы при заданных значениях K и N. Помогите директору написать программу, которая поможет вычислить это количество. Например, если K=3 и N=4, то существуют следующие маршруты: 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2, 1+3, 3+1. Т.е. при данных значениях у зайца всего 7 различных маршрутов добраться до вершины лестницы. 

Для проверки: 1 и 3 = 1 вариант, 2 и 7 = 21 вариант, 3 и 10 = 274 варианта

## Example output

go run Rabbit.go -N 4 -K 3

    1+1+1+1
    1+1+2
    1+2+1
    1+3
    2+1+1
    2+2
    3+1
