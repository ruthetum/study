package com.example.demo.domain;

import lombok.Getter;
import lombok.NoArgsConstructor;

import java.time.Instant;
import java.util.List;

import static lombok.AccessLevel.PROTECTED;

@Getter
@NoArgsConstructor(access = PROTECTED)
public class Order {
    Long id;

    String orderNo;

    String productName;

    int quantity;

    long price;

    List<String> items;

    Instant orderedAt;

    public static Order of(String orderNo, String productName, int quantity, long price, List<String> items, Instant orderedAt) {
        Order order = new Order();
        order.orderNo = orderNo;
        order.productName = productName;
        order.quantity = quantity;
        order.price = price;
        order.items = items;
        order.orderedAt = orderedAt;
        return order;
    }
}
