package com.architecture.clean.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Data
@ConfigurationProperties(prefix = "clean")
public class CleanArchitectureConfigurationProperties {

    private long transferThreshold = Long.MAX_VALUE;

}