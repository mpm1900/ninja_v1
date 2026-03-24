import { useSuspenseQuery } from '@tanstack/react-query'
import { Card, CardContent } from './ui/card'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { useStore } from '@tanstack/react-store'
import { actionsQuery } from '#/lib/queries/actions'
import { clientsStore, type Client } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { useState } from 'react'
import type { Context } from '#/lib/game/context'
import { ActionControl } from './action-control'

function ActorControl({ actor_ID, enabled }: { actor_ID: string, enabled: boolean }) {
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
    <Card className="rounded-t-none border-t-0 mx-2 mb-2 py-2">
      <CardContent className='px-2'>
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
          enabled={enabled}
          context={context}
          onContextChange={setContext}
        />
      </CardContent>
    </Card>
  )
}

export { ActorControl }
