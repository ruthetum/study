# Item 45. 스트림은 주의해서 사용하라

## 스트림 API
- 다량의 데이터 처리를 돕고자 자바 8에 추가되었다.
- 스트림 API가 제공하는 추상 개념 중 핵심은 두 가지로 **스트림**과 **스트림 파이프라인**이다.
    - 스트림(stream) : 데이터 원소의 유한 혹은 무한 시퀀스를 의미
    - 스트림 파이프라인(stream pipeline) : 데이터들로 수행하는 연산단계를 표현하는 개념
- 스트림 API는 메서드 연쇄를 지원하는 플루언트 API이다.
    - 플루언트 API(fluent API) : 파이프 라인 하나를 구성하는 모든 호출을 연결하여 하나의 표현식으로 완성할 수 있는 API

## 스트림 API와 유지보수
- 스트림 API를 과용하면 프로그램이 읽거나 유지보수하기 어려워진다.
- 아래의 예(아나그램)를 통해 확인해보자.
    - 아나그램 : 'staple' -> 'aelpst'처럼 철자를 구성하는 알파벳이 같고 순서만 다른 언어

### 스트림을 활용하지 않았을 때
```java
public class Anagrams {
    public static void main(String[] args) throws IOException {
        File dictionary = new File(args[0]);
        int minGroupSize = Integer.parseInt(args[1]);

        Map<String, Set<String>> groups = new HashMap<>();
        try (Scanner s = new Scanner(dictionary)) {
            while (s.hasNext()) {
                String word = s.next();
                groups.computeIfAbsent(alphabetize(word), (unused) -> new TreeSet<>()).add(word);   // ⓐ
            }
        }

        for (Set<String> group : groups.values())
            if (group.size() >= minGroupSize)
                System.out.println(group.size() + ": " + group);
    }

    private static String alphabetize(String s) {
        char[] a = s.toCharArray();
        Arrays.sort(a);
        return new String(a);
    }
}
```
- ⓐ : **computeIfAbsent** 메서드를 활용해서 맵 안에 키가 있는지 찾은 다음, 있으면 단순히 그 키에 매핑된 값을 반환한다.

### 과한 스트림 사용했을 때
```java
public class Anagrams {
    public static void main(String[] args) throws IOException {
        Path dictionary = Paths.get(args[0]);
        int minGroupSize = Integer.parseInt(args[1]);

        try (Stream<String> words = Files.lines(dictionary)) {
            words.collect(
                    groupingBy(word -> word.chars().sorted()
                            .collect(StringBuilder::new,
                                    (sb, c) -> sb.append((char) c),
                                    StringBuilder::append).toString()))
                    .values().stream()
                    .filter(group -> group.size() >= minGroupSize)
                    .map(group -> group.size() + ": " + group)
                    .forEach(System.out::println);
        }
    }
}
```
- 스트림을 과하게 사용하는 경우 코드는 짧아지지만 가독성 및 유지보수하기 어려워진다.
- 따라서 너무 과하게 사용하기 보다는 적절히 활용하여 코드를 깔끔하고 명료하게 만들 수 있다.

### 스트림을 적절히 활용했을 때
```java
public class Anagrams {
    public static void main(String[] args) throws IOException {
        Path dictionary = Paths.get(args[0]);
        int minGroupSize = Integer.parseInt(args[1]);

        try (Stream<String> words = Files.lines(dictionary)) {                  // ⓐ
            words.collect(groupingBy(word -> alphabetize(word)))                // ⓑ
                    .values().stream()                                       
                    .filter(group -> group.size() >= minGroupSize)             
                    .forEach(g -> System.out.println(g.size() + ": " + g));
        }
    }

    private static String alphabetize(String s) {
        char[] a = s.toCharArray();
        Arrays.sort(a);
        return new String(a);
    }
}
```
- ⓐ : try-with-resource 블록에서 사전 파일을 열고, 파일의 모든 라인으로 구성된 스트림을 얻는다.
- ⓑ : 스트림 변수의 이름을 **words**로 지어 스트림 안의 각 원소가 단어(word)임을 명확히 한다.

### char 스트림
- 위의 코드에서 **alphabetize** 메서드도 스트림을 사용하여 다르게 구현할 수 있다.
    - 하지만 그렇게 할 경우 명확성이 떨어지고 잘못 구현할 가능성이 커진다.
    - 심지어 더 느려질 수도 있다.
- 자바는 기본 타입인 char용 스트림을 지원하지 않는다.
- 아래의 예처럼 char 값들을 스트림으로 처리하는 코드가 있을 때

```java
"Hello world!".chars().forEach(System.out::print);
```

- `Hello world!`가 출력될 것으로 기대되지만 `721011091091113211911111410810033`을 출력한다.
    - `"Hello world!".chars()`가 반환하는 스트림 원소는 char가 아닌 int 값이기 때문이다.
- 따라서 이를 올바르게 출력하기 위해서는 아래와 같이 형변환을 명시적으로 해줘야 한다.

```java
"Hello world!".chars().forEach(x -> System.out.print((char) x));
```

- 되도록이면 char값을 처리할 때는 스트림을 삼가는 편이 좋다.

## 스트림 API를 적용하기 좋은 상황
- 원소들의 시퀀스를 일관되게 변환한다.
- 원소들의 시퀀스를 필터링한다.
- 원소들의 시퀀스를 하나의 연산을 사용해 결합한다.
- 원소들의 시퀀스를 컬렉션에 모은다.
- 원소들의 시퀀스에서 특정 조건을 만족하는 원소를 찾는다.

## 스트림 API를 적용하기 어려운 상황
- 한 데이터가 파이프라인의 여러 단계를 통과할 때 이 데이타의 각 단계에서 값들에 동시에 접근하기 어려운 경우
    - 스트림 파이프라인은 일단 한 값을 다른 값에 매핑하고 나면 원래의 값을 잃는 구조이기 때문이다.

## 정리
- 스트림을 사용해야 효율적으로 처리할 수 있는 일이 있고, 단순하게 반복문으로 처리하는 것이 효율적일 때도 있다.
- 그리고 많은 작업들이 이 둘을 조합했을 때 가장 효율적으로 해결된다.
- 어느 방법을 선택하는지에 대한 확실한 규칙은 없지만 참고할만한 지침이 있기 때문에 해당 지침들을 선행적으로 참고해야 한다.
- 스트림과 반복 중 어느 쪽이 나은지 확신하기 어려운 경우에는 둘 다 해보고 더 나은 방법을 선택해야 한다.