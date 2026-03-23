import z from 'zod'

const ContextSchema = z.object({
  source_player_ID: z.string().nullable(),

  parent_actor_ID: z.string().nullable(),
  source_actor_ID: z.string().nullable(),

  target_actor_IDs: z.array(z.string()),
  target_position_IDs: z.array(z.string()),
})

type Context = z.output<typeof ContextSchema>

function contextToString(c: Context): string {
  return `${c.parent_actor_ID}.${c.source_actor_ID}.${c.source_player_ID}.${c.target_actor_IDs.join('+')}.${c.target_position_IDs.join('+')}`
}

export { ContextSchema, contextToString }
export type { Context }
