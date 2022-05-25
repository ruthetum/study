package chap05.lifecycle;

import org.junit.jupiter.api.*;

public class LifeCycleTest {

    public LifeCycleTest() {
        System.out.println("new LifecycleTest");
    }

    @BeforeAll
    public static void beforeAll() { System.out.println("Before All");}

    @AfterAll
    public static void afterAll() { System.out.println("After All"); }

    @BeforeEach
    void setUp() { System.out.println("setUp"); }

    @Test
    void a() { System.out.println("A"); }

    @Test
    void b() { System.out.println("B"); }

    @AfterEach
    void tearDown() { System.out.println("tearDown"); }
}
