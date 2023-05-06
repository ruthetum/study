import lombok.SneakyThrows;

import java.util.concurrent.*;

public class FutureHelper {

    // getFuture: 새로운 스레드를 생성하여 1을 반환
    public static Future<Integer> getFuture() {
        ExecutorService executor = Executors.newSingleThreadExecutor();

        try {
            return executor.submit(() -> {
                return 1;
            });
        } finally {
            executor.shutdown();
        }
    }

    // getFuture: 새로운 스레드를 생성하고 1초 대기 후 1을 반환
    public static Future<Integer> getFutureCompleteAfter1s() {
        ExecutorService executor = Executors.newSingleThreadExecutor();

        try {
            return executor.submit(() -> {
                Thread.sleep(1000);
                return 1;
            });
        } finally {
            executor.shutdown();
        }
    }

    // finishedStage: 1을 반환하는 완료된 CompletableFuture 반환
    @SneakyThrows
    public static CompletionStage<Integer> finishedStage() {
        CompletionStage<Integer> future = CompletableFuture.supplyAsync(() -> {
            System.out.println("supplyAsync");
            return 1;
        });
        Thread.sleep(100);
        return future;
    }

    // runningStage: 1초를 sleep한 후 1을 반환하는 CompletableFuture
    public static CompletionStage<Integer> runningStage() {
        return CompletableFuture.supplyAsync(() -> {
            try {
                Thread.sleep(1000);
                System.out.println("I'm running!");
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
            return 1;
        });
    }
}

