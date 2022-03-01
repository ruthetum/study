package com.example.chatredispubsub.entity;

import lombok.Getter;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

@Entity
@Getter
public class Room {

    @Id
    @GeneratedValue
    @Column(name = "room_id")
    private Long id;

    @Column
    private String name;

    public void setName(String name) {
        this.name = name;
    }
}
