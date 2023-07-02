package com.example.orderservice.order.interfaces.dto;

import lombok.Getter;

@Getter
public class CreateOrderRequest {
    private Long productId;
    private int quantity;

    public CreateOrderRequest(Long productId, int quantity) {
        this.productId = productId;
        this.quantity = quantity;
    }
}
