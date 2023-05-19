import { ThreadWorker } from 'poolifier'
import { create, all } from 'mathjs'
import messages from '../../../../../release/node/v3/transformer_pb.js'

const math = create(
  { all },
  {
    matrix: 'Array',
    number: 'BigNumber',
    precision: 15 // Hardcoded!
  }
)

function evaluate(data) {
  // SEE: https://github.com/protocolbuffers/protobuf-javascript/blob/main/docs/index.md
  // SEE: https://protobuf.dev/programming-guides/proto3/
  const req = new messages.TransformManyRequest(data)
  const resp = new messages.TransformManyResponse()
  const expression = req.getTransform().getExpression()
  const datapoints = req.getDatapointsList()

  const code = math.compile(expression)

  for (let i = 0; i < datapoints.length; i++) {
    const datapoint = datapoints[i]
    const clone = datapoint.clone()
    const obj = datapoint.toObject()

    // prepare compatible eval sandbox
    obj.t = obj.t.seconds
    obj.lt = obj.lt.seconds
    obj.v = clone.hasV() ? obj.v : null
    obj.da = obj.daList
    obj.va = obj.vaList
    obj.gc = obj.gcList

    try {
      code.evaluate(obj)
    } catch (_) {}

    // transform v
    if (obj.v === null) clone.clearV()
    else clone.setV(math.number(obj.v)) // handle BigNumber

    //
    // NOTE: NOT SUPPORTED YET - transform d, da, va
    //

    resp.addDatapoints(clone)
  }

  return resp.array
}

export default new ThreadWorker(evaluate, {
  maxInactiveTime: 60000
})
