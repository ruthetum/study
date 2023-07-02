package com.example.orderservice.order;

import com.example.orderservice.product.Product;
import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Entity
@Table(name = "orders")
@NoArgsConstructor(access = AccessLevel.PROTECTED)
@Getter
public class Order {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @OneToOne
    private Product product;

    private int quantity;

    public static Order create(Product product, int quantity) {
        Order order = new Order();
        order.product = product;
        order.quantity = quantity;
        return order;
    }
}
