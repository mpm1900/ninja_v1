import type { Actor } from '#/lib/game/actor'
import type { Game } from '#/lib/game/game'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import { HealthBar } from './health-bar'
import { NatureBadge } from './nature-badge'
import { Item, ItemActions, ItemContent, ItemTitle } from './ui/item'
import { Badge } from './ui/badge'
import { X } from 'lucide-react'
import { sendContextMessage } from '#/lib/stores/socket'
import { NULL_CONTEXT } from '#/lib/game/context'
import { ActorThumbnail } from './actor-thumbnail'
import { StageBadge } from './stage-badge'

const actorVariants = cva(
  cn(
    'group/item flex flex-wrap items-center rounded-md border text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50',
    'p-2 w-86 border-transparent',
    'border-black border-2'
  ),
  {
    variants: {
      player: {
        player: 'bg-gray-800',
        enemy: '',
      },
      selected: {
        selected: 'bg-gray-700! border-gray-400',
        // source: 'scale-105 bg-blue-900! border-blue-300/40',
        targeted: 'bg-red-900! border-red-300/40',
        source: 'bg-yellow-900! border-yellow-300/40',
      },
    },
    defaultVariants: {},
  }
)

function ActorCard({
  actor,
  client_ID,
  game,
  selected,
  source,
  targeted,
  className,
  ...props
}: React.ComponentProps<typeof Item> & {
  actor: Actor | undefined
  client_ID?: string
  game: Game
  selected?: boolean
  targeted?: 'targeted'
  source?: 'source'
}) {
  const modifiers = (game.modifiers ?? [])
    .map((m) => m.mutation)
    .concat(actor?.innate_modifiers ?? [])

  const is_player = actor?.player_ID === client_ID
  const action_tx = game.actions.find(
    (t) => t.context.source_actor_ID === actor?.ID
  )
  const has_queued_action = game.queued_actions[actor?.ID ?? '']

  return (
    <div className={cn('flex flex-col', className)}>
      <div className="flex flex-wrap gap-3">
        {Object.entries(actor?.applied_modifiers ?? {}).map(([ID, count]) => (
          <span key={ID}>
            {modifiers.find((m) => m.group_ID === ID)?.name}
            {count > 1 ? ` (${count})` : null}
          </span>
        ))}
      </div>
      <div
        className={actorVariants({
          className: cn(is_player && 'cursor-pointer', 'gap-2'),
          player: is_player ? 'player' : 'enemy',
          selected: selected ? 'selected' : source ? source : targeted,
        })}
        {...props}
      >
        {actor && (
          <div className="relative">
            <ActorThumbnail actor={actor} size={50} />
            <div className="absolute font-bold px-1 h-4 leading-5 rounded whitespace-nowrap -bottom-1 z-10 bg-mist-300 text-background text-lg nanum-brush-script-regular">
              LV {actor.level}
            </div>
          </div>
        )}
        <ItemContent className="relative">
          {actor && (
            <div className="flex justify-between items-end gap-4">
              <ItemTitle>
                <span
                  className={cn({
                    'text-blue-300': actor.player_ID === client_ID,
                    'text-red-400': actor.player_ID !== client_ID,
                    'text-foregroud': selected,
                  })}
                >
                  {actor.name}
                </span>
              </ItemTitle>
              {!action_tx || !is_player || game.status === 'running' ? (
                <ItemActions className="gap-0">
                  {(Object.keys(actor.natures) as Array<NatureSet>)
                    .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                    .map((nature) => (
                      <NatureBadge key={nature} nature={nature} />
                    ))}
                </ItemActions>
              ) : (
                <Badge
                  onClick={() => {
                    if (has_queued_action) return

                    sendContextMessage({
                      type: 'remove-action',
                      client_ID: client_ID!,
                      context: {
                        ...NULL_CONTEXT,
                        action_ID: action_tx.ID,
                      },
                    })
                  }}
                >
                  {action_tx.mutation.config.name} {!has_queued_action && <X />}
                </Badge>
              )}
            </div>
          )}
          <div className="space-y-2">
            {actor && <HealthBar actor={actor} selected={selected} />}
          </div>
          {actor && (
            <div className="absolute -bottom-1.5 flex gap-1 px-2">
              {Object.entries(actor.staged_stats).map(([key, stage]) => (
                <StageBadge stage={stage as any} stat={key as any} />
              ))}
            </div>
          )}
        </ItemContent>
      </div>
    </div>
  )
}

export { ActorCard }
