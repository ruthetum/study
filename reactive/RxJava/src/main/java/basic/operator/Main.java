package basic.operator;

import io.reactivex.Observable;

public class Main {
    public static void main(String[] args) {

        Observable<Integer> observable = Observable.just(1, 2, 3, 4, 5);

        // Map
        System.out.println("Map Result");
        observable.map(e -> e * 2).subscribe(System.out::println);

        // Filter
        System.out.println("Filter Result");
        observable.filter(e -> e % 2 == 0).subscribe(System.out::println);

        // Count
        System.out.println("Count Result");
        observable.count().subscribe(System.out::println);

        // Zip
        Observable<String> other = Observable.just("A", "B", "C");
        System.out.println("Zip Result");
        Observable.zip(
                observable,
                other,
                (x, y) -> x + y
        ).forEach(System.out::println);
    }
}
