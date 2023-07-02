package com.example.orderservice.order.application.service;

import com.example.orderservice.order.application.port.OrderPort;
import com.example.orderservice.order.domain.Order;
import com.example.orderservice.order.interfaces.dto.CreateOrderRequest;
import com.example.orderservice.product.domain.Product;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@RequiredArgsConstructor
public class OrderService {
    private final OrderPort orderPort;

    @Transactional
    public void createOrder(CreateOrderRequest request) {
        Long productId = request.getProductId();

        Product product = orderPort.getProductById(productId);

        Order order = Order.create(product, request.getQuantity());
        orderPort.save(order);
    }
}
