package com.example.simple.netty.decoder;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler.Sharable;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.ByteToMessageDecoder;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class Decoder extends ByteToMessageDecoder {

    private static final int DATA_LENGTH = 2048;

    @Override
    protected void decode(ChannelHandlerContext ctx, ByteBuf in, List<Object> out) {
        if (in.readableBytes() < DATA_LENGTH)
            return;
        out.add(in.readBytes(DATA_LENGTH));
    }
}