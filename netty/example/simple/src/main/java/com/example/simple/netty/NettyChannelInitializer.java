package com.example.simple.netty;

import com.example.simple.netty.decoder.Decoder;
import com.example.simple.netty.handler.NettyHandler;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.ssl.SslContext;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class NettyChannelInitializer extends ChannelInitializer<SocketChannel> {

    // private final SslContext sslContext;
    private final NettyHandler handler;

    @Override
    protected void initChannel(SocketChannel ch) {
        ChannelPipeline pipeline = ch.pipeline();
//        if (sslContext != null)
//            pipeline.addLast(sslContext.newHandler(ch.alloc()));

        pipeline.addLast(new Decoder());
        pipeline.addLast(handler);
    }
}
