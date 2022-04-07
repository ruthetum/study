package com.example.telnetserver;

import com.example.telnetserver.config.TelnetServerConfig;
import com.example.telnetserver.telnet.TelnetServer;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.support.AbstractApplicationContext;

public class TelnetServerApplication {

    public static void main(String[] args) {

        AbstractApplicationContext springContext = null;

        try {
            springContext = new AnnotationConfigApplicationContext(TelnetServerConfig.class);
            springContext.registerShutdownHook();

            TelnetServer server = springContext.getBean(TelnetServer.class);
            server.start();
        } finally {
            springContext.close();
        }
    }
}
