package com.example.orderservice.product;

interface ProductPort {
    void save(Product product);

    Product getProduct(Long productId);
}
