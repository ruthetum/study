package com.example.cacheabledemo.config;

import com.example.cacheabledemo.domain.Webtoon;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.PostConstruct;
import javax.persistence.EntityManager;
import java.time.DayOfWeek;

@Component
@RequiredArgsConstructor
public class Init {

    private final InitService initService;

    @PostConstruct
    public void init() {
        initService.dbInit();
    }

    @Component
    @Transactional
    @RequiredArgsConstructor
    static class InitService {

        private final EntityManager em;

        public void dbInit() {
            Webtoon sunWebtoon1 =  Webtoon.createWebtoon("suw1", "독립일기", DayOfWeek.SUNDAY);
            em.persist(sunWebtoon1);
            Webtoon sunWebtoon2 =  Webtoon.createWebtoon("suw2", "싸움독학", DayOfWeek.SUNDAY);
            em.persist(sunWebtoon2);
            Webtoon sunWebtoon3 =  Webtoon.createWebtoon("suw3", "입학용병", DayOfWeek.SUNDAY);
            em.persist(sunWebtoon3);

            Webtoon monWebtoon1 =  Webtoon.createWebtoon("mow1", "참교육", DayOfWeek.MONDAY);
            em.persist(monWebtoon1);
            Webtoon monWebtoon2 =  Webtoon.createWebtoon("mow2", "신의 탑", DayOfWeek.MONDAY);
            em.persist(monWebtoon2);
            Webtoon monWebtoon3 =  Webtoon.createWebtoon("mow3", "뷰티풀 군바리", DayOfWeek.MONDAY);
            em.persist(monWebtoon3);

            Webtoon tueWebtoon1 =  Webtoon.createWebtoon("tuw1", "김부장", DayOfWeek.TUESDAY);
            em.persist(tueWebtoon1);
            Webtoon tueWebtoon2 =  Webtoon.createWebtoon("tuw2", "여신강림", DayOfWeek.TUESDAY);
            em.persist(tueWebtoon2);
            Webtoon tueWebtoon3 =  Webtoon.createWebtoon("tuw3", "대학원 탈출일지", DayOfWeek.TUESDAY);
            em.persist(tueWebtoon3);

            Webtoon wenWebtoon1 =  Webtoon.createWebtoon("wew1", "화산귀환", DayOfWeek.WEDNESDAY);
            em.persist(wenWebtoon1);
            Webtoon wenWebtoon2 =  Webtoon.createWebtoon("wew2", "전지적 독자 시점", DayOfWeek.WEDNESDAY);
            em.persist(wenWebtoon2);
            Webtoon wenWebtoon3 =  Webtoon.createWebtoon("wew3", "헬퍼2", DayOfWeek.WEDNESDAY);
            em.persist(wenWebtoon3);

            Webtoon thrWebtoon1 =  Webtoon.createWebtoon("thw1", "연애혁명", DayOfWeek.THURSDAY);
            em.persist(thrWebtoon1);
            Webtoon thrWebtoon2 =  Webtoon.createWebtoon("thw2", "더 복서", DayOfWeek.THURSDAY);
            em.persist(thrWebtoon2);
            Webtoon thrWebtoon3 =  Webtoon.createWebtoon("thw3", "기기괴괴", DayOfWeek.THURSDAY);
            em.persist(thrWebtoon3);

            Webtoon friWebtoon1 =  Webtoon.createWebtoon("frw1", "외모지상주의", DayOfWeek.FRIDAY);
            em.persist(friWebtoon1);
            Webtoon friWebtoon2 =  Webtoon.createWebtoon("frw2", "대학원 탈출일지", DayOfWeek.FRIDAY);
            em.persist(friWebtoon2);
            Webtoon friWebtoon3 =  Webtoon.createWebtoon("frw3", "나 혼자 만렙", DayOfWeek.FRIDAY);
            em.persist(friWebtoon3);

            Webtoon satWebtoon1 =  Webtoon.createWebtoon("saw1", "프리드로우", DayOfWeek.SATURDAY);
            em.persist(satWebtoon1);
            Webtoon satWebtoon2 =  Webtoon.createWebtoon("saw2", "신림/남/22", DayOfWeek.SATURDAY);
            em.persist(satWebtoon2);
            Webtoon satWebtoon3 =  Webtoon.createWebtoon("saw3", "취사병 전설이 되다", DayOfWeek.SATURDAY);
            em.persist(satWebtoon3);
        }
    }
}
