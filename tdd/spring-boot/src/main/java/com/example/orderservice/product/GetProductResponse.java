package com.example.orderservice.product;

import lombok.Getter;

@Getter
class GetProductResponse {
    private long id;
    private String name;
    private int price;
    private DiscountPolicy discountPolicy;

    public GetProductResponse(long id, String name, int price, DiscountPolicy discountPolicy) {
        this.id = id;
        this.name = name;
        this.price = price;
        this.discountPolicy = discountPolicy;
    }

    static GetProductResponse fromEntity(Product product) {
        return new GetProductResponse(product.getId(), product.getName(), product.getPrice(), product.getDiscountPolicy());
    }
}
