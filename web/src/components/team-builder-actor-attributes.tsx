import type { ActorDef, ActorFocus } from '#/lib/game/actor'
import { AbilitySelect } from './ability-select'
import { FocusSelect } from './focus-select'
import { ItemSelect } from './item-select'

function TeamBuilderActorAttributes({
  def,
  focus,
  onFocusChange,
  abilityID,
  onAbilityIDChange,
  itemID,
  onItemIDChange,
}: {
  def: ActorDef
  focus: ActorFocus
  onFocusChange: (f: ActorFocus) => void
  abilityID: string | undefined | null
  onAbilityIDChange: (a: string) => void
  itemID: string | undefined | null
  onItemIDChange: (i: string | null) => void
}) {
  return (
    <div className="space-y-2">
      <FocusSelect value={focus ?? 'none'} onValueChange={onFocusChange} />
      <AbilitySelect
        options={def.abilities}
        value={abilityID ?? null}
        onValueChange={onAbilityIDChange}
      />
      <ItemSelect value={itemID ?? null} onValueChange={onItemIDChange} />
    </div>
  )
}

export { TeamBuilderActorAttributes }
