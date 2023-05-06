import java.util.function.Consumer;
import java.util.stream.Stream;

public class MethodReference {
    public static class Person {
        private final String name;

        public Person(String name) {
            this.name = name;
        }

        public String getName() {
            return name;
        }

        public Boolean compareTo(Person p) {
            return p.name.compareTo(name) > 0;
        }
    }

    public static void print(String name) {
        System.out.println(name);
    }

    public static void main(String[] args) {
        Person target = new Person("f");
        Consumer<String> staticPrint = MethodReference::print;

        Stream.of("a", "b", "c", "d")
                .map(Person::new)           // constructor reference
                .filter(target::compareTo)  // method reference
                .map(Person::getName)       // instance method reference
                .forEach(staticPrint);      // static method reference
    }
}
