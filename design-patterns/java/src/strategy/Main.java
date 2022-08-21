package strategy;

import strategy.behavior.fly.FlyRocketPowered;
import strategy.duck.Duck;
import strategy.duck.MallardDuck;
import strategy.duck.ModelDuck;

public class Main {
    public static void main(String[] args) {
        Duck mallard = new MallardDuck();
        mallard.performQuack();
        mallard.performFly();

        Duck model = new ModelDuck();
        model.performFly();
        // 실행 중에 오리의 행동을 바꾸고 싶으면 setter 메서드를 통해 수정
        model.setFlyBehavior(new FlyRocketPowered());
        model.performFly();
    }
}
