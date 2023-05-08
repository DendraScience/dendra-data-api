import asyncio
import concurrent.futures
import logging
import os

# import pprint

import grpc
from grpc_reflection.v1alpha import reflection
from google.protobuf import empty_pb2

from v3 import converter_pb2
from v3 import converter_pb2_grpc
from v3 import metadata_pb2_grpc

# coroutines to be invoked when the event loop is shutting down
_cleanup_coroutines = []

uoms = None


class ConverterService(converter_pb2_grpc.ConverterServiceServicer):
    async def ConvertMany(
        self,
        request: converter_pb2.ConvertManyRequest,
        context: grpc.aio.ServicerContext,
    ) -> converter_pb2.ConvertManyResponse:
        logging.info("convert many request received")

        loop = asyncio.get_running_loop()

        def convert(datapoints):
            for datapoint in datapoints:
                # DEBUG: same value conversion
                datapoint.uv = datapoint.v
            return datapoints

        # TODO: init pint using from/to_unit, pass to convert
        # TODO: configure pool
        with concurrent.futures.ThreadPoolExecutor() as executor:
            datapoints = await loop.run_in_executor(
                executor, convert, request.datapoints
            )

        return converter_pb2.ConvertManyResponse(datapoints=datapoints)


# TODO: schedule this to occur periodically
async def fetch_uoms() -> None:
    logging.info("fetching uoms...")
    addr = os.environ.get("METADATA_SERVICE")

    # TODO: add timeout
    async with grpc.aio.insecure_channel(addr) as channel:
        stub = metadata_pb2_grpc.MetadataServiceStub(channel)
        response = await stub.ListUoms(empty_pb2.Empty())

    uoms = response.uoms
    logging.info("cached (%d) uoms", len(uoms))
    # TODO: init index/map for lookup?


async def serve() -> None:
    server = grpc.aio.server()
    converter_pb2_grpc.add_ConverterServiceServicer_to_server(
        ConverterService(), server
    )
    SERVICE_NAMES = (
        converter_pb2.DESCRIPTOR.services_by_name["ConverterService"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    listen_addr = "[::]:" + os.environ.get("PORT", "50051")
    server.add_insecure_port(listen_addr)
    await server.start()
    logging.info("server listening at %s", listen_addr)

    async def server_graceful_shutdown():
        logging.info("stopping server...")
        # shuts down the server with 5 seconds of grace period
        await server.stop(5)

    _cleanup_coroutines.append(server_graceful_shutdown())
    await server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(fetch_uoms())
        loop.run_until_complete(serve())
    except KeyboardInterrupt as e:
        print("caught keyboard interrupt")
    finally:
        loop.run_until_complete(*_cleanup_coroutines)
        loop.close()
    logging.info("bye!")
