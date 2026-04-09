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

function ActorCombobox({
  onClick,
  selected = [],
  value,
  onValueChange,
}: {
  onClick?: () => void
  selected?: string[]
  value?: string | null
  onValueChange?: (value: string | null) => void
}) {
  const actors = useSuspenseQuery(actorsQuery)
  const sortedActors = actors.data.sort((a, b) => a.name.localeCompare(b.name))
  const selectedActor =
    sortedActors.find((actor) => actor.actor_ID === value) ?? null

  const handleValueChange = (actor: ActorDef | null) => {
    onValueChange?.(actor?.actor_ID ?? null)
  }
  return (
    <Combobox<ActorDef>
      items={sortedActors}
      itemToStringValue={(actor) => actor.actor_ID}
      itemToStringLabel={(actor) => actor.name}
      value={selectedActor}
      onValueChange={handleValueChange}
    >
      <ComboboxInput
        className="h-14 gap-2"
        before={
          selectedActor ? (
            <InputGroupAddon className="relative size-12 cursor-pointer">
              <img
                src={selectedActor.sprite_url}
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
        {selectedActor && (
          <div>
            {(Object.keys(selectedActor.natures) as Array<NatureSet>)
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
