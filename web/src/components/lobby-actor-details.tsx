import { MODIFIER_ICONS, SHINOBI_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { NatureBadge } from './nature-badge'

function LobbyActorDetails({
  actor,
  className,
  enabled,
  ...props
}: React.ComponentProps<'button'> & { actor: Actor; enabled: boolean }) {
  const ItemIcon = actor?.item?.icon
    ? MODIFIER_ICONS[actor.item.icon]
    : undefined
  const AbilityIcon = actor?.ability?.icon
    ? MODIFIER_ICONS[actor.ability.icon]
    : undefined
  return (
    <button
      {...props}
      className={cn(
        'relative overflow-hidden border border-stone-300/30 ring ring-black bg-stone-900 rounded-lg text-left cursor-pointer',
        enabled && 'bg-stone-600',
        'p-2 flex flex-col gap-2',
        className
      )}
    >
      <img
        src={actor.sprite_url}
        draggable={false}
        className={cn('absolute left-0 bottom-0 opacity-40')}
        width={128}
        height={128}
      />
      <div className="absolute z-0 opacity-30 -right-5 -bottom-7">
        {actor.affiliations
          ?.filter((_, i) => i == 0)
          .map((a) => {
            const C = SHINOBI_ICONS[a]
            return C ? <C key={a} className="w-36" /> : null
          })}
      </div>
      <div className="flex justify-between z-10">
        <div className="text-3xl nanum-brush-script-regular text-shadow-[2px_2px_0px_#000000]">
          {actor.name}
        </div>
        <div className="flex items-start">
          {(Object.keys(actor.natures) as Array<NatureSet>)
            .sort((a, b) => natureIndexes[a] - natureIndexes[b])
            .map((nature) => (
              <NatureBadge
                key={nature}
                nature={nature}
                className="block text-xs"
              />
            ))}
        </div>
      </div>
      <div className="z-10 bg-stone-700 rounded-xs overflow-hidden ring ring-black mb-2 text-shadow-[1px_1px_0px_#000000]">
        <div className="mb-1 h-px w-full bg-gradient-to-r to-stone-100/35 from-transparent" />
        <div className="flex gap-8 [&>div]:flex-1 [&>div]:text-nowrap px-2">
          <div className="capitalize">{actor.focus}</div>
          <div className="flex gap-1 items-center">
            {ItemIcon && <ItemIcon />}
            {actor.item?.name ?? '-'}
          </div>
          <div className="flex gap-1 items-center">
            {AbilityIcon && <AbilityIcon />}
            {actor.ability?.name ?? '-'}
          </div>
        </div>
        <div className="mt-1 h-px w-full bg-gradient-to-r to-transparent from-stone-100/35" />
      </div>
      <table className="z-10 [&_td]:px-2 [&_td]:whitespace-nowrap text-shadow-[1px_1px_0px_#000000]">
        <tbody>
          {actor.actions
            .filter((a) => !a.meta.switch)
            .map((a) => (
              <tr
                key={a.ID}
                className={cn({
                  'text-destructive': a.disabled,
                })}
              >
                <td className="w-6">
                  {a.config.nature && <NatureBadge nature={a.config.nature} />}
                </td>
                <td>{a.config.name}</td>
              </tr>
            ))}
        </tbody>
      </table>
    </button>
  )
}

export { LobbyActorDetails }
