import type { TeamBuilderForm } from '#/hooks/use-team-builder-form'
import type { ActorDef } from '#/lib/game/actor'
import {
  getResistance,
  getWeakness,
  natureIndexes,
  type NatureSet,
} from '#/lib/game/nature'
import { NatureBadge } from './nature-badge'
import { TeamBuilderActorAttributes } from './team-builder-actor-attributes'
import { TeamBuilderStats } from './team-builder-stats'

function NatureWR({ nature }: { nature: NatureSet }) {
  const weaknesses = getWeakness(nature)
  const resistances = getResistance(nature)
  if (weaknesses.length == 0 && resistances.length == 0) {
    return null
  }
  return (
    <div>
      <NatureBadge nature={nature} className="text-xs" />
      {weaknesses.length > 0 && (
        <>
          <span className="text-xs text-muted-foreground"> is weak to </span>
          {weaknesses?.map((n) => (
            <NatureBadge key={n} nature={n} className="text-xs" />
          ))}
        </>
      )}

      {resistances.length > 0 && (
        <>
          <span className="text-xs text-muted-foreground"> but resists </span>
          {resistances?.map((n) => (
            <NatureBadge key={n} nature={n} className="text-xs" />
          ))}
        </>
      )}
    </div>
  )
}

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
        return {
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
      {({ actor, selected_index, total }) => (
        <div className="flex gap-4">
          <div className="flex flex-col gap-2 min-w-1/4">
            <div className="flex">
              <img src={def.sprite_url} className="object-cover size-16" />
              <div>
                <div>{def.name}</div>
                <div className="flex">
                  {(Object.keys(def.natures) as Array<NatureSet>)
                    .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                    .map((nature) => (
                      <NatureBadge
                        key={nature}
                        nature={nature}
                        className="text-xs block"
                      />
                    ))}
                </div>
              </div>
            </div>
            <TeamBuilderActorAttributes
              def={def}
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
            <div>
              <div className="flex justify-between items-center">
                {(Object.keys(def.natures) as Array<NatureSet>)
                  .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                  .map((nature) => (
                    <NatureBadge key={nature} nature={nature} />
                  ))}
                <div className="text-muted-foreground text-xs flex-1 text-center">
                  {' '}
                  is weak to{' '}
                </div>
                {getWeakness(...(Object.keys(def.natures) as Array<NatureSet>))
                  .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                  .map((nature) => (
                    <NatureBadge key={nature} nature={nature} />
                  ))}
              </div>
              <div className="flex justify-between items-center">
                {(Object.keys(def.natures) as Array<NatureSet>)
                  .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                  .map((nature) => (
                    <NatureBadge key={nature} nature={nature} />
                  ))}
                <div className="text-muted-foreground text-xs flex-1 text-center">
                  {' '}
                  resists{' '}
                </div>
                {getResistance(
                  ...(Object.keys(def.natures) as Array<NatureSet>)
                )
                  .sort((a, b) => natureIndexes[a] - natureIndexes[b])
                  .map((nature) => (
                    <NatureBadge key={nature} nature={nature} />
                  ))}
              </div>
            </div>
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
