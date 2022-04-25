package com.mvc.mail.utils;

import com.mvc.mail.dto.SimpleMailRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class MailUtil {

    private final JavaMailSender mailSender;

    private static final String FROM_ADDRESS = "sample@gmail.com";

    public void send(SimpleMailRequest request) {
        SimpleMailMessage message = new SimpleMailMessage();
        message.setTo(request.getTo());
        message.setFrom(FROM_ADDRESS);
        message.setSubject(request.getTitle());
        message.setText(request.getContent());
        mailSender.send(message);
    }
}
