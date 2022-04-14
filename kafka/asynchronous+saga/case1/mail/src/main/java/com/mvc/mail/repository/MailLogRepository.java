package com.mvc.mail.repository;

import com.mvc.mail.domain.MailLog;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MailLogRepository extends JpaRepository<MailLog, Long> {
}
