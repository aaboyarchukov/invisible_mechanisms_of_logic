# "Знание" кода

Многие говорят, что одна часть кода знает о другой. Здесь имеют в виду скрытую зависимость, вот в чем она может проявляться:

Речь идёт прежде всего о том, что такие элементы программы, как модули, функции или классы, **зависят друг от друга**. Это зависимость может выражаться в нескольких формах:

**Прямые вызовы**: Когда одна функция вызывает другую напрямую, она "знает" о её существовании и сигнатуре.

**Обмен данными**: Когда данные, созданные или изменённые в одной части кода, передаются и используются в другой.

**Общие глобальные переменные**: Когда разные части программы обращаются к одним и тем же глобальным переменным.

**Наследование и полиморфизм**: Когда один класс наследует свойства и методы другого.

**Общие библиотеки и модули**: Когда несколько частей программы используют одни и те же внешние библиотеки или модули.

**Интерфейсы и контракты**: Предопределённые соглашения о том, как части кода будут взаимодействовать друг с другом.

Такое "знание" создаёт зависимости между компонентами, и важно управлять ними, чтобы поддерживать понимаемость и сопровождаемость кода.

## Примеры

1. Переопределение удаленного метода

```java
class Animal {
    public void makeSound() {
        System.out.println("Some generic animal sound");
    }
}

class Cat extends Animal {
    // Переопределение метода makeSound
    public void makeSound() {
        System.out.println("Meow");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal myCat = new Cat();
        myCat.makeSound();  // "Meow"
    }
}
```

Но

```java
class Animal {
    // Изменен метод в суперклассе
    public void makeGenericSound() {
        System.out.println("Some generic animal sound");
    }
}
```

В Java методы определяются по типу, то есть у Animal - будет отсутствовать нужный метод и будет ошибка. Здесь задействован один из механизмов логики - зависимость за счет наследования. То есть, когда мы зависим от родительского класса, мы не можем гарантировать что он не изменится, из-за этого, может происходить такое поведение.

2. Переписывание несуществующего метода родительского класса

```java
class Animal {
    public void makeSound() {
        System.out.println("Some generic animal sound");
    }
}

class Cat extends Animal {
    @Override
    public void makeSound(int numberOfSounds) {
        for (int i = 0; i < numberOfSounds; i++) {
            System.out.println("Meow");
        }
    }

    @Override
    public void makeSound() {
        System.out.println("Meow");
    }
}

public class Main {
    public static void main(String[] args) {
        Animal cat = new Cat();
        cat.makeSound();
        cat.makeSound(3);
    }
}
```

Здесь похожая ошибка и на первый пример, так как мы смотрим на тип Animal - получается мы пытаемся переписать несуществующий тип, что ошибочно. Здесь задействован один из механизмов логики - зависимость за счет наследования

3. Перекодирование данных

```java
/*
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.9.10</version>
</dependency>
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.12.5</version>
</dependency>
*/

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        // Создаем объект ObjectMapper для парсинга JSON
        ObjectMapper objectMapper = new ObjectMapper();

        String jsonString = "{\"name\":\"John\", \"age\":30}";

        try {
            // Парсим JSON-строку в HashMap
            Map<String, Object> result = objectMapper.readValue(jsonString, HashMap.class);

            System.out.println("Name: " + result.get("name"));
        } catch (IOException e) {
            // Обработка ошибки парсинга
            e.printStackTrace();
        }

        try {
            String prettyJson = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(result);
            System.out.println("Pretty JSON: " + prettyJson);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

Здесь честно не совсем понимаю, что происходит, но судя по функциям, сначала мы определили строку в `HashMap`, но результат типа `Map` - который является родительским (но здесь мы "сужаем" наш тип, поднимаясь по иерархии, с этим пропадают свойства и методы, который специфичны для `HashMap`). Далее мы из `Map` переводим уже в строку обратно, но подозреваю, что будет ошибка, так как обычный `Map` к строке не приведешь.
Здесь также зависимость от родительского класса.

На что еще не обратил внимание - `result` находится в другом блоке `try`, соответсвенно у них разная видимость, здесь один из механизмов - общие переменные.
