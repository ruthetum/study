package basic.asyncsequence;

import io.reactivex.Observable;

import java.util.concurrent.TimeUnit;

public class Main {

    public static void main(String[] args) throws InterruptedException {
        Observable.interval(1, TimeUnit.SECONDS)
                .subscribe(e -> System.out.println("Received: " + e));
        Thread.sleep(5000);
    }
}
