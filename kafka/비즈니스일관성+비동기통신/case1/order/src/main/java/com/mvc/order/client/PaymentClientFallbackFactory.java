package com.mvc.order.client;

import com.mvc.order.dto.PaymentRequest;
import com.mvc.order.dto.PaymentResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.openfeign.FallbackFactory;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class PaymentClientFallbackFactory implements FallbackFactory<PaymentClient> {

    @Override
    public PaymentClient create(Throwable cause) {
        return new PaymentClient() {
            @Override
            public PaymentResponse pay(PaymentRequest request) {
                log.error(cause.toString());

                // cause를 통해 정확한 exception 처리를 실행할 수 있음

                return null;
            }
        };
    }
}
