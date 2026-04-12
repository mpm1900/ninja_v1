import { useActiveActor } from '#/hooks/use-active-actor'
import { useGetTargets } from '#/hooks/use-get-targets'
import type { Actor } from '#/lib/game/actor'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { ActorStat } from './actor-stat'
import { NatureBadge } from './nature-badge'
import { HoverCard, HoverCardContent, HoverCardTrigger } from './ui/hover-card'
import { NULL_CONTEXT } from '#/lib/game/context'
import { Button } from './ui/button'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { clientsStore } from '#/lib/stores/clients'
import { sendContextMessage } from '#/lib/stores/socket'
import { setActionID } from '#/lib/stores/battle-context'

function SwitchButton({ actor }: { actor: Actor }) {
  const active = useActiveActor()
  const client = useStore(clientsStore, (c) => c.me!)
  const game = useStore(gameStore, (g) => g)
  const idle = game.status === 'idle'
  const switch_action = active?.actions.find((a) => a.config.name === 'Switch')
  const context = {
    ...NULL_CONTEXT,
    action_ID: switch_action?.ID ?? null,
    parent_actor_ID: active?.ID ?? null,
    source_actor_ID: active?.ID ?? null,
    source_player_ID: active?.player_ID ?? null,
    target_actor_IDs: [],
  }

  const { context: t_context } = useGetTargets(context)
  if (!idle || !t_context?.target_actor_IDs?.includes(actor.ID)) {
    return null
  }

  return (
    <Button
      onClick={() => {
        sendContextMessage({
          type: 'push-action',
          client_ID: client.ID,
          context: {
            ...context,
            target_actor_IDs: [actor.ID],
          },
        })

        setActionID(context.source_actor_ID!, context.action_ID!, game)
      }}
    >
      Switch {actor.name}
    </Button>
  )
}

function ActorTooltip({
  actor,
  disabled = false,
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  actor: Actor
  disabled?: boolean
}) {
  return (
    <HoverCard openDelay={100} closeDelay={100}>
      <HoverCardTrigger {...props} />
      {!disabled && (
        <HoverCardContent sideOffset={8} collisionPadding={8}>
          <div className="flex justify-between">
            <div>{actor.name}</div>
            <div>
              {(Object.keys(actor.natures) as Array<NatureSet>)
                .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                .map((nature) => (
                  <NatureBadge key={nature} nature={nature} />
                ))}
            </div>
          </div>
          <table className="[&_td]:px-2">
            <tbody>
              <tr>
                <td>HP</td>
                <td>
                  <ActorStat actor={actor} showBase={false} stat={'hp'} />
                </td>
              </tr>
              <tr>
                <td>Stamina</td>
                <td>
                  <ActorStat actor={actor} showBase={false} stat={'stamina'} />
                </td>
              </tr>
              <tr>
                <td>Attack</td>
                <td>
                  <ActorStat actor={actor} showBase={false} stat={'attack'} />
                </td>
              </tr>
              <tr>
                <td>Defense</td>
                <td>
                  <ActorStat actor={actor} showBase={false} stat={'defense'} />
                </td>
              </tr>
              <tr>
                <td>C.Attack</td>
                <td>
                  <ActorStat
                    actor={actor}
                    showBase={false}
                    stat={'chakra_attack'}
                  />
                </td>
              </tr>
              <tr>
                <td>C.Defense</td>
                <td>
                  <ActorStat
                    actor={actor}
                    showBase={false}
                    stat={'chakra_defense'}
                  />
                </td>
              </tr>
              <tr>
                <td>Speed</td>
                <td>
                  <ActorStat actor={actor} showBase={false} stat={'speed'} />
                </td>
              </tr>
            </tbody>
          </table>
          <SwitchButton actor={actor} />
        </HoverCardContent>
      )}
    </HoverCard>
  )
}

export { ActorTooltip }
