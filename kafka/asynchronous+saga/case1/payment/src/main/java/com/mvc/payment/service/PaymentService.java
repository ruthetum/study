package com.mvc.payment.service;

import com.mvc.payment.client.MailClient;
import com.mvc.payment.domain.Payment;
import com.mvc.payment.dto.MailRequest;
import com.mvc.payment.dto.MailResponse;
import com.mvc.payment.dto.PaymentRequest;
import com.mvc.payment.dto.PaymentResponse;
import com.mvc.payment.repository.PaymentRepository;
import com.mvc.payment.utils.KakaoPayUtil;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class PaymentService {

    private final PaymentRepository paymentRepository;
    private final MailClient mailClient;
    private final KakaoPayUtil kakaoPayUtil;

    @Transactional
    public PaymentResponse pay(PaymentRequest request) {

        // 결제
        try {
            kakaoPayUtil.pay();
        } catch (Exception e) {
            log.error("kakao pay error");
            throw new IllegalStateException("kakao pay error");
        }
        
        // 결제 성공 시 결제 내역 저장
        Payment payment = Payment.createPayment(request.getUserId(), request.getOrderId(), request.getPrice());
        payment = paymentRepository.save(payment);

        // 메일 전송 요청
        MailRequest req = MailRequest.fromPayment(payment);
        MailResponse res = mailClient.send(req);

        if (res.getState() == null) {
            log.error("mail server error");
            payment.failed();
            return new PaymentResponse(false);
        }

        if (Boolean.FALSE.equals(res.getState())) {
            log.error("mail send error");
            payment.failed();
        }

        return new PaymentResponse(res.getState());
    }
}
