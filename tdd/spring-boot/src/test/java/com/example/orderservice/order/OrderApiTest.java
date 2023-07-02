package com.example.orderservice.order;

import com.example.orderservice.ApiTest;
import io.restassured.response.ExtractableResponse;
import io.restassured.response.Response;
import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;

import static com.example.orderservice.product.ProductSteps.상품등록요청;
import static com.example.orderservice.product.ProductSteps.상품등록요청_생성;
import static org.assertj.core.api.Assertions.assertThat;

public class OrderApiTest extends ApiTest {

    @Test
    void 주문생성() {
        // given
        상품등록요청(상품등록요청_생성());
        CreateOrderRequest request = OrderSteps.주문생성요청_생성();

        // when
        ExtractableResponse<Response> response = OrderSteps.주문생성요청(request);

        // then
        assertThat(response.statusCode()).isEqualTo(HttpStatus.CREATED.value());
    }
}
