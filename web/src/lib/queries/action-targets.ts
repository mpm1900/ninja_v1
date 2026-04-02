import { createServerFn } from '@tanstack/react-start'
import { ContextSchema, type Context } from '../game/context'
import { z } from 'zod'
import { queryOptions } from '@tanstack/react-query'

const requestSchema = z.object({
  instanceID: z.string(),
  context: ContextSchema,
})

const getActionTargets = createServerFn()
  .inputValidator(requestSchema)
  .handler(async ({ data }) => {
    const response = await fetch(
      `${process.env.API_URL}/api/${data.instanceID}/targets`,
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
  context: Context,
  deps: any[] = []
) {
  return queryOptions({
    queryKey: ['action-targets', instanceID, context.action_ID, context.source_actor_ID, ...deps],
    queryFn: async () => {
      return await getActionTargets({
        data: {
          instanceID,
          context,
        },
      })
    },
    enabled: !!context.action_ID,
  })
}

export { getActionTargets, actionTargetsQuery }
