import { Card, CardContent } from './ui/card'
import { useStore } from '@tanstack/react-store'
import type { Actor } from '#/lib/game/actor'
import { ActorStats } from './actor-stats'
import { sendContextMessage } from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { NULL_CONTEXT } from '#/lib/game/context'
import { AbilitySelect } from './ability-select'
import { ItemSelect } from './item-select'

function ActorControl({ actor }: { actor: Actor }) {
  const client = useStore(clientsStore, (c) => c.me!)

  return (
    <Card className="flex flex-row rounded-t-none border-t-0 mx-2 mb-2 py-2 gap-0">
      <CardContent className="grid grid-cols-3 w-full">
        <div className="space-y-4">
          <div className="h-16 w-16 overflow-hidden">
            <img
              src={actor.sprite_url}
              className="h-full w-full object-cover"
              width={64}
              height={64}
            />
          </div>
          <AbilitySelect
            options={actor.abilities}
            value={actor.ability?.ID ?? null}
            onValueChange={(ability_ID) => {
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
            }}
          />
          <ItemSelect
            value={actor.item?.ID ?? null}
            onValueChange={(item_ID) => {
              if (!item_ID) return
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
            }}
          />
        </div>
        <ActorStats actor={actor} />
      </CardContent>
    </Card>
  )
}

export { ActorControl }
