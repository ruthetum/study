package com.mvc.payment.client;

import com.mvc.payment.dto.MailRequest;
import com.mvc.payment.dto.MailResponse;

public class MailClientFallback implements MailClient {

    @Override
    public MailResponse send(MailRequest request) {
        return null;
    }
}
