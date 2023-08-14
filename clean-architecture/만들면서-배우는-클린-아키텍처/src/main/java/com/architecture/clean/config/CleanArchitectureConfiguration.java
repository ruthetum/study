package com.architecture.clean.config;

import com.architecture.clean.account.application.service.MoneyTransferProperties;
import com.architecture.clean.account.domain.Money;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@EnableConfigurationProperties(CleanArchitectureConfigurationProperties.class)
public class CleanArchitectureConfiguration {

    /**
     * Adds a use-case-specific {@link MoneyTransferProperties} object to the application context. The properties
     * are read from the Spring-Boot-specific {@link CleanArchitectureConfigurationProperties} object.
     */
    @Bean
    public MoneyTransferProperties moneyTransferProperties(CleanArchitectureConfigurationProperties cleanArchitectureConfigurationProperties){
        return new MoneyTransferProperties(Money.of(cleanArchitectureConfigurationProperties.getTransferThreshold()));
    }

}