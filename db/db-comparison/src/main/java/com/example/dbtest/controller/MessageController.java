package com.example.dbtest.controller;

import com.example.dbtest.dto.MessageDto;
import com.example.dbtest.service.MessageServiceImplMongo;
import com.example.dbtest.service.MessageServiceImplMysql;
import com.example.dbtest.service.MessageServiceImplRedis;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RestController
@RequiredArgsConstructor
public class MessageController {

    private final MessageServiceImplMysql mysqlService;
    private final MessageServiceImplMongo mongoService;
    private final MessageServiceImplRedis redisService;

    private static final int MYSQL = 1;
    private static final int MONGO = 2;
    private static final int REDIS = 3;

    @GetMapping
    public ResponseEntity<Void> save(
            @RequestParam("type") int type,
            @RequestParam("id") Long id,
            @RequestParam("content") String content
    ) {
        log.info("type: {}, id: {}, content: {}", type, id, content);
        MessageDto message = new MessageDto(type, id, content);

        if (type == MYSQL)
            mysqlService.save(message);

        if (type == MONGO)
            mongoService.save(message);

        if (type == REDIS)
            redisService.save(message);

        return ResponseEntity.ok().build();
    }

    @GetMapping("/{id}")
    public ResponseEntity<Void> findById(
            @PathVariable("id") String id,
            @RequestParam("type") int type
    ) {
        log.info("type: {}, id: {}", type, id);

        if (type == MYSQL)
            mysqlService.findById(Long.parseLong(id));

        if (type == MONGO)
            mongoService.findById(Long.parseLong(id));

        if (type == REDIS)
            redisService.findById(Long.parseLong(id));

        return ResponseEntity.ok().build();
    }
}
