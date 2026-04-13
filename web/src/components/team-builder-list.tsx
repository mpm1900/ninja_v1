import type { TeamBuilderForm } from '#/hooks/use-team-builder-form'
import { makeConfigFromDef } from '#/lib/game/team'
import { ActorCombobox } from './actor-combobox'

function TeamBuilderList({ form }: { form: TeamBuilderForm }) {
  return (
    <form.Field name="actors" mode="array">
      {(field) => (
        <form.Subscribe
          selector={(state) => ({
            selected: state.values.actors.map((a) => a.actor_ID!),
            active: state.values.actors[state.values.selected_index]?.actor_ID,
          })}
        >
          {({ selected, active }) => (
            <div className="flex flex-col gap-2 min-w-sm">
              <div className="">Team: {selected.length}/6</div>
              {field.state.value.map((_, i) => (
                <form.Field key={i} name={`actors[${i}]`}>
                  {(actorID) => (
                    <ActorCombobox
                      active={active}
                      selected={selected}
                      value={actorID.state.value?.actor_ID}
                      onValueChange={(actor) => {
                        actorID.handleChange(makeConfigFromDef(actor))
                        form.setFieldValue('selected_index', i)
                      }}
                      onClick={() => form.setFieldValue('selected_index', i)}
                    />
                  )}
                </form.Field>
              ))}
              {selected.length < 6 && (
                <ActorCombobox
                  active={active}
                  selected={selected}
                  value={undefined}
                  onValueChange={(actor) => {
                    form.pushFieldValue('actors', makeConfigFromDef(actor))
                    form.setFieldValue('selected_index', selected.length)
                  }}
                />
              )}
            </div>
          )}
        </form.Subscribe>
      )}
    </form.Field>
  )
}

export { TeamBuilderList }
