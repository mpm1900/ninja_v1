import type { Actor } from '#/lib/game/actor'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { ActorStat } from './actor-stat'
import { NatureBadge } from './nature-badge'
import { HoverCard, HoverCardContent, HoverCardTrigger } from './ui/hover-card'

function ActorTooltip({
  actor,
  disabled = false,
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  actor: Actor
  disabled?: boolean
}) {
  return (
    <HoverCard openDelay={100} closeDelay={0}>
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
        </HoverCardContent>
      )}
    </HoverCard>
  )
}

export { ActorTooltip }
