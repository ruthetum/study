package com.example.orderservice.product.application.port;

import com.example.orderservice.product.domain.Product;

interface ProductPort {
    void save(Product product);

    Product getProduct(Long productId);
}
