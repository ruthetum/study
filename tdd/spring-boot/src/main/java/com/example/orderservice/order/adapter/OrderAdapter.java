package com.example.orderservice.order.adapter;

import com.example.orderservice.order.application.port.OrderPort;
import com.example.orderservice.order.domain.Order;
import com.example.orderservice.order.domain.OrderRepository;
import com.example.orderservice.product.domain.Product;
import com.example.orderservice.product.domain.ProductRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class OrderAdapter implements OrderPort {

    private final ProductRepository productRepository;
    private final OrderRepository orderRepository;

    @Override
    public Product getProductById(Long productId) {
        return productRepository.findById(productId)
                .orElseThrow(() -> new IllegalArgumentException("상품이 존재하지 않습니다."));
    }

    @Override
    public void save(Order order) {
        orderRepository.save(order);
    }
}
