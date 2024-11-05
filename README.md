[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)

# maze
# Консольная игра "Лабиринт"

---
## Описание проекта
Проект представляет реализацию консольной программы для генерации лабиринта и поиска пути в нём. Программа генерирует лабиринт, алгоритм для генерации выбирается пользователем при запуске программы (алгоритм Прима или алгоритм Краскала), различного размера, размер также указывается пользователем (ширина и высота лабиринта), а также показывает поиск пути в лабиринте от точки S (начало лабиринта) к точке E (конец лабиринта) алгоритмами BFS или A*

---
## Технологии
* Go 1.23.0
* DDD (Domain Driven Design)

---
## Запуск игры

**1. Клонировать репозиторий:**
```
git clone https://github.com/KazikovAP/maze
```

**2. Запустить игру:**
```
go run cmd/maze/main.go
```

## Пример игры
```
Welcome to the Maze Game!
Enter maze width (default: 15): 
"Invalid width input, using default = 15"
Enter maze height (default: 15):         
"Invalid height input, using default = 15"
Choose maze generation algorithm (1 - Prim, 2 - Kruskal, default Prim): 
"Invalid choice, using default Prim generation algorithm"    
Choose pathfinding algorithm (1 - BFS, 2 - A*, default BFS): 
"Invalid choice, using default BFS pathfinding algorithm"        
Generating a 15x15 maze using Prim generation and BFS pathfinding
Start: {1 2}, End: {13 9}
Path found using BFS!

▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓   ▓ ▓ ▓▓▓   ▓
▓S▓      ▓ ▓ ▓▓
▓ ▓ ▓▓ ▓      ▓
▓  ▓ ▓  ▓ ▓▓ ▓▓
▓▓    ▓ ▓   ▓ ▓
▓  ▓ ▓  ▓ ▓   ▓
▓▓  ▓▓ ▓  ▓ ▓ ▓
▓  ▓  ▓▓ ▓   ▓▓
▓ ▓ ▓   ▓  ▓ E▓
▓     ▓  ▓  ▓ ▓
▓ ▓ ▓▓▓ ▓▓▓ ▓ ▓
▓▓     ▓ ▓ ▓  ▓
▓▓▓ ▓▓      ▓ ▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
```
```
Welcome to the Maze Game!
Enter maze width (default: 15): 25
Enter maze height (default: 15): 25
Choose maze generation algorithm (1 - Prim, 2 - Kruskal, default Prim): 2
Choose pathfinding algorithm (1 - BFS, 2 - A*, default BFS): 2
Generating a 25x25 maze using Kruskal generation and A* pathfinding
Start: {1 10}, End: {23 20}
Path found using BFS!

▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓     ▓     ▓   ▓ ▓   ▓ ▓
▓▓▓ ▓ ▓ ▓ ▓▓▓ ▓ ▓ ▓ ▓▓▓ ▓
▓   ▓ ▓ ▓ ▓ ▓ ▓   ▓     ▓
▓ ▓▓▓ ▓ ▓▓▓ ▓▓▓ ▓▓▓▓▓ ▓▓▓
▓ ▓     ▓ ▓ ▓ ▓     ▓   ▓
▓▓▓ ▓ ▓▓▓ ▓ ▓ ▓▓▓ ▓▓▓▓▓ ▓
▓ ▓ ▓   ▓   ▓ ▓     ▓   ▓
▓ ▓▓▓ ▓ ▓▓▓ ▓ ▓▓▓▓▓ ▓ ▓ ▓
▓   ▓ ▓       ▓     ▓ ▓ ▓
▓S▓ ▓▓▓▓▓ ▓▓▓ ▓ ▓▓▓▓▓ ▓ ▓
▓ ▓ ▓   ▓ ▓   ▓       ▓ ▓
▓ ▓ ▓▓▓ ▓▓▓ ▓▓▓▓▓ ▓ ▓ ▓▓▓
▓     ▓         ▓ ▓ ▓   ▓
▓ ▓▓▓▓▓▓▓▓▓ ▓▓▓ ▓ ▓▓▓▓▓▓▓
▓ ▓         ▓ ▓     ▓   ▓
▓ ▓▓▓ ▓▓▓▓▓ ▓ ▓ ▓ ▓▓▓▓▓ ▓
▓     ▓       ▓ ▓ ▓     ▓
▓ ▓▓▓▓▓ ▓ ▓ ▓ ▓▓▓ ▓▓▓ ▓▓▓
▓ ▓     ▓ ▓ ▓ ▓         ▓
▓ ▓▓▓▓▓ ▓▓▓ ▓ ▓▓▓▓▓ ▓ ▓E▓
▓     ▓   ▓ ▓   ▓   ▓ ▓ ▓
▓ ▓ ▓ ▓▓▓▓▓▓▓ ▓▓▓ ▓ ▓▓▓ ▓
▓ ▓ ▓   ▓     ▓   ▓     ▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
```
```
Welcome to the Maze Game!
Enter maze width (default: 15): 20
Enter maze height (default: 15): 20
Choose maze generation algorithm (1 - Prim, 2 - Kruskal, default Prim): 1
Choose pathfinding algorithm (1 - BFS, 2 - A*, default BFS): 2
Generating a 20x20 maze using Prim generation and A* pathfinding
Start: {1 2}, End: {18 2}
Path found using BFS!

▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓         ▓  ▓  ▓  ▓
▓S▓ ▓ ▓ ▓   ▓ ▓   E▓
▓▓  ▓▓ ▓▓ ▓     ▓  ▓
▓ ▓      ▓▓▓ ▓▓▓▓ ▓▓
▓ ▓ ▓ ▓▓  ▓▓     ▓ ▓
▓    ▓   ▓   ▓▓▓▓  ▓
▓ ▓ ▓ ▓ ▓  ▓     ▓ ▓
▓  ▓  ▓ ▓▓  ▓ ▓ ▓  ▓
▓ ▓ ▓ ▓   ▓▓▓▓    ▓▓
▓      ▓ ▓  ▓ ▓ ▓  ▓
▓ ▓ ▓ ▓    ▓     ▓ ▓
▓▓  ▓  ▓ ▓ ▓ ▓▓ ▓  ▓
▓  ▓  ▓▓  ▓  ▓   ▓ ▓
▓ ▓ ▓   ▓ ▓ ▓  ▓  ▓▓
▓   ▓ ▓ ▓ ▓  ▓ ▓▓  ▓
▓▓ ▓  ▓  ▓  ▓    ▓ ▓
▓   ▓ ▓ ▓ ▓ ▓▓ ▓  ▓▓
▓ ▓ ▓ ▓   ▓  ▓  ▓  ▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓

```

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)
