package com.mvc.orderkafka.service;

import com.mvc.orderkafka.domain.Item;
import com.mvc.orderkafka.domain.Order;
import com.mvc.orderkafka.domain.OrderItem;
import com.mvc.orderkafka.dto.OrderRequest;
import com.mvc.orderkafka.dto.OrderResponse;
import com.mvc.orderkafka.repository.ItemRepository;
import com.mvc.orderkafka.repository.OrderRepository;
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

    private static final Long DEFAULT_USER_ID = 1L;

    @Transactional
    public Order order(OrderRequest request) {

        // 주문 저장
        Item item = itemRepository.findById(request.getItemId()).get();
        OrderItem orderItem = OrderItem.createOrderItem(item, item.getPrice(), request.getQuantity());
        Order order = Order.createOrder(DEFAULT_USER_ID, orderItem);
        order = orderRepository.save(order);
        return order;
    }
}
