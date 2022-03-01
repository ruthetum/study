package com.example.chatredispubsub.repository;

import com.example.chatredispubsub.service.RedisSubscriber;
import com.example.chatredispubsub.entity.ChatRoom;
import lombok.RequiredArgsConstructor;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.listener.ChannelTopic;
import org.springframework.data.redis.listener.RedisMessageListenerContainer;
import org.springframework.stereotype.Repository;

import javax.annotation.PostConstruct;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@Repository
@RequiredArgsConstructor
public class ChatRoomRepository {

    private final RedisMessageListenerContainer redisMessageListenerContainer;
    private final RedisSubscriber redisSubscriber;

    private static final String CHAT_ROOM = "CHAT_ROOM";
    private final RedisTemplate<String, Object> redisTemplate;
    private HashOperations<String, String, ChatRoom> opsHashChatRoom;
    private Map<String, ChannelTopic> topics;

    @PostConstruct
    private void init() {
        opsHashChatRoom = redisTemplate.opsForHash();
        topics = new HashMap<>();
    }

    public List<ChatRoom> findAllChatRoom() {
        return opsHashChatRoom.values(CHAT_ROOM);
    }

    public ChatRoom findChatRoomByRid(String id) {
        return opsHashChatRoom.get(CHAT_ROOM, id);
    }

    public ChatRoom createChatRoom(String name) {
        ChatRoom room = ChatRoom.create(name);
        opsHashChatRoom.put(CHAT_ROOM, room.getId(), room);
        return room;
    }

    public void enterChatRoom(String id) {
        ChannelTopic topic = topics.get(id);
        if (Objects.isNull(topic)) {
            topic = new ChannelTopic(id);
            redisMessageListenerContainer.addMessageListener(redisSubscriber, topic);
            topics.put(id, topic);
        }
    }

    public ChannelTopic getTopic(String id) {
        return topics.get(id);
    }
}
