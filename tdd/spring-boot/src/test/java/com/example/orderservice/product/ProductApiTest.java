package com.example.orderservice.product;

import com.example.orderservice.ApiTest;
import io.restassured.RestAssured;
import io.restassured.response.ExtractableResponse;
import io.restassured.response.Response;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;

import static com.example.orderservice.product.ProductSteps.*;
import static org.assertj.core.api.Assertions.assertThat;

class ProductApiTest extends ApiTest {

    @Autowired
    private ProductRepository productRepository;
    @Test
    void 상품등록() {
        final AddProductRequest request = 상품등록요청_생성();

        // API 요청
        final ExtractableResponse<Response> response = 상품등록요청(request);

        assertThat(response.statusCode()).isEqualTo(HttpStatus.CREATED.value());
    }

    @Test
    void 상품조회() {
        // 상품 등록
        상품등록요청(상품등록요청_생성());
        Long productId = 1L;

        final ExtractableResponse<Response> response = 상품조회요청(productId);

        assertThat(response.statusCode()).isEqualTo(HttpStatus.OK.value());
        assertThat(response.jsonPath().getString("name")).isEqualTo("초코에몽");
    }

    @Test
    void 상품수정() {
        // 상품 등록
        상품등록요청(상품등록요청_생성());
        Long productId = 1L;
        final UpdateProductRequest request = 상품수정요청_생성();

        final ExtractableResponse<Response> response = 상품수정요청(productId, request);

        assertThat(response.statusCode()).isEqualTo(HttpStatus.OK.value());
        assertThat(productRepository.findById(productId).get().getName()).isEqualTo("초코에몽L");
        assertThat(productRepository.findById(productId).get().getPrice()).isEqualTo(2000);
    }
}
