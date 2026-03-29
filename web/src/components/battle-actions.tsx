import { Card, CardContent, CardHeader, CardTitle } from './ui/card'
import { ActionControl } from './action-control'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { useState } from 'react'
import type { Action } from '#/lib/game/action'
import type { Context } from '#/lib/game/context'
import type { Actor } from '#/lib/game/actor'
import { Button } from './ui/button'

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
  return (
    <Card className="w-xl">
      <CardHeader>
        <CardTitle>Select an action</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="gap-2 grid grid-cols-3">
          {actor.actions.map((a) => (
            <Button
              variant={action?.ID === a.ID ? 'default' : 'secondary'}
              onClick={() => setAction(a)}
            >
              {a.config.name}
            </Button>
          ))}
        </div>
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
  )
}

export { BattleActions }
