package com.example.orderservice.order;

import com.example.orderservice.product.ProductService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static com.example.orderservice.product.ProductSteps.상품등록요청_생성;
import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
public class OrderServiceTest {

    @Autowired
    private ProductService productService;

    @Autowired
    private OrderService orderService;

    @Autowired
    private OrderRepository orderRepository;

    @Test
    void 주문생성() {
        // given
        productService.addProduct(상품등록요청_생성());
        Long productId = 1L;

        // when
        CreateOrderRequest request = new CreateOrderRequest(productId, 3);
        orderService.createOrder(request);

        // then
        Order savedOrder = orderRepository.findById(productId).get();
        assertThat(savedOrder.getProduct().getId()).isEqualTo(productId);
        assertThat(savedOrder.getQuantity()).isEqualTo(3);
    }
}
