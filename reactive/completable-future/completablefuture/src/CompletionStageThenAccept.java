import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;

/**
 * thenAccept, thenAcceptAsync sample code
 */
public class CompletionStageThenAccept {

    public static void thenAcceptAtFinishingStageExample() throws InterruptedException {
        CompletionStage<Integer> stage = FutureHelper.finishedStage();
        stage.thenAccept(i -> {
            System.out.printf("%d in thenAccept\n", i);
        }).thenAccept(i -> {
            System.out.printf("%d in thenAccept2\n", i);
        });
        System.out.println("after thenAccept\n");

        Thread.sleep(1000);
    }

    public static void thenAcceptAsyncAtFinishingStageExample() throws InterruptedException {
        CompletionStage<Integer> stage = FutureHelper.finishedStage();
        stage.thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync\n", i);
        }).thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync2\n", i);
        });
        System.out.println("after thenAccept\n");

        Thread.sleep(1000);
    }

    public static void thenAcceptAtRunningStageExample() throws InterruptedException {
        CompletionStage<Integer> stage = FutureHelper.runningStage();
        stage.thenAccept(i -> {
            System.out.printf("%d in thenAccept\n", i);
        }).thenAccept(i -> {
            System.out.printf("%d in thenAccept2\n", i);
        });
        System.out.println();

        Thread.sleep(1000);
    }

    public static void thenAcceptAsyncAtRunningStageExample() throws InterruptedException {
        CompletionStage<Integer> stage = FutureHelper.runningStage();
        stage.thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync\n", i);
        }).thenAcceptAsync(i -> {
            System.out.printf("%d in thenAcceptAsync2\n", i);
        });
        System.out.println();

        Thread.sleep(1000);
    }

    public static void main(String[] args) throws InterruptedException {
        thenAcceptAtFinishingStageExample();
        thenAcceptAsyncAtFinishingStageExample();
        thenAcceptAtRunningStageExample();
        thenAcceptAsyncAtRunningStageExample();
    }
}
