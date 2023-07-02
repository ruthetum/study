package com.example.orderservice.product;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/products")
public class ProductController {

    private final ProductService productService;

    ProductController(ProductService productService) {
        this.productService = productService;
    }

    @PostMapping
    public ResponseEntity<Void> addProduct(@RequestBody AddProductRequest request) {
        productService.addProduct(request);
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @GetMapping("/{productId}")
    public ResponseEntity<GetProductResponse> getProduct(@PathVariable Long productId) {
        GetProductResponse response = productService.getProduct(productId);
        return ResponseEntity.ok(response);
    }

    @PutMapping("/{productId}")
    public ResponseEntity<Void> updateProduct(@PathVariable Long productId, @RequestBody UpdateProductRequest request) {
        productService.updateProduct(productId, request);
        return ResponseEntity.ok().build();
    }
}
