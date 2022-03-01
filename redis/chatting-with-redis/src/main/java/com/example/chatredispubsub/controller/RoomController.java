package com.example.chatredispubsub.controller;

import com.example.chatredispubsub.service.RoomService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

@Controller
@Slf4j
@RequiredArgsConstructor
@RequestMapping("/chat")
public class RoomController {

    private final RoomService roomService;

    //채팅방 목록 조회
    @GetMapping("/rooms")
    public String showRooms(Model model){
        log.info("GET /chat/rooms");
        model.addAttribute("list", roomService.getAllRooms());
        return "rooms";
    }

    //채팅방 개설
    @PostMapping("/room")
    public String create(@RequestParam String name, RedirectAttributes rttr){
        log.info("POST /chat/rooms");
        String roomName = roomService.createRoom(name);
        rttr.addFlashAttribute("roomName", roomName);
        return "redirect:/chat/rooms";
    }

    //채팅방 조회
    @GetMapping("/room/{roomId}")
    public String getRoom(
            @PathVariable Long roomId,
            Model model
    ){
        log.info("GET /chat/rooms/{}", roomId);
        model.addAttribute("room", roomService.getRoomDetail(roomId));
        return "room";
    }
}
