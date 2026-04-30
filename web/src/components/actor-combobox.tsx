import type { ActorDef } from '#/lib/game/actor'
import { actorsQuery } from '#/lib/queries/actors'
import { useSuspenseQuery } from '@tanstack/react-query'
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
  ComboboxTrigger,
  ComboboxValue,
} from './ui/combobox'
import { ChevronsUpDown, Plus } from 'lucide-react'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { NatureBadge } from './nature-badge'
import { Button } from './ui/button'
import { cn } from '#/lib/utils'

function ActorCombobox({
  className,
  onClick,
  selected = [],
  active,
  value,
  onValueChange,
}: {
  className?: string
  onClick?: () => void
  selected?: string[]
  active?: string
  value?: string | null
  onValueChange?: (def: ActorDef) => void
}) {
  const actors = useSuspenseQuery(actorsQuery)
  const sortedActors = actors.data.sort((a, b) => a.name.localeCompare(b.name))
  const actor = sortedActors.find((actor) => actor.actor_ID === value) ?? null
  const is_active = !!actor && active === actor.actor_ID

  const handleValueChange = (actor: ActorDef | null) => {
    if (!actor?.actor_ID) return
    onValueChange?.(actor)
  }
  return (
    <Combobox<ActorDef>
      items={sortedActors}
      itemToStringValue={(actor) => actor.actor_ID}
      itemToStringLabel={(actor) => actor.name}
      value={actor}
      onValueChange={handleValueChange}
    >
      <ComboboxTrigger
        render={
          <Button
            variant="outline"
            className={cn(
              'relative justify-between font-normal h-16 px-2',
              {
                'bg-input!': is_active,
              },
              className
            )}
          >
            <div className="flex gap-4">
              {actor ? (
                <img
                  src={actor.sprite_url}
                  className={cn(
                    'size-15 p-0.5 -m-1.5 bg-stone-200/20 border border-stone-950 rounded cursor-pointer',
                    is_active && 'bg-stone-300'
                  )}
                  onPointerDown={(e) => {
                    e.preventDefault()
                    e.stopPropagation()
                  }}
                  onClick={(e) => {
                    e.preventDefault()
                    e.stopPropagation()
                    onClick?.()
                  }}
                />
              ) : (
                <Plus className="text-muted-foreground/60 size-8" />
              )}
              <div className="text-left space-y-1">
                <div
                  className={cn(
                    'font-semibold text-md',
                    !is_active && 'text-muted-foreground!'
                  )}
                >
                  <ComboboxValue placeholder="Select a shinobi..." />
                </div>
                {actor ? (
                  <div className={cn('flex', !is_active && 'opacity-50')}>
                    {(Object.keys(actor.natures) as Array<NatureSet>)
                      .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                      .map((nature) => (
                        <NatureBadge
                          key={nature}
                          nature={nature}
                          className="text-xs block"
                        />
                      ))}
                  </div>
                ) : (
                  <div className="text-xs text-muted-foreground/50 -mt-1">
                    to add to your team
                  </div>
                )}
              </div>
            </div>
            <div className="text-muted-foreground pr-3">
              <ChevronsUpDown />
            </div>
          </Button>
        }
      />

      <ComboboxContent className="min-w-(--anchor-width) w-(--anchor-width) max-w-(--anchor-width)">
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No Shinobi found.</ComboboxEmpty>
        <ComboboxList>
          {(actor) => (
            <ComboboxItem
              key={actor.actor_ID}
              value={actor}
              disabled={selected.includes(actor.actor_ID)}
            >
              {actor.name}
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { ActorCombobox }
