import { createServerFn } from '@tanstack/react-start'
import { ContextSchema, type Context } from '../game/context'
import { z } from 'zod'
import { queryOptions } from '@tanstack/react-query'

const requestSchema = z.object({
  instanceID: z.string(),
  actionID: z.string(),
  context: ContextSchema,
})

const getActionTargets = createServerFn()
  .inputValidator(requestSchema)
  .handler(async ({ data }) => {
    const response = await fetch(
      `${process.env.API_URL}/api/${data.instanceID}/${data.actionID}/targets`,
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data.context),
      }
    )

    const response_data = await response.json()
    return response_data as Array<string>
  })

function actionTargetsQuery(
  instanceID: string,
  actionID: string | undefined,
  context: Context
) {
  return queryOptions({
    queryKey: ['action-targets', instanceID, actionID, context.source_actor_ID],
    queryFn: async () => {
      return await getActionTargets({
        data: {
          instanceID,
          actionID: actionID!,
          context,
        },
      })
    },
    enabled: !!actionID,
  })
}

export { getActionTargets, actionTargetsQuery }
