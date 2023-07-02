package com.example.orderservice.order;

import com.example.orderservice.product.Product;

public interface OrderPort {
    Product getProductById(Long productId);

    void save(Order order);
}
