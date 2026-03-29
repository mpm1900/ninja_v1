import { Card, CardContent } from './ui/card'
import { useStore } from '@tanstack/react-store'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { useState } from 'react'
import type { Context } from '#/lib/game/context'
import { ActionControl } from './action-control'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select'
import { sendContextMessage } from '#/lib/stores/socket'
import { PositionSelect } from './position-select'
import type { Actor } from '#/lib/game/actor'
import { ActionsTable } from './actions-table'

function ActorControl({ actor, enabled }: { actor: Actor; enabled: boolean }) {
  const client = useStore(clientsStore, (c) => c.me!)
  const game = useStore(gameStore, (g) => g)
  const player = game.players.find((p) => p.ID == actor.player_ID)
  const [activeActionID, setActiveActionID] = useState<string>()
  const [context, setContext] = useState<Context>({
    source_player_ID: client.ID,
    source_actor_ID: actor.ID,
    parent_actor_ID: actor.ID,
    target_actor_IDs: [],
    target_position_IDs: [],
  })


  return (
    <Card className="grid grid-cols-2 rounded-t-none border-t-0 mx-2 mb-2 py-2 gap-0">
      <div>
        <CardContent className="px-2 flex flex-col gap-2">
          <div className="flex gap-2">
            <div className="h-16 w-16 overflow-hidden">
              <img
                src={actor.sprite_url}
                className="h-full w-full object-cover"
                width={64}
                height={64}
              />
            </div>
            <div className="flex items-center gap-2">
              <Select
                disabled={!enabled}
                value={actor.player_ID}
                onValueChange={(playerID) => {
                  sendContextMessage({
                    type: 'set-actor-player',
                    client_ID: client.ID,
                    context: {
                      source_player_ID: playerID,
                      source_actor_ID: null,
                      parent_actor_ID: null,
                      target_actor_IDs: [actor.ID],
                      target_position_IDs: [],
                    },
                  })
                }}
              >
                <SelectTrigger>
                  <SelectValue>
                    {game.players.find((p) => p.ID == actor.player_ID)?.user.email ?? (
                      <span className="text-red-300">{actor.player_ID}</span>
                    )}
                  </SelectValue>
                </SelectTrigger>
                <SelectContent>
                  {game.players.map((item) => (
                    <SelectItem key={item.ID} value={item.ID}>
                      {item.user.username || item.user.email}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>
          <PositionSelect actor={actor} game={game} />
        </CardContent>
      </div>
      <div>
        <CardContent className="px-2">
          <ActionsTable
            cooldowns={actor.action_cooldowns}
            data={actor.actions}
            enabled={enabled && !!player && !!actor.position_ID}
            selected={activeActionID}
            onSelectedChange={setActiveActionID}
            subRow={({ row }) => (
              <ActionControl
                action={actor.actions.find((a) => a.ID === activeActionID)}
                enabled={
                  enabled &&
                  !!player &&
                  !!actor.position_ID &&
                  actor.action_cooldowns[row.original.ID] == undefined
                }
                context={context}
                onContextChange={setContext}
              />
            )}
          />
        </CardContent>
      </div>
    </Card>
  )
}

export { ActorControl }
