package com.mvc.order.client;

import com.mvc.order.dto.PaymentRequest;
import com.mvc.order.dto.PaymentResponse;

public class PaymentClientFallback implements PaymentClient {

    @Override
    public PaymentResponse pay(PaymentRequest request) {
        return null;
    }
}
