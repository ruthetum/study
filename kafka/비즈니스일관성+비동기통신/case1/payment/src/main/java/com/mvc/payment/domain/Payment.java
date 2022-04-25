package com.mvc.payment.domain;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table(name = "payment")
@Getter
@Setter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class Payment {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "payment_id")
    private Long id;

    private Long userId;

    private Long orderId;

    private int price;

    private Boolean status;

    public static Payment createPayment(Long userId, Long orderId, int price) {
        Payment payment = new Payment();
        payment.setUserId(userId);
        payment.setOrderId(orderId);
        payment.setPrice(price);
        payment.setStatus(true);
        return payment;
    }

    public void failed() {
        this.status = false;
    }
}
