import type { TeamBuilderForm } from '#/hooks/use-team-builder-form'
import type { ActorDef } from '#/lib/game/actor'
import {
  getResistance,
  getWeakness,
  natureIndexes,
  type NatureSet,
} from '#/lib/game/nature'
import { NatureBadge } from './nature-badge'
import { NastureSetDetails } from './natureset-details'
import { TeamBuilderActorAttributes } from './team-builder-actor-attributes'
import { TeamBuilderStats } from './team-builder-stats'

function TeamBuilderActorConfig({
  def,
  form,
}: {
  def: ActorDef
  form: TeamBuilderForm
}) {
  return (
    <form.Subscribe
      selector={(state) => {
        const actor = state.values.actors[state.values.selected_index]
        const items = state.values.actors.map(a => a.config.item_ID).filter(id => id !== null)
        return {
          items,
          selected_index: state.values.selected_index,
          actors: state.values.actors,
          actor,
          total: Object.values(actor?.config.aux_stats ?? {}).reduce(
            (sum, value) => sum + value,
            0
          ),
        }
      }}
    >
      {({ actor, selected_index, total, items }) => (
        <div className="flex gap-4">
          <div className="flex flex-col gap-2 min-w-1/4">
            <div className="flex">
              <img src={def.sprite_url} className="object-cover size-16" />
              <div>
                <div className='w-full px-2 flex gap-6 justify-between'>
                  <div>{def.name}</div>
                  <div className="flex items-start">
                    {(Object.keys(def.natures) as Array<NatureSet>)
                      .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                      .map((nature) => (
                        <NatureBadge
                          key={nature}
                          nature={nature}
                          className="text-xs"
                        />
                      ))}
                  </div>
                </div>
                <div className='px-2 text-xs text-muted-foreground'>info here</div>
              </div>
            </div>
            <TeamBuilderActorAttributes
              def={def}
              otherItemIDs={items}
              focus={actor.config?.focus ?? 'none'}
              onFocusChange={(focus) => {
                form.setFieldValue(
                  `actors[${selected_index}].config.focus`,
                  focus
                )
              }}
              abilityID={actor.config?.ability_ID}
              onAbilityIDChange={(ability_ID) => {
                form.setFieldValue(
                  `actors[${selected_index}].config.ability_ID`,
                  ability_ID
                )
              }}
              itemID={actor.config?.item_ID}
              onItemIDChange={(item_ID) => {
                form.setFieldValue(
                  `actors[${selected_index}].config.item_ID`,
                  item_ID
                )
              }}
            />
            <NastureSetDetails
              natures={Object.keys(def.natures) as Array<NatureSet>}
            />
          </div>

          <TeamBuilderStats
            total={total}
            config={actor.config}
            def={def}
            onConfigChange={(config) => {
              form.setFieldValue(`actors[${selected_index}].config`, config)
            }}
          />
        </div>
      )}
    </form.Subscribe>
  )
}

export { TeamBuilderActorConfig }
