import grpc
import logging
import torch

from concurrent.futures import ThreadPoolExecutor

from grpc_pb2 import (
    BlurResponse,
)  # fix import after creating grpc server for python using the .proto file
from grpc_pb2_grpc import (
    ImageServiceServicer,
    add_ImageServiceServicer_to_server,
)  # fix import after creating grpc server for python using the .proto file

from infer import infer

import time


class ImageServiceServer(ImageServiceServicer):
    def __init__(self):
        logging.info("Initializing ImageServiceServer")
        device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        self.model = torch.jit.load("./model/model_name.ts").to(device)
        logging.info("successfully loaded the model")

    def BlurBackground(self, request, context):
        logging.info("got request with blur: %r", request.blur_amount)
        start = time.time()
        encoded_image = infer(request, self.model)  # TODO: add, not currently included
        output_bytes = encoded_image.tobytes()
        end = time.time()
        print("elapsed time: %f" % (end - start))

        return BlurResponse(image=output_bytes)


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s - %(levelname)s - %(message)s",
    )
    server = grpc.server(ThreadPoolExecutor())
    add_ImageServiceServicer_to_server(ImageServiceServer(), server)

    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()

    logging.info("server ready on port %r", port)

    server.wait_for_termination()
