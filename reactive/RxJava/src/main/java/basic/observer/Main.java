package basic.observer;

import io.reactivex.Observable;
import io.reactivex.Observer;
import org.reactivestreams.Subscriber;
import org.reactivestreams.Subscription;

import java.util.concurrent.Executors;
import java.util.concurrent.Future;

/**
 * Observable을 생성하는 방식은 RxJava 1.2.7부터 사용되지 않음
 * 생성하는 것이 너무 많고, subscriber에게 과도한 부하를 줄 수 있기 때문에 안전하지 않음
 * 본 코드는 흐름을 이해하기 위한 기초 코드
 */

public class Main {
    public static void main(String[] args) {

        // Observer + Subscriber
        Observable.create(
                sub -> {
                    sub.onNext("Hello, World!");
                    sub.onComplete();
                }
        ).subscribe(
                System.out::println,                    // onNext()
                System.err::println,                    // onError()
                () -> System.out.println("Completed!")  // onCompleted()
        );

        // Iterable 컬렉션을 이용한 Observable 생성
        Observable.just("1", "2", "3", "4");
        Observable.fromArray(new String[]{"A", "B", "C"});

        // Callable, Future 활용한 Observable 생성
        Observable<String> hello = Observable.fromCallable(() -> "Hello, ");
        Future<String> future = Executors.newCachedThreadPool().submit(() -> "World");
        Observable<String> world = Observable.fromFuture(future);

        // concat()을 통해 입력 스트림을 다운 스트림 Observable로 보낼 수 있음
        Observable.concat(hello, world, Observable.just("!"))
                .forEach(System.out::print);
    }
}
