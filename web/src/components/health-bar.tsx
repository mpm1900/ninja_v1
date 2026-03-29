import { useEffect, useMemo } from 'react'
import { motion, useMotionValue, useSpring, useTransform } from 'motion/react'
import type { Actor } from '#/lib/game/actor'

function clamp01(value: number) {
  return Math.max(0, Math.min(1, value))
}

function HealthBar({
  actor,
  selected = false,
}: {
  actor: Actor
  selected?: boolean
}) {
  const maxHp = Math.max(1, actor.stats.hp)
  const maxChakra = Math.max(1, actor.stats.chakra)

  const hpCurrent = Math.max(0, actor.stats.hp - actor.damage)
  const chakraCurrent = Math.max(0, actor.stats.chakra - actor.chakra_damage)

  const hpRatio = clamp01(hpCurrent / maxHp)
  const chakraRatio = clamp01(chakraCurrent / maxChakra)

  const hpTarget = useMotionValue(hpRatio * 100)
  const chakraTarget = useMotionValue(chakraRatio * 100)
  const hpGhostTarget = useMotionValue(hpRatio * 100)

  useEffect(() => {
    hpTarget.set(hpRatio * 100)
    chakraTarget.set(chakraRatio * 100)
    hpGhostTarget.set(hpRatio * 100)
  }, [hpRatio, chakraRatio, hpTarget, chakraTarget, hpGhostTarget])

  const hpWidth = useTransform(
    useSpring(hpTarget, { stiffness: 280, damping: 30, mass: 0.8 }),
    (v) => `${v}%`
  )

  const hpGhostWidth = useTransform(
    useSpring(hpGhostTarget, { stiffness: 50, damping: 26, mass: 1.2 }),
    (v) => `${v}%`
  )

  const chakraWidth = useTransform(
    useSpring(chakraTarget, { stiffness: 200, damping: 26, mass: 0.95 }),
    (v) => `${v}%`
  )

  const hpFillStyle = useMemo(() => {
    if (hpRatio > 0.8) {
      return 'bg-gradient-to-r from-emerald-500 to-emerald-700'
    }
    if (hpRatio > 0.6) {
      return 'bg-gradient-to-r from-green-600 to-lime-400/50'
    }
    if (hpRatio > 0.3) {
      return 'bg-gradient-to-r from-yellow-300 via-orange-400 to-yellow-600'
    }
    if (hpRatio > 0.2) {
      return 'bg-gradient-to-r from-amber-300 via-orange-400 to-orange-600'
    }
    return 'bg-gradient-to-r from-rose-500 via-rose-600 to-rose-800'
  }, [hpRatio])

  return (
    <div className="relative w-full select-none">
      {selected && (
        <div className="pointer-events-none absolute -inset-1 rounded-sm bg-gradient-to-r from-cyan-400/15 via-indigo-400/10 to-red-500/15 blur-md" />
      )}

      <div className="relative h-6 overflow-hidden rounded-sm border border-white/10 bg-slate-900/95">
        {/* CHAKRA LAYER (behind, border-like energy shell) */}
        <div className="absolute inset-0 bg-slate-800" />
        <motion.div
          className={`absolute inset-y-0 left-0 rounded-sm bg-gradient-to-r from-sky-300 via-sky-200 to-indigo-300 ${selected ? 'shadow-lg ring-2 ring-sky-300/40' : ''}`}
          style={{ width: chakraWidth }}
        />

        {/* Inner chamber where HP lives */}
        <div className="absolute inset-px overflow-hidden rounded-sm border border-black/50 bg-black/35 shadow-inner">
          {/* low-contrast track */}
          <div className="absolute inset-0 bg-gradient-to-b from-white/5 to-black/25" />

          {/* delayed damage ghost */}
          <motion.div
            className="absolute inset-y-0 left-0 bg-white/40"
            style={{ width: hpGhostWidth }}
          />

          {/* main HP fill */}
          <motion.div
            className={`absolute inset-y-0 left-0 ${hpFillStyle} ${selected ? 'shadow-md' : ''}`}
            style={{ width: hpWidth }}
          />

          {/* moving highlight for premium sheen */}
          {selected && (
            <motion.div
              className="absolute inset-y-0 left-[-30%] w-[30%] bg-gradient-to-r from-transparent via-white/20 to-transparent mix-blend-screen"
              animate={{ x: ['0%', '460%'] }}
              transition={{ duration: 2.6, ease: 'linear', repeat: Infinity }}
            />
          )}

          {/* subtle grid texture */}
          <div className="absolute inset-0 bg-white/20 opacity-[0.15]" />
        </div>

        {/* text + values */}
        <div className="pointer-events-none absolute inset-0 z-10 flex items-center justify-end px-2">
          <div className="flex flex-row-reverse items-center gap-2">
            <span className="text-sm font-black tabular-nums text-white drop-shadow">
              {Math.round(hpRatio * 100)}%
            </span>
          </div>
        </div>
      </div>
    </div>
  )
}

export { HealthBar }
