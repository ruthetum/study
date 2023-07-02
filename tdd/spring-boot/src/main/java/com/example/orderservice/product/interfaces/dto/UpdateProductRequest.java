package com.example.orderservice.product.interfaces.dto;

import com.example.orderservice.product.DiscountPolicy;
import lombok.Getter;

@Getter
public class UpdateProductRequest {

    private String name;
    private int price;
    private DiscountPolicy discountPolicy;

    public UpdateProductRequest(String name, int price, DiscountPolicy discountPolicy) {
        this.name = name;
        this.price = price;
        this.discountPolicy = discountPolicy;
    }
}
