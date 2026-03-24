import { useSuspenseQuery } from '@tanstack/react-query'
import { Card, CardContent, CardHeader } from './ui/card'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { useStore } from '@tanstack/react-store'
import { actionsQuery } from '#/lib/queries/actions'
import { clientsStore, type Client } from '#/lib/stores/clients'
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
import { Button } from './ui/button'

function ActorControl({
  actor_ID,
  enabled,
}: {
  actor_ID: string
  enabled: boolean
}) {
  const actions = useSuspenseQuery(actionsQuery)
  const client = useStore(clientsStore, (c) => c[0] as Client)
  const game = useStore(gameStore, (g) => g)
  const source = game.actors.find((a) => a.actor_ID === actor_ID)!
  const [activeActionID, setActiveActionID] = useState(actions.data[0]?.ID)
  const [context, setContext] = useState<Context>({
    source_player_ID: client.ID,
    source_actor_ID: source.ID,
    parent_actor_ID: source.ID,
    target_actor_IDs: [],
    target_position_IDs: [],
  })

  return (
    <Card className="grid grid-cols-2 rounded-t-none border-t-0 mx-2 mb-2 py-2 gap-0">
      <div>
        <CardHeader className="px-2 flex justify-between">
          <div className="w-full grid grid-cols-3">
            {actions.data.map((a) => (
              <Button
                key={a.ID}
                variant={a.ID === activeActionID ? 'default' : 'outline'}
                onClick={() => {
                  setActiveActionID(a.ID)
                }}
              >
                {a.config.name}
              </Button>
            ))}
          </div>
        </CardHeader>
        <CardContent className="px-2">
          <ActionControl
            action_ID={activeActionID}
            enabled={enabled}
            context={context}
            onContextChange={setContext}
          />
        </CardContent>
      </div>
      <div>
        <CardHeader className="px-2">
          <div className="flex items-center justify-end gap-2">
            <div>Player:</div>
            <Select
              disabled={!enabled}
              value={source.player_ID}
              onValueChange={(playerID) => {
                sendContextMessage({
                  type: 'set-actor-player',
                  client_ID: client.ID,
                  context: {
                    source_player_ID: playerID,
                    source_actor_ID: null,
                    parent_actor_ID: null,
                    target_actor_IDs: [source.ID],
                    target_position_IDs: [],
                  },
                })
              }}
            >
              <SelectTrigger className="">
                <SelectValue>
                  {game.players.includes(source.player_ID) ? (
                    source.player_ID
                  ) : (
                    <span className="text-red-300">{source.player_ID}</span>
                  )}
                </SelectValue>
              </SelectTrigger>
              <SelectContent>
                {game.players.map((item) => (
                  <SelectItem key={item} value={item}>
                    {item}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </CardHeader>
      </div>
    </Card>
  )
}

export { ActorControl }
