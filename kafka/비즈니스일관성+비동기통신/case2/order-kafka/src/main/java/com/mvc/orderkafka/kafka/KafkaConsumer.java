package com.mvc.orderkafka.kafka;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.mvc.orderkafka.domain.Order;
import com.mvc.orderkafka.kafka.event.PaymentResponse;
import com.mvc.orderkafka.repository.OrderRepository;
import com.mvc.orderkafka.service.OrderService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
@RequiredArgsConstructor
public class KafkaConsumer {

    private final ObjectMapper objectMapper;
    private final OrderRepository orderRepository;

    @KafkaListener(topics = "${spring.kafka.topic.payment}")
    @Transactional
    public void subscribePayment(String message) {
        try {
            PaymentResponse paymentResponse = objectMapper.readValue(message, PaymentResponse.class);

            Order order = orderRepository.findById(paymentResponse.getOrderId())
                    .orElseThrow(() -> new RuntimeException());

            if (paymentResponse.getState()) {
                order.payment();
                return;
            }
            order.failed();
        } catch (Exception e){
            e.printStackTrace();
            throw new RuntimeException();
        }
    }
}
