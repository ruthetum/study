package com.example.orderservice.product;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static com.example.orderservice.product.ProductSteps.상품등록요청_생성;
import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
public class ProductServiceTest {

    @Autowired
    private ProductService productService;

    @Test
    void 상품조회() {
        // 상품 등록
        productService.addProduct(상품등록요청_생성());
        Long productId = 1L;

        // 상품 조회
        GetProductResponse response = productService.getProduct(productId);

        // 상품 응답 검증
        assertThat(response).isNotNull();
    }

    @Test
    void 상품수정() {
        // 상품 등록
        productService.addProduct(상품등록요청_생성());
        Long productId = 1L;

        UpdateProductRequest request = new UpdateProductRequest("초코에몽L", 2000, DiscountPolicy.NONE);
        productService.updateProduct(productId, request);

        // 검증
        assertThat(productService.getProduct(productId).getName()).isEqualTo("초코에몽L");
        assertThat(productService.getProduct(productId).getPrice()).isEqualTo(2000);
    }
}
