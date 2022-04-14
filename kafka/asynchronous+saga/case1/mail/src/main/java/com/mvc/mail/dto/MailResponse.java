package com.mvc.mail.dto;

import com.mvc.mail.domain.MailLog;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class MailResponse {
    private Boolean state;

    public static MailResponse fromEntity(MailLog mailLog) {
        return new MailResponse(mailLog.getStatus());
    }
}
