import { sendContextMessage } from '#/lib/stores/socket'
import { useSuspenseQuery } from '@tanstack/react-query'
import { Button } from './ui/button'
import { Card, CardContent } from './ui/card'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { useStore } from '@tanstack/react-store'
import { actionsQuery } from '#/lib/queries/actions'
import { modifiersQuery } from '#/lib/queries/modifiers'
import { clientsStore, type Client } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { useState } from 'react'
import type { Context } from '#/lib/game/context'
import { ActionControl } from './action-control'

function ActorControl({ actor_ID }: { actor_ID: string }) {
  const actions = useSuspenseQuery(actionsQuery)
  const modifiers = useSuspenseQuery(modifiersQuery)
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
    <Card className="rounded-t-none border-t-0 mx-2 mb-2">
      <CardContent>
        {modifiers.data.map((m) => (
          <Button
            key={m.ID}
            onClick={() => {
              sendContextMessage({
                type: 'add-modifier',
                clientID: client.ID,
                modifierID: m.ID,
                context,
              })
            }}
          >
            {m.name}
          </Button>
        ))}
        <hr className="my-4" />
        <Tabs value={activeActionID} onValueChange={setActiveActionID}>
          <TabsList>
            {actions.data.map((a) => (
              <TabsTrigger key={a.ID} value={a.ID}>
                {a.config.name}
              </TabsTrigger>
            ))}
          </TabsList>
        </Tabs>
        <ActionControl
          action_ID={activeActionID}
          context={context}
          onContextChange={setContext}
        />
      </CardContent>
    </Card>
  )
}

export { ActorControl }
