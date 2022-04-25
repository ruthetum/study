package com.mvc.mail.domain;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "mail_log")
@Getter
@Setter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class MailLog {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "mail_log_id")
    private Long id;

    private Long orderId;

    private Boolean status;

    private LocalDateTime sendDt;

    public static MailLog createMailLog(Long orderId) {
        MailLog log = new MailLog();
        log.setOrderId(orderId);
        log.setStatus(true);
        log.setSendDt(LocalDateTime.now());
        return log;
    }

    public void failed() { this.status = false; }
}
