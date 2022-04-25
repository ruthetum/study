package com.mvc.orderkafka.repository;

import com.mvc.orderkafka.domain.Item;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ItemRepository extends JpaRepository<Item, Long> {
}
