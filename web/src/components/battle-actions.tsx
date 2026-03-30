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
import type { Actor } from '#/lib/game/actor'
import { ActionCard } from './action-card'
import { battleContext, setContextAction } from '#/lib/stores/battle-context'

function BattleActions({ actor }: { actor: Actor }) {
  const game = useStore(gameStore, (g) => g)
  const context = useStore(battleContext, (c) => c)
  const action = actor.actions.find((a) => a.ID === context.action_ID)

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
                onContextChange={(c) => battleContext.setState((_) => c)}
              />
            </div>
          )}
        </CardContent>
      </Card>
      <div className="flex gap-2 absolute -bottom-20">
        {actor.actions.map((a) => (
          <ActionCard
            key={a.ID}
            action={a}
            onClick={() => setContextAction(a.ID)}
          />
        ))}
      </div>
    </div>
  )
}

export { BattleActions }
