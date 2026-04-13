import type { Actor } from '#/lib/game/actor'
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
import { ActorModifiers } from './actor-modifiers'
import { ActorStatus } from './actor-status'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'

const actorVariants = cva(
  cn(
    'group/item flex flex-wrap items-center rounded-md border text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50',
    'p-2 w-92'
  ),
  {
    variants: {
      player: {
        player: 'border-transparent', //'bg-gray-900 border-gray-700',
        enemy: 'border-transparent',
        summon: 'border-transparent', //'bg-stone-600! border-stone-400! shadow-xs',
      },
      selected: {
        selected:
          'bg-stone-500! border-stone-300! shadow-[0px_5px_10px_rgba(0,0,0,0.7)]',
        // source: 'scale-105 bg-blue-900! border-blue-300/40',
        targeted:
          'bg-red-900! border-red-300/40! shadow-[0px_5px_10px_rgba(0,0,0,0.7)]',
        source:
          'bg-yellow-900! border-yellow-300/40 shadow-[0px_5px_10px_rgba(0,0,0,0.7)]',
      },
    },
    defaultVariants: {},
  }
)

function parseSummon(summon: Actor['summon'], parent: Actor): Actor['summon'] {
  if (!summon) return summon
  if (summon.proxy) {
    return {
      ...summon,
      ...parent,
      ID: summon.ID,
      stats: {
        ...parent.stats,
        hp: summon.stats.hp,
      },
      damage: summon.damage,
      sprite_url: summon.sprite_url,
      applied_modifiers: {},
      summon: undefined,
    }
  }

  return summon
}

function ActorCard({
  actor,
  client_ID,
  selected,
  source,
  targeted,
  className,
  summon,
  summonClass,
  ...rest
}: React.ComponentProps<typeof Item> & {
  actor: Actor | undefined
  client_ID?: string
  selected?: boolean
  targeted?: boolean
  source?: boolean
  summon?: boolean
  summonClass?: string
}) {
  const status = useStore(gameStore, (s) => s.status)
  const actions = useStore(gameStore, (s) => s.actions)
  const modifiers = useStore(gameStore, (g) =>
    (g.modifiers ?? [])
      .map((m) => m.mutation)
      .concat(g.actors.filter((a) => a.ability).map((a) => a.ability!))
      .concat(g.actors.filter((a) => a.item).map((a) => a.item!))
  )
  const has_queued_action = useStore(
    gameStore,
    (s) => s.queued_actions[actor?.ID ?? '']
  )
  const action_tx = actions.find((t) => t.context.source_actor_ID === actor?.ID)
  const is_player = actor?.player_ID === client_ID

  return (
    <div
      className={cn(
        'relative',
        summon && 'pointer-events-none',
        'flex flex-col',
        className
      )}
    >
      {actor?.summon?.proxy && (
        <ActorCard
          summon
          actor={parseSummon(actor.summon, actor)}
          client_ID={client_ID}
          selected={selected}
          targeted={targeted}
          source={source}
          className={cn('absolute top-0 left-0 z-20', summonClass)}
        />
      )}
      {actor && <ActorModifiers actor={actor} modifiers={modifiers} />}
      <div
        className={actorVariants({
          className: cn(is_player && 'cursor-pointer', 'gap-2'),
          player: summon ? 'summon' : is_player ? 'player' : 'enemy',
          selected: selected
            ? 'selected'
            : source
              ? 'source'
              : targeted
                ? 'targeted'
                : undefined,
        })}
        {...rest}
      >
        {actor && (
          <div className="relative">
            <ActorThumbnail actor={actor} size={50} />
            <ActorStatus actor={actor} />
          </div>
        )}
        <ItemContent className="relative gap-0">
          {actor && (
            <div className="flex justify-between items-end gap-4">
              <ItemTitle className={cn('px-2 pb-2 -mb-2 rounded-t')}>
                <span
                  className={cn(
                    'font-semibold text-lg text-shadow-[1px_1px_0px_#000000]',
                    {
                      'text-foreground':
                        actor.player_ID === client_ID ||
                        selected ||
                        targeted ||
                        source,
                      'text-red-300': actor.player_ID !== client_ID,
                    }
                  )}
                >
                  {actor.name}
                </span>
              </ItemTitle>
              {!action_tx || !is_player || status === 'running' ? (
                <ItemActions className="gap-px pb-1">
                  {(Object.keys(actor.natures) as Array<NatureSet>)
                    .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                    .map((nature) => (
                      <NatureBadge key={nature} nature={nature} />
                    ))}
                </ItemActions>
              ) : (
                <Badge
                  onClick={(e) => {
                    e.preventDefault()
                    e.stopPropagation()
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
              {Object.entries(actor.staged_stats)
                .filter(([key]) => key !== 'evasion' && key !== 'accuracy')
                .map(([key, stage]) => (
                  <StageBadge
                    key={key}
                    stage={stage as any}
                    stat={key as any}
                  />
                ))}
            </div>
          )}
        </ItemContent>
      </div>
    </div>
  )
}

export { ActorCard }
