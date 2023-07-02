package com.example.orderservice.product;

import com.example.orderservice.product.domain.DiscountPolicy;
import com.example.orderservice.product.interfaces.dto.AddProductRequest;
import com.example.orderservice.product.interfaces.dto.UpdateProductRequest;
import io.restassured.RestAssured;
import io.restassured.response.ExtractableResponse;
import io.restassured.response.Response;
import org.springframework.http.MediaType;

public class ProductSteps {
    public static AddProductRequest 상품등록요청_생성() {
        final String name = "초코에몽";
        final int price = 1000;
        final DiscountPolicy discountPolicy = DiscountPolicy.NONE;
        return new AddProductRequest(name, price, discountPolicy);
    }

    public static ExtractableResponse<Response> 상품등록요청(AddProductRequest request) {
        return RestAssured.given().log().all()
                .contentType(MediaType.APPLICATION_JSON_VALUE)
                .body(request)
                .when()
                .post("/products")
                .then()
                .log().all().extract();
    }

    public static ExtractableResponse<Response> 상품조회요청(Long productId) {
        return RestAssured.given().log().all()
                .when()
                .get("/products/{productId}", productId)
                .then()
                .log().all().extract();
    }

    public static UpdateProductRequest 상품수정요청_생성() {
        final String name = "초코에몽L";
        final int price = 2000;
        final DiscountPolicy discountPolicy = DiscountPolicy.NONE;
        return new UpdateProductRequest(name, price, discountPolicy);
    }

    public static ExtractableResponse<Response> 상품수정요청(Long productId, UpdateProductRequest request) {
        return RestAssured.given().log().all()
                .contentType(MediaType.APPLICATION_JSON_VALUE)
                .body(request)
                .when()
                .put("/products/{productId}", productId)
                .then()
                .log().all().extract();
    }
}
