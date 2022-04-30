package com.mvc.orderkafka.repository;

import com.mvc.orderkafka.domain.Order;
import org.springframework.data.jpa.repository.JpaRepository;

public interface OrderRepository extends JpaRepository<Order, Long> {
}
