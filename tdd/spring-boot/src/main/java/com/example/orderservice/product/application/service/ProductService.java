package com.example.orderservice.product.application.service;

import com.example.orderservice.product.domain.Product;
import com.example.orderservice.product.ProductPort;
import com.example.orderservice.product.interfaces.dto.AddProductRequest;
import com.example.orderservice.product.interfaces.dto.GetProductResponse;
import com.example.orderservice.product.interfaces.dto.UpdateProductRequest;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class ProductService {
    private final ProductPort productPort;

    ProductService(ProductPort productPort) {
        this.productPort = productPort;
    }

    @Transactional
    public void addProduct(AddProductRequest request) {
        final Product product = Product.create(request.getName(), request.getPrice(), request.getDiscountPolicy());
        productPort.save(product);
    }

    public GetProductResponse getProduct(Long productId) {
        Product product = productPort.getProduct(productId);
        return GetProductResponse.fromEntity(product);
    }

    @Transactional
    public void updateProduct(Long productId, UpdateProductRequest request) {
        Product product = productPort.getProduct(productId);
        product.update(request.getName(), request.getPrice(), request.getDiscountPolicy());
        productPort.save(product);
    }
}
