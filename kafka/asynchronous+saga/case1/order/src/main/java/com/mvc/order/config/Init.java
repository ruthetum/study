package com.mvc.order.config;

import com.mvc.order.domain.Item;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.PostConstruct;
import javax.persistence.EntityManager;

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
            Item itemA = Item.createItem("itemA", 10000);
            em.persist(itemA);

            Item itemB = Item.createItem("itemB", 8000);
            em.persist(itemB);

            Item itemC = Item.createItem("itemC", 6000);
            em.persist(itemC);
        }
    }
}
