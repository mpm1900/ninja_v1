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
} from './ui/combobox'
import { CircleQuestionMark } from 'lucide-react'
import { InputGroupAddon } from './ui/input-group'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { NatureBadge } from './nature-badge'
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
      <ComboboxInput
        className={cn(
          'h-14 gap-2',
          is_active && 'border-neutral-400 bg-input!',
          className
        )}
        before={
          actor ? (
            <InputGroupAddon className="relative size-12 cursor-pointer">
              <img
                src={actor.sprite_url}
                className="h-full w-full object-cover absolute inset-0 z-10 rounded-lg ml-1"
                onClick={onClick}
              />
            </InputGroupAddon>
          ) : (
            <InputGroupAddon>
              <CircleQuestionMark />
            </InputGroupAddon>
          )
        }
        placeholder="Select a Shinobi"
      >
        {actor && (
          <div>
            {(Object.keys(actor.natures) as Array<NatureSet>)
              .sort((a, b) => natureIndexes[a] - natureIndexes[b])
              .map((nature) => (
                <NatureBadge key={nature} nature={nature} />
              ))}
          </div>
        )}
      </ComboboxInput>

      <ComboboxContent>
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
