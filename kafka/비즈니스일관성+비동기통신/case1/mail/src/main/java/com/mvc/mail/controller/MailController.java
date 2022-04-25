package com.mvc.mail.controller;

import com.mvc.mail.dto.MailRequest;
import com.mvc.mail.dto.MailResponse;
import com.mvc.mail.service.MailService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
public class MailController {

    private final MailService mailService;

    @PostMapping("/mail")
    public MailResponse sendMail(@RequestBody MailRequest request) {
        return mailService.sendMail(request);
    }

}
