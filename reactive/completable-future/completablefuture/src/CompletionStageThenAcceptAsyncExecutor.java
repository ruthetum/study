import java.util.concurrent.CompletionStage;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class CompletionStageThenAcceptAsyncExecutor {

    public static void main(String[] args) throws InterruptedException {
        ExecutorService single = Executors.newSingleThreadExecutor();
        ExecutorService fixed = Executors.newFixedThreadPool(10);

        System.out.println("start main");
        CompletionStage<Integer> stage = FutureHelper.runningStage();
        stage.thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync\n", i);
        }, fixed).thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync2\n", i);
        }, single);
        System.out.println("after thenAccept");
        Thread.sleep(200);

        single.shutdown();
        fixed.shutdown();
    }
}
