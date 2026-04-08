import type { Actor } from '#/lib/game/actor'
import { NULL_CONTEXT } from '#/lib/game/context'
import { sendContextMessage } from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { Field, FieldContent, FieldLabel } from './ui/field'
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from './ui/select'
import { clientsStore } from '#/lib/stores/clients'

function AbilitySelect({ actor }: { actor: Actor }) {
  const client = useStore(clientsStore, s => s.me!)
  return (
    <Field>
      <FieldLabel>Ability</FieldLabel>
      <FieldContent>
        <Select value={actor.ability?.ID} onValueChange={ability_ID => {
          sendContextMessage({
            type: 'update-actor',
            actor_config: {
              ability_ID,
            },
            context: {
              ...NULL_CONTEXT,
              source_actor_ID: actor.ID,
            },
            client_ID: client.ID,
          })
        }}>
          <SelectTrigger className='w-full'>
            <SelectValue placeholder="Select an Ability " />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              {actor.abilities.map((ability) => (
                <SelectItem key={ability.ID} value={ability.ID}>
                  {ability.name}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </FieldContent>
    </Field>
  )
}

export { AbilitySelect }
