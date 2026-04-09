import { Card, CardContent } from './ui/card'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { ActionControl } from './action-control'
import type { Actor } from '#/lib/game/actor'
import { ActionsTable } from './actions-table'
import { useGameContext } from '#/hooks/use-game-context'
import { ActorStats } from './actor-stats'
import { FocusSelect } from './focus-select'
import { sendContextMessage } from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { NULL_CONTEXT } from '#/lib/game/context'
import { useQuery } from '@tanstack/react-query'
import { actionsQuery } from '#/lib/queries/actions'
import { AbilitySelect } from './ability-select'
import { ItemSelect } from './item-select'

function ActorControl({ actor, enabled }: { actor: Actor; enabled: boolean }) {
  const game = useStore(gameStore, (g) => g)
  const client = useStore(clientsStore, (c) => c.me!)
  const player = game.players.find((p) => p.ID == actor.player_ID)
  const actions_query = useQuery({
    ...actionsQuery,
    select: (all) => all.filter((a) => actor.action_IDs.includes(a.ID)),
  })
  const { context, onContextChange } = useGameContext(actor, undefined, [game])

  return (
    <Card className="flex flex-row rounded-t-none border-t-0 mx-2 mb-2 py-2 gap-0">
      <CardContent className="grid grid-cols-3 w-full">
        <div className="space-y-4">
          <div className="h-16 w-16 overflow-hidden">
            <img
              src={actor.sprite_url}
              className="h-full w-full object-cover"
              width={64}
              height={64}
            />
          </div>
          <AbilitySelect actor={actor} />
          <ItemSelect actor={actor} />
        </div>
        <ActorStats actor={actor} />
        <ActionsTable
          data={actions_query.data ?? []}
          enabled={enabled && !!player}
          rowSelection={Object.fromEntries(
            actor.actions.map((a) => [a.ID, true])
          )}
          onRowSelectionChange={(rowSelection) => {
            sendContextMessage({
              type: 'update-actor',
              client_ID: client.ID,
              context: {
                ...NULL_CONTEXT,
                source_actor_ID: actor.ID,
              },
              actor_config: {
                action_IDs: Object.keys(rowSelection),
              },
            })
          }}
          subRow={() => (
            <ActionControl
              action={actor.actions.find((a) => a.ID === context.action_ID)}
              enabled={enabled && !!player}
              context={context}
              onContextChange={onContextChange}
            />
          )}
        />
      </CardContent>
    </Card>
  )
}

export { ActorControl }
