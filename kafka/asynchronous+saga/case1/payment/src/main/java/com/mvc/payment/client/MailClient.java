package com.mvc.payment.client;

import com.mvc.payment.dto.MailRequest;
import com.mvc.payment.dto.MailResponse;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "MailFeign",
        url = "http://localhost:9002",
        fallback = MailClientFallback.class)
public interface MailClient {

    @PostMapping("/mail")
    MailResponse send(@RequestBody MailRequest request);
}
