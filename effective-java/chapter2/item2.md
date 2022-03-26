# Item 2. 생성자에 매개변수가 많다면 빌더를 고려하라

- 점층적 생성자 패턴
- 자바빈즈 패턴
- <strong>빌더 패턴</strong>

## 점층적 생성자 패턴
- 매개변수가 늘어남에 따라 프로그래머들은 점층적 생성자 패턴을 즐겨 사용했다.

```java
public class NutritionFacts {
    private final int servingSize; // 필수
    private final int servings; // 필수
    private final int calories; // 선택
    private final int fat; // 선택
    private final int sodium; // 선택
    private final int carbohydrate; // 선택

    public NutritionFacts(int servingSize, int servings) { // 필수 2개
        this(servingSize, servings, 0);
    }

    public NutritionFacts(int servingSize, int servings, int calories) { // 필수 2개 + 선택 1개
        this(servingSize, servings, calories, 0);
    }

    public NutritionFacts(int servingSize, int servings, int calories, int fat) { // 필수 2개 + 선택 2개
        this(servingSize, servings, calories, fat, 0);
    }

    public NutritionFacts(int servingSize, int servings, int calories, int fat, int sodium) { // 필수 2개 + 선택 3개
        this(servingSize, servings, calories, fat, sodium, 0);
    }

    public NutritionFacts(int servingSize, int servings, int calories, int fat, int sodium, int carbohydrate) { // 필수 2개 +선택 4개
        this.servingSize = servingSize;
        this.servings = servings;
        this.calories = calories;
        this.fat = fat;
        this.sodium = sodium;
        this.carbohydrate = carbohydrate;
    }
}
```

- 이와 같이 정층적 생성자 패턴을 사용할 수는 있지만, 매개변수의 개수가 많아지면 클라이언트 코드를 작성하거나 읽기 어렵다.

## 자바빈즈 패턴 (JavaBeans Pattern)
- 선택 매개변수가 많을 때 활용할 수 있는 두 번째 대안은 자바빈즈 패턴이다.
- 매개변수가 없는 생성자로 객체를 만든 후 `setter` 메서드를 호출해 원하는 매개변수의 값을 설정한다.

```java
@Setter
public class NutritionFacts {
    // 매개변수들은 기본값이 있는 경우 기본값으로 초기화된다.
    private final int servingSize = -1; // 필수
    private final int servings = -1; // 필수
    private final int calories = 0; // 선택
    private final int fat = 0; // 선택
    private final int sodium = 0; // 선택
    private final int carbohydrate = 0; // 선택
}
```
  
```java
NutritionFacts coke = new NutritionFacts();
coke.setServingSize(240);
coke.setServings(8);
coke.setCalories(100);
coke.setSodium(35);
coke.setCarbohydrate(27);
```
  
- 자바빈즈 패턴을 활용하면 점층적 생성자 패턴에 비해 쉽게 인스턴스를 만들 수 있고, 가독성 좋은 코드를 작성할 수 있다.
- 하지만 자바빈즈 패턴에서 객체 하나를 만들려면 메서드를 여러 개 호출애햐 하고, 객체가 완전히 생성되기 전까지 일관성이 무너진다.
- 또한 자바빈즈 패턴에서는 클래스를 불변 클래스로 만들 수 없으며 스레드 안전성을 얻으러면 프로그래머가 추가 작업을 해줘야 한다.

## 빌더 패턴 (Builder Pattern)

```java
public class NutritionFacts {
    private final int servingSize;
    private final int servings;
    private final int calories;
    private final int fat;
    private final int sodium;
    private final int carbohydrate;

    public static class Builder {
        // 필수 매개변수
        private final int servingSize;
        private final int servings;

        // 선택 매개변수 - 기본값으로 초기화한다.
        private int calories = 0;
        private int fat = 0;
        private int sodium = 0;
        private int carbohydrate = 0;

        public Builder(int servingSize, int servings) {
            this.servingSize = servingSize;
            this.servings = servings;
        }

        public Builder calories(int val) {
            calories = val;
            return this;
        }

        public Builder fat(int val) {
            fat = val;
            return this;
        }

        public Builder sodium(int val) {
            sodium = val;
            return this;
        }

        public Builder carbohydrate(int val) {
            calories = val;
            return this;
        }
        
        public NutritionFacts build() {
            return new NutritionFacts(this);
        }
    }
    
    private NutritionFacts(Builder builder) {
        servingSize = builder.servingSize;
        servings = builder.servings;
        calories = builder.calories;
        fat = builder.fat;
        sodium = builder.sodium;
        carbohydrate = builder.carbohydrate;
    }
}
```
- 클라이언트는 필요한 객체를 직접 마드는 대신, 필수 매개변수만으로 생성자(혹은 정적 팩터리 메서드)를 호출해 빌더 객체를 얻는다.
- 그 다음 빌더 객체가 제공하는 세터 메서드들을 통해 선택 매개변수를 설정한다.
- 잘못된 매개변수를 일찍 발견하려면 빌더의 생성자와 메서드에서 입력 매개변수를 검사한다.
  - 매개변수를 잘못 입력했을 때 `IllegalArgumentException`을 던져준다.
- 빌더 패턴은 계층적으로 설계된 클래스와 함께 쓰기에 좋다.

## 정리
> - 생성자나 정적 팩터리가 처리해야 할 매개변수가 많다면 빌더 패턴을 선택하는 게 좋다.
> - 매개변수 중 다수가 필수가 아니거나 같은 타입이라면 특히 사용을 권장한다.
> - 리더는 점층적 생성자보다 코드를 읽고 쓰기에 간결하고, 자바빈즈보다 안전한다.
