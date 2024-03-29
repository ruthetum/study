package com.example.orderservice.product;

import com.example.orderservice.product.domain.DiscountPolicy;
import com.example.orderservice.product.domain.Product;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

public class ProductTest {

    @Test
    void update() {
        Product product = Product.create("초코에몽", 1000, DiscountPolicy.NONE);

        product.update("초코에몽L", 2000, DiscountPolicy.NONE);

        assertThat(product.getName()).isEqualTo("초코에몽L");
        assertThat(product.getPrice()).isEqualTo(2000);
    }
}
