package com.mvc.orderkafka.kafka.event;

import com.mvc.orderkafka.domain.Order;
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
