package com.example.orderservice.order.application.port;

import com.example.orderservice.order.domain.Order;
import com.example.orderservice.product.domain.Product;

public interface OrderPort {
    Product getProductById(Long productId);

    void save(Order order);
}
