package grpc

import (
	"context"
	"log"

	proto "/goproto" // will have to be fixed after building the grpc server

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BackgroundBlur(image []byte, blurAmount uint) []byte {
	// setup proto server
	addr := "localhost:9999"

	const TEN_MB = 10485760
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
		return image
	}

	defer conn.Close()

	client := proto.NewImageServiceClient(conn)

	response, err := client.BlurBackground(context.Background(), &proto.BlurRequest{
		Image:      image,
		BlurAmount: uint32(blurAmount),
	}, grpc.MaxSendMsgSizeCallOption{MaxSendMsgSize: TEN_MB}, grpc.MaxRecvMsgSizeCallOption{MaxRecvMsgSize: TEN_MB})

	if err != nil {
		log.Fatal(err)
		return image
	}

	return response.Image
}
