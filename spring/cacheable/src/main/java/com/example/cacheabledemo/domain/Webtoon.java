package com.example.cacheabledemo.domain;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.io.Serializable;
import java.time.DayOfWeek;

@Entity
@Table(name = "webtoon")
@Getter
@Setter
@NoArgsConstructor(access = AccessLevel.PROTECTED)
public class Webtoon implements Serializable {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "webtoon_id")
    private Long id;

    @Column(name = "webtoon_no", unique = true)
    private String webtoonNo;

    private String title;

    private int saved;

    private DayOfWeek dayOfWeek;

    public static Webtoon createWebtoon(String webtoonNo, String title, DayOfWeek dayOfWeek) {
        Webtoon webtoon = new Webtoon();
        webtoon.setWebtoonNo(webtoonNo);
        webtoon.setTitle(title);
        webtoon.setSaved(0);
        webtoon.setDayOfWeek(dayOfWeek);
        return webtoon;
    }
}
