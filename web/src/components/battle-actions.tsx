import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from './ui/card'
import { ActionControl } from './action-control'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { useEffect, useState } from 'react'
import type { Action } from '#/lib/game/action'
import type { Context } from '#/lib/game/context'
import type { Actor } from '#/lib/game/actor'
import { Button } from './ui/button'
import { ActionCard } from './action-card'

function BattleActions({ actor }: { actor: Actor }) {
  const game = useStore(gameStore, (g) => g)
  const client = useStore(clientsStore, (c) => c.me!)
  const [action, setAction] = useState<Action>()
  const [context, setContext] = useState<Context>({
    source_player_ID: client.ID,
    source_actor_ID: actor.ID,
    parent_actor_ID: actor.ID,
    target_actor_IDs: [],
    target_position_IDs: [],
  })

  useEffect(() => {
    setContext({
      ...context,
      source_actor_ID: actor.ID,
      parent_actor_ID: actor.ID,
    })
  }, [actor.ID])

  return (
    <div className="flex flex-col items-center gap-4">
      <Card className="w-xl">
        <CardHeader>
          <CardTitle>Select Targets</CardTitle>
          <CardDescription>For {action?.config.name}</CardDescription>
        </CardHeader>
        <CardContent>
          {action && (
            <div>
              <ActionControl
                action={action}
                enabled={
                  game.status === 'idle' &&
                  !!actor.position_ID &&
                  actor.action_cooldowns[action.ID] == undefined
                }
                context={context}
                onContextChange={setContext}
              />
            </div>
          )}
        </CardContent>
      </Card>
      <div className="flex gap-2 absolute -bottom-20">
        {actor.actions.map((a) => (
          <ActionCard key={a.ID} action={a} onClick={() => setAction(a)} />
        ))}
      </div>
    </div>
  )
}

export { BattleActions }
