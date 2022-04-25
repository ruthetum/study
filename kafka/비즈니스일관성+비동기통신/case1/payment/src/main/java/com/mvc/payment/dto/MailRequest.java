package com.mvc.payment.dto;

import com.mvc.payment.domain.Payment;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class MailRequest {
    private Long userId;
    private Long orderId;
    private int price;

    public static MailRequest fromPayment(Payment payment) {
        MailRequest mailRequest = new MailRequest();
        mailRequest.setUserId(payment.getUserId());
        mailRequest.setOrderId(payment.getOrderId());
        mailRequest.setPrice(payment.getPrice());
        return mailRequest;
    }
}
