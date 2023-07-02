package com.example.orderservice.order;

import io.restassured.RestAssured;
import io.restassured.response.ExtractableResponse;
import io.restassured.response.Response;
import org.springframework.http.MediaType;

public class OrderSteps {

    static CreateOrderRequest 주문생성요청_생성() {
        Long productId = 1L;
        int quantity = 3;
        return new CreateOrderRequest(productId, quantity);
    }

    static ExtractableResponse<Response> 주문생성요청(CreateOrderRequest request) {
        return RestAssured.given().log().all()
                .contentType(MediaType.APPLICATION_JSON_VALUE)
                .body(request)
                .when()
                .post("/orders")
                .then()
                .log().all().extract();
    }
}
