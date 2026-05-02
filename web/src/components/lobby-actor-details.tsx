import type { Actor } from '#/lib/game/actor'
import { natureIndexes, type NatureSet } from '#/lib/game/nature'
import { cn } from '#/lib/utils'
import { NatureBadge } from './nature-badge'

function LobbyActorDetails({
  actor,
  className,
  ...props
}: React.ComponentProps<'button'> & { actor: Actor }) {
  return (
    <button
      {...props}
      className={cn(
        'border border-stone-300/40 ring ring-black bg-stone-800 rounded text-left cursor-pointer',
        'p-2 flex flex-col gap-2',
        className
      )}
    >
      <div className="flex justify-between">
        <div className="text-lg font-bold">{actor.name}</div>
        <div className="flex items-start">
          {(Object.keys(actor.natures) as Array<NatureSet>)
            .sort((a, b) => natureIndexes[a] - natureIndexes[b])
            .map((nature) => (
              <NatureBadge key={nature} nature={nature} className="text-xs" />
            ))}
        </div>
      </div>
      <div className="grid grid-cols-3 text-sm">
        <div className="capitalize">{actor.focus}</div>
        <div>{actor.item?.name ?? '-'}</div>
        <div>{actor.ability?.name ?? '-'}</div>
      </div>
      <table className="[&_td]:px-2 [&_td]:whitespace-nowrap">
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
