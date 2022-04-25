package com.mvc.orderkafka.domain;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "orders")
@Getter
@Setter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class Order {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "order_id")
    private Long id;

    private Long userId;

    @OneToMany(mappedBy = "order", cascade = CascadeType.ALL)
    private List<OrderItem> orderItems = new ArrayList<>();

    @Enumerated(EnumType.STRING)
    private OrderStatus status;

    public void addOrderItem(OrderItem orderItem) {
        orderItems.add(orderItem);
        orderItem.setOrder(this);
    }

    public static Order createOrder(Long userId, OrderItem... orderItems) {
        Order order = new Order();
        order.setUserId(userId);
        for (OrderItem orderItem: orderItems) {
            order.addOrderItem(orderItem);
        }
        order.setStatus(OrderStatus.WAIT);
        return order;
    }

    public int getTotalPrice() {
        return orderItems.stream()
                .mapToInt(OrderItem::getTotalPrice)
                .sum();
    }

    public void payment() {
        this.status = OrderStatus.PAYMENT;
    }

    public void complete() {
        this.status = OrderStatus.COMPLETED;
    }

    public void failed() { this.status = OrderStatus.FAILED; }
}