package com.example.chatredispubsub.controller;

import com.example.chatredispubsub.service.RedisPublisher;
import com.example.chatredispubsub.dto.MessageDto;
import com.example.chatredispubsub.repository.ChatRoomRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.stereotype.Controller;

@Slf4j
@Controller
@RequiredArgsConstructor
public class ChatController {

    // private final SimpMessagingTemplate template;
    private final RedisPublisher redisPublisher;
    private final ChatRoomRepository chatRoomRepository;

    @MessageMapping("/chat/enter")
    public void enter(MessageDto message){
        message.setMessage(message.getWriter() + "님이 채팅방에 참여하였습니다.");
        // template.convertAndSend("/sub/chat/room/" + message.getRoomId(), message);
        chatRoomRepository.enterChatRoom(message.getRoomId());
        redisPublisher.publish(chatRoomRepository.getTopic(message.getRoomId()), message);
    }

    @MessageMapping("/chat/message")
    public void message(MessageDto message){
        // template.convertAndSend("/sub/chat/room/" + message.getRoomId(), message);
        chatRoomRepository.enterChatRoom(message.getRoomId());
        redisPublisher.publish(chatRoomRepository.getTopic(message.getRoomId()), message);
    }
}