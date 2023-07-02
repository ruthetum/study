package com.example.orderservice.product;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
class ProductService {
    private final ProductPort productPort;

    ProductService(ProductPort productPort) {
        this.productPort = productPort;
    }

    @Transactional
    public void addProduct(AddProductRequest request) {
        final Product product = new Product(request.name(), request.price(), request.discountPolicy());
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
