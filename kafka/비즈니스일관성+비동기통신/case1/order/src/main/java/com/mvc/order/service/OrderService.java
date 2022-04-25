package com.mvc.order.service;

import com.mvc.order.client.PaymentClient;
import com.mvc.order.domain.Item;
import com.mvc.order.domain.Order;
import com.mvc.order.domain.OrderItem;
import com.mvc.order.dto.OrderRequest;
import com.mvc.order.dto.OrderResponse;
import com.mvc.order.dto.PaymentRequest;
import com.mvc.order.dto.PaymentResponse;
import com.mvc.order.repository.ItemRepository;
import com.mvc.order.repository.OrderRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class OrderService {

    private final OrderRepository orderRepository;
    private final ItemRepository itemRepository;

    private final PaymentClient paymentClient;

    private static final Long DEFAULT_USER_ID = 1L;

    @Transactional
    public OrderResponse order(OrderRequest request) {
        
        // 주문 저장
        Item item = itemRepository.findById(request.getItemId()).get();
        OrderItem orderItem = OrderItem.createOrderItem(item, item.getPrice(), request.getQuantity());
        Order order = Order.createOrder(DEFAULT_USER_ID, orderItem);
        order = orderRepository.save(order);

        // 결제 서비스 요청
        PaymentRequest req = PaymentRequest.fromEntity(order);
        PaymentResponse res = paymentClient.pay(req);

        if (res == null) {
            log.error("payment server error");
            order.failed();
            return new OrderResponse(order.getId(), false);
        }

        if (Boolean.FALSE.equals(res.getState()))
            order.failed();

        return new OrderResponse(order.getId(), res.getState());
    }
}
