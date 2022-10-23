package strategy.behavior.fly;

public class FlyNoWay implements FlyBehavior {
    public void fly() {
        System.out.println("저는 못 날아요.");
    }
}
