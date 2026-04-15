import type { Actor } from '#/lib/game/actor'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import { HealthBar } from './health-bar'
import { NatureBadge } from './nature-badge'
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

const frameVariants = cva('', {
  variants: {
    variant: {
      default: 'bg-zinc-900 text-zinc-200',
      selected: 'bg-zinc-200 text-zinc-900 text-shadow-none!',
      targeted: 'bg-red-900 text-zinc-200',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})

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
}: React.ComponentProps<'div'> & {
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
        className={cn(
          is_player && 'cursor-pointer',
          'group/item flex flex-wrap items-center rounded-md text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50 p-2 w-92'
        )}
        {...rest}
      >
        {actor && (
          <div
            className={cn(
              'relative p-2 -mb-2 pr-3 -mr-1 rounded-lg rounded-tr-none rounded-l-2xl',
              frameVariants({
                variant: targeted
                  ? 'targeted'
                  : selected || source
                    ? 'selected'
                    : 'default',
              })
            )}
          >
            <ActorThumbnail actor={actor} size={50} imgClass="rounded-bl-xl" />
            <ActorStatus actor={actor} />
          </div>
        )}
        <div className="flex flex-1 flex-col relative gap-0">
          {actor && (
            <div className="flex justify-between items-end gap-4">
              <div
                className={cn(
                  'pl-2 pr-4 pb-2 pt-1.5 -mb-2 rounded-md rounded-tl-none shadow-[4px_1px_3px_rgba(0,0,0,0.2)]',
                  frameVariants({
                    variant: targeted
                      ? 'targeted'
                      : selected || source
                        ? 'selected'
                        : 'default',
                  })
                )}
              >
                <span className={cn('font-semibold text-lg')}>
                  {actor.name}
                </span>
              </div>
              {!action_tx || !is_player || status === 'running' ? (
                <div className="flex gap-px pb-1">
                  {(Object.keys(actor.natures) as Array<NatureSet>)
                    .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                    .map((nature) => (
                      <NatureBadge key={nature} nature={nature} />
                    ))}
                </div>
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
        </div>
      </div>
    </div>
  )
}

export { ActorCard }
