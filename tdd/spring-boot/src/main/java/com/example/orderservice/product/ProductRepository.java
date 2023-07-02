package com.example.orderservice.product;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.HashMap;
import java.util.Map;

@Repository
public
interface ProductRepository  extends JpaRepository<Product, Long> {
}
