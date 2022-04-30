package com.mvc.orderkafka.controller;

import com.mvc.orderkafka.domain.Order;
import com.mvc.orderkafka.dto.OrderRequest;
import com.mvc.orderkafka.dto.OrderResponse;
import com.mvc.orderkafka.kafka.KafkaProducer;
import com.mvc.orderkafka.kafka.event.PaymentRequest;
import com.mvc.orderkafka.service.OrderService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;
    private final KafkaProducer kafkaProducer;

    @PostMapping("/order")
    public OrderResponse order(@RequestBody OrderRequest request) {
        Order order = orderService.order(request);
        kafkaProducer.publishOrder(PaymentRequest.fromEntity(order));
        return OrderResponse.fromEntity(order);
    }
}
