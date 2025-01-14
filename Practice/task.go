func solution(arr []string, n int) int {
  // Словарь для хранения уже обработанных последовательностей
  processed := make(map[string]bool)
  // Счетчик повторяющихся последовательностей
  count := 0

  // Проходим по массиву раздвижным окном
  for i := 0; i <= len(arr)-n; {
    // Получаем текущую последовательность длины n
    seq := strings.Join(arr[i:i+n], ",")

    // Проверяем, была ли эта последовательность уже обработана
    if processed[seq] {
      i++
      continue
    }

    // Отмечаем текущую последовательность как обработанную
    processed[seq] = true

    // Считаем количество повторений этой последовательности начиная с текущей позиции
    repeats := 0
    for j := i; j <= len(arr)-n; j += n {
      if strings.Join(arr[j:j+n], ",") == seq {
        repeats++
      } else {
        break
      }
    }

    // Добавляем количество повторений минус одно (т.к. первая последовательность не считается повтором)
    if repeats > 1 {
      count += repeats - 1
    }

    // Сдвигаем указатель на позицию после последнего повторения
    i += repeats * n
  }

  return count
}

func main() {
  arr := []string{"A", "B", "A", "B", "A", "B", "A", "C", "A", "B", "A", "B", "A", "D", "A", "E", "A", "E", "A", "B", "A", "B", "A", "B", "A"}
  n := 2
  log.Println(solution(arr, n)) // Ожидаемый результат: 9
}
