package com.mvc.mail.service;

import com.mvc.mail.domain.MailLog;
import com.mvc.mail.dto.MailRequest;
import com.mvc.mail.dto.MailResponse;
import com.mvc.mail.dto.SimpleMailRequest;
import com.mvc.mail.repository.MailLogRepository;
import com.mvc.mail.utils.MailUtil;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Slf4j
@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class MailService {

    private final MailLogRepository mailLogRepository;
    private final MailUtil mailUtil;

    private static final String TO = "khd0117@naver.com";

    @Transactional
    public MailResponse sendMail(MailRequest request) {

        MailLog mailLog = MailLog.createMailLog(request.getOrderId());

        // 메일 발송
        try {
            SimpleMailRequest simpleMailRequest = new SimpleMailRequest(
                    TO,
                    "[주문 번호 : " + request.getOrderId() + "] 결제 완료",
                    request.getPrice() + "원 결제 완료되었습니다.");

            mailUtil.send(simpleMailRequest);
        } catch (Exception e) {
            log.error("send mail error");
            mailLog.failed();
        }

        // 메일 발송 내역 저장
        mailLog = mailLogRepository.save(mailLog);

        return MailResponse.fromEntity(mailLog);
    }
}
