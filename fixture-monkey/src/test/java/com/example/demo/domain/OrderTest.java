package com.example.demo.domain;

import com.navercorp.fixturemonkey.FixtureMonkey;
import com.navercorp.fixturemonkey.api.introspector.FieldReflectionArbitraryIntrospector;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.BDDAssertions.then;

public class OrderTest {

    @Test
    void order() {
        // given
        FixtureMonkey sut = FixtureMonkey.builder()
                .objectIntrospector(FieldReflectionArbitraryIntrospector.INSTANCE)
                .build();

        // when
        Order actual = sut.giveMeBuilder(Order.class)
                .set("orderNo", "1")
                .set("productName", "Line Sally")
                .minSize("items", 1)
                .sample();

        // then
        then(actual.getOrderNo()).isEqualTo("1");
        then(actual.getProductName()).isEqualTo("Line Sally");
        then(actual.getItems()).hasSizeGreaterThanOrEqualTo(1);
    }
}
