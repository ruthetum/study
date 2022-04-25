package com.mvc.payment.client;

import com.mvc.payment.dto.MailRequest;
import com.mvc.payment.dto.MailResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.openfeign.FallbackFactory;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class MailClientFallbackFactory implements FallbackFactory<MailClient> {
    @Override
    public MailClient create(Throwable cause) {
        return new MailClient() {
            @Override
            public MailResponse send(MailRequest request) {
                log.error(cause.toString());
                return null;
            }
        };
    }
}
