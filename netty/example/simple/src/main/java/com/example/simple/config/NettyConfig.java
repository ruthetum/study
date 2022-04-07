package com.example.simple.config;

import com.example.simple.netty.NettyChannelInitializer;
import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.logging.LogLevel;
import io.netty.handler.logging.LoggingHandler;
import io.netty.handler.ssl.SslContext;
import io.netty.handler.ssl.SslContextBuilder;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.net.ssl.SSLException;
import java.io.File;
import java.io.FileNotFoundException;
import java.net.InetSocketAddress;

@Configuration
@RequiredArgsConstructor
public class NettyConfig {

    @Value("${server.host}")
    private String host;

    @Value("${server.port}")
    private int port;

    @Value("${server.netty.boss.thread.count}")
    private int bossCount;

    @Value("${server.netty.worker.thread.count}")
    private int workerCount;

    @Bean
    public ServerBootstrap serverBootstrap(NettyChannelInitializer nettyChannelInitializer) {
        ServerBootstrap b = new ServerBootstrap();
        b.group(bossGroup(), workerGroup())
                .channel(NioServerSocketChannel.class)
                .handler(new LoggingHandler(LogLevel.DEBUG))
                .childHandler(nettyChannelInitializer);
        return b;
    }

    @Bean(destroyMethod = "shutdownGracefully")
    public NioEventLoopGroup bossGroup() {
        return new NioEventLoopGroup(bossCount);
    }

    @Bean(destroyMethod = "shutdownGracefully")
    public NioEventLoopGroup workerGroup() {
        return new NioEventLoopGroup(workerCount);
    }

    @Bean
    public InetSocketAddress inetSocketAddress() {
        return new InetSocketAddress(host, port);
    }

//    @Bean
//    public SslContext sslContext() {
//        SslContext sslCtx = null;
//
//        try {
//            File certChainFile =
//                    ConfigReader.getInstance().getConfigFile(CoreConstantsName.SSL_PUBLIC_KEY);
//            File keyFile =
//                    ConfigReader.getInstance().getConfigFile(CoreConstantsName.SSL_PRIVATE_KEY);
//
//            sslCtx = SslContextBuilder.forServer(certChainFile, keyFile, null).build();
//
//        }
//        catch (SSLException | FileNotFoundException e) {
//            e.printStackTrace();
//        }
//    }
}
