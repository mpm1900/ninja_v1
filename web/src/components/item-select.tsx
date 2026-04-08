import type { Actor } from "#/lib/game/actor";
import { itemsQuery } from "#/lib/queries/items";
import { useQuery } from "@tanstack/react-query";
import { Field, FieldContent, FieldLabel } from "./ui/field";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { sendContextMessage } from "#/lib/stores/socket";
import { NULL_CONTEXT } from "#/lib/game/context";
import { useStore } from "@tanstack/react-store";
import { clientsStore } from "#/lib/stores/clients";

function ItemSelect({ actor }: { actor: Actor }) {
  const query = useQuery(itemsQuery)
  const client = useStore(clientsStore, c => c.me!)
  return (
    <Field>
      <FieldLabel>Item</FieldLabel>
      <FieldContent>
        <Select value={actor.item?.ID} onValueChange={item_ID => {
          sendContextMessage({
            type: 'update-actor',
            actor_config: {
              item_ID,
            },
            context: {
              ...NULL_CONTEXT,
              source_actor_ID: actor.ID,
            },
            client_ID: client.ID,
          })
        }}>
          <SelectTrigger className='w-full'>
            <SelectValue placeholder="Select an Item " />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              {query.data?.map((ability) => (
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

export { ItemSelect }
