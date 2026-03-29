import { createServerFn } from '@tanstack/react-start'
import z from 'zod'
import { ContextSchema, contextToString, type Context } from '../game/context'
import { queryOptions } from '@tanstack/react-query'

const requestSchema = z.object({
  context: ContextSchema,
})

const getIsActionContextValid = createServerFn()
  .inputValidator(requestSchema)
  .handler(async ({ data }) => {
    const response = await fetch(
      `${process.env.API_URL}/api/validate`,
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data.context),
      }
    )

    const response_data = await response.json()
    return response_data as boolean
  })

function isActionContextValidQuery(
  context: Context
) {
  return queryOptions({
    queryKey: ['action-targets', contextToString(context)],
    queryFn: async () => {
      return await getIsActionContextValid({
        data: {
          context,
        },
      })
    },
    enabled: !!context.action_ID,
  })
}

export { isActionContextValidQuery }
