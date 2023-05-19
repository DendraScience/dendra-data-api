/**
 * WaterOneFlow Web Service app.
 *
 * @author J. Scott Smith
 * @license BSD-3-Clause
 * @module app
 */

// SEE: https://grpc.github.io/grpc/node/index.html
import grpc from '@grpc/grpc-js'
import { addReflection } from 'grpc-server-reflection'
import messages from '../../../../../release/node/v3/transformer_pb.js'
import services from '../../../../../release/node/v3/transformer_grpc_pb.js'
import { DynamicThreadPool } from 'poolifier'

class TransformerService {
  constructor(options) {
    Object.assign(this, options)
  }

  transformMany(call, callback) {
    // DEBUG
    // this.logger.info('transform many request received')

    // pass the initial data array to the worker (see structured clone algorithm)
    this.pool
      .execute(call.request.array)
      .then(data => {
        callback(null, new messages.TransformManyResponse(data))
      })
      .catch(err => {
        callback(err, null)
      })
  }
}

export default async logger => {
  const app = {}

  // App setup
  app.start = async () => {
    // TODO: make configurable
    const pool = new DynamicThreadPool(10, 100, './src/worker.js', {
      errorHandler: e => logger.error(e),
      onlineHandler: () => logger.info('worker thread is online')
    })
    app.pool = pool

    const port = '[::]:' + (process.env.PORT || '50051')
    const server = new grpc.Server()

    server.addService(
      services.TransformerServiceService,
      new TransformerService({ logger, pool })
    )
    addReflection(server, '../../../../release/node/descriptor_set.bin')

    await new Promise((resolve, reject) => {
      server.bindAsync(port, grpc.ServerCredentials.createInsecure(), err =>
        err ? reject(err) : resolve()
      )
    })

    server.start()
    app.server = server

    logger.info('server listening at %s', port)
  }

  app.stop = async () => {
    if (app.pool) app.pool.destroy()

    if (app.server) {
      logger.info('stopping server...')
      // shuts down the server with 5 seconds of grace period
      setTimeout(() => app.server.forceShutdown(), 5000)
      return new Promise(resolve => app.server.tryShutdown(resolve))
    }
  }

  return app
}
