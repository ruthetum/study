package com.example.chatredispubsub.dto;

import com.example.chatredispubsub.entity.Room;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.web.socket.WebSocketSession;

import java.util.HashSet;
import java.util.Set;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class RoomDtoResponse {
    private Long roomId;
    private String name;
    private Set<WebSocketSession> sessions = new HashSet<>();

    public static RoomDtoResponse fromEntity(Room room) {
        RoomDtoResponse roomDtoResponse = new RoomDtoResponse();
        roomDtoResponse.setRoomId(room.getId());
        roomDtoResponse.setName(room.getName());
        return roomDtoResponse;
    }
}
