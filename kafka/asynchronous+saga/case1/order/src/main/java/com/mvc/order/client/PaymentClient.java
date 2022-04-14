package com.mvc.order.client;

import com.mvc.order.dto.PaymentRequest;
import com.mvc.order.dto.PaymentResponse;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "PaymentFeign",
        url = "http://localhost:9001",
        fallback = PaymentClientFallback.class)
public interface PaymentClient {

    @PostMapping("/payment")
    PaymentResponse pay(@RequestBody PaymentRequest request);
}
