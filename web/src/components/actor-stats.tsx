import type { Actor } from '#/lib/game/actor'
import { useForm, useStore } from '@tanstack/react-form'
import { ActorStat, NatureDamageStat, NatureResistanceStat } from './actor-stat'
import { NatureBadge } from './nature-badge'
import { Input } from './ui/input'
import z from 'zod'
import { Button } from './ui/button'
import { NULL_CONTEXT } from '#/lib/game/context'
import { clientsStore } from '#/lib/stores/clients'
import { sendContextMessage } from '#/lib/stores/socket'

const AuxStatsSchema = z
  .object({
    hp: z.number().min(0, 'negative').max(32, 'too big'),
    attack: z.number().min(0, 'negative').max(32, 'too big'),
    stamina: z.number().min(0, 'negative').max(32, 'too big'),
    defense: z.number().min(0, 'negative').max(32, 'too big'),
    speed: z.number().min(0, 'negative').max(32, 'too big'),
    chakra_attack: z.number().min(0, 'negative').max(32, 'too big'),
    chakra_defense: z.number().min(0, 'negative').max(32, 'too big'),
  })
  .superRefine((stats, ctx) => {
    const total = Object.values(stats).reduce((sum, value) => sum + value, 0)
    if (total > 66) {
      ctx.addIssue({
        code: 'custom',
        message: 'Total aux stats cannot exceed 66',
      })
    }
  })

function ActorAuxStats({ actor }: { actor: Actor }) {
  const client = useStore(clientsStore, (c) => c.me!)
  const form = useForm({
    defaultValues: actor.aux_stats,
    validators: {
      // @ts-ignore
      onChange: AuxStatsSchema,
    },
    onSubmit: ({ value }) => {
      console.log(value)
      sendContextMessage({
        type: 'update-actor',
        client_ID: client.ID,
        context: {
          ...NULL_CONTEXT,
          source_actor_ID: actor.ID,
        },
        actor_config: {
          aux_stats: value,
        },
      })
    },
  })
  return (
    <div>
      <form.Subscribe>
        {({ values }) => (
          <div className="flex justify-center">
            {Object.values(values).reduce((sum, value) => sum + value, 0)} / 66
          </div>
        )}
      </form.Subscribe>
      <div className="grid grid-cols-2">
        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground text-nowrap">HP</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'hp'} />
          </div>

          <form.Field name="hp">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Attack</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'attack'} />
          </div>

          <form.Field name="attack">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Stamina</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'stamina'} />
          </div>
          <form.Field name="stamina">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Defense</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'defense'} />
          </div>
          <form.Field name="defense">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Speed</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'speed'} />
          </div>
          <form.Field name="speed">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground text-nowrap">C.Attack</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'chakra_attack'} />
          </div>
          <form.Field name="chakra_attack">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Accuracy</div>
          <div className="whitespace-nowrap pr-4">
            <ActorStat actor={actor} showBase={false} stat={'accuracy'} />
          </div>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground text-nowrap">C.Defense</div>
          <div className="whitespace-nowrap">
            <ActorStat actor={actor} showBase={false} stat={'chakra_defense'} />
          </div>
          <form.Field name="chakra_defense">
            {(field) => (
              <Input
                type="number"
                className="w-16"
                value={field.state.value}
                onValueChange={(v) => field.handleChange(Number(v ?? 0))}
              />
            )}
          </form.Field>
        </div>

        <div className="flex justify-end items-center gap-2 mb-1">
          <div className="text-muted-foreground">Evasion</div>
          <div className="whitespace-nowrap pr-4">
            <ActorStat actor={actor} showBase={false} stat={'evasion'} />
          </div>
        </div>

        <div className="flex justify-end">
          <form.Subscribe>
            {({ canSubmit }) => (
              <Button disabled={!canSubmit} onClick={() => form.handleSubmit()}>
                Save
              </Button>
            )}
          </form.Subscribe>
        </div>
      </div>
    </div>
  )
}

function ActorStats({ actor }: { actor: Actor }) {
  return (
    <div className="flex items-start gap-4">
      <ActorAuxStats actor={actor} />
      <div>
        <div className="uppercase text-muted-foreground font-bold text-center">
          DMG
        </div>
        {Object.keys(actor.nature_damage).map((key) => (
          <div key={key} className="flex justify-between gap-1">
            <div className="text-muted-foreground">
              <NatureBadge nature={key as keyof typeof actor.nature_damage} />
            </div>
            <div>
              <NatureDamageStat
                actor={actor}
                nature={key as keyof typeof actor.nature_damage}
              />
            </div>
          </div>
        ))}
      </div>
      <div>
        <div className="uppercase text-muted-foreground font-bold text-center">
          RES
        </div>
        {Object.keys(actor.resolved_nature_resistance).map((key) => (
          <div key={key} className="flex justify-between gap-1">
            <div className="text-muted-foreground">
              <NatureBadge
                nature={key as keyof typeof actor.resolved_nature_resistance}
              />
            </div>
            <div>
              <NatureResistanceStat
                actor={actor}
                nature={key as keyof typeof actor.resolved_nature_resistance}
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export { ActorStats }
