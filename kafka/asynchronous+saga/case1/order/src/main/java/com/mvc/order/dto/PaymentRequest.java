package com.mvc.order.dto;

import com.mvc.order.domain.Order;
import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class PaymentRequest {
    private Long userId;
    private Long orderId;
    private int price;

    public static PaymentRequest fromEntity(Order order) {
        return PaymentRequest.builder()
                .userId(order.getUserId())
                .orderId(order.getId())
                .price(order.getTotalPrice())
                .build();
    }
}
