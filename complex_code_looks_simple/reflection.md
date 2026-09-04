# Как код, выглядящий просто, может оказаться сложным

Нам дан пример:

```java
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateExample {
    public static void main(String[] args) {
        String dateString = "2024-05-13 14:30:00";
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        try {
            Date date = format.parse(dateString);
            System.out.println("Date: " + date);
        } catch (ParseException e) {
            e.printStackTrace();
        }
    }
}
```

Необходимо выяснить, какие есть недостатки у данного решения и сделать решение лучше.

Как выглядит на Go:

```go
package complexcodelookssimple

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2024-05-13 14:30:00"

	date, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		fmt.Printf("parse date: %v\n", err)
		panic("parse date")
	}

	fmt.Println("Date:", date)
}
```

Недостатки:

- формат даты вынес бы в отдельную переменную: "yyyy-MM-dd HH:mm:ss", либо использовал гоотвые константы (если такие есть в Java библиотеках)
- после нахождения исключения - мы сразу печатаем stackTrace и аварийно завершаем программу
- из-за аварийного завершения программы - мы не можем протестировать ошибочное поведение программы
- нет UTC - таймзоны, из-за этого будет подставлена UTC конкретной JVM - недетерменированное поведение

Улучшенный пример:

```go
package complexcodelookssimple

import (
	"fmt"
	"time"
)

const (
	inputDate = "2024-05-13 14:30:00"
	timeZone  = "Europe/Moscow"
)

func ParseDate(s string, loc *time.Location) (time.Time, error) {
	date, err := time.ParseInLocation(time.DateTime, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse date %q: %w", s, err)
	}

	return date, nil
}

func ExampleParseDate() {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Printf("load location %q: %v\n", timeZone, err)
		return
	}

	date, err := ParseDate(inputDate, loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Date:", date)
}
```
