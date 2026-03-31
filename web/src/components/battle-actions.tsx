import { ActionControl } from './action-control'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import type { Actor } from '#/lib/game/actor'
import { ActionCard } from './action-card'
import {
  battleContext,
  setContext,
  setContextAction,
} from '#/lib/stores/battle-context'
import { AnimatePresence, LayoutGroup, motion } from 'motion/react'
import { cn } from '#/lib/utils'

function BattleActions({ actor }: { actor: Actor }) {
  const game = useStore(gameStore, (g) => g)
  const context = useStore(battleContext, (c) => c)
  const action = actor.actions.find((a) => a.ID === context.action_ID)
  const idle = game.status === 'idle'
  const queued = game.actions.find(
    (tx) => tx.context.source_actor_ID === actor.ID
  )

  const pulseScale =
    game.turn.phase === 'start' ? 1.01 : game.turn.phase === 'end' ? 1.006 : 1

  return (
    <LayoutGroup id="battle-actions">
      <div className="pointer-events-none relative flex w-full flex-col items-center gap-4 pb-8">
        <div
          className="pointer-events-auto w-xl transition-opacity duration-150"
          style={{
            pointerEvents: idle ? 'auto' : 'none',
            opacity: idle ? 1 : 0,
          }}
        >
          {!queued && (
            <div className="grid place-items-center text-muted-foreground mb-6">
              <div>{action ? 'Select Targets' : 'Choose an Action'}</div>
            </div>
          )}

          <AnimatePresence mode="wait" initial={false}>
            {action && (
              <motion.div
                key={action.ID}
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                transition={{
                  type: 'spring',
                  stiffness: 520,
                  damping: 36,
                  mass: 0.6,
                }}
              >
                <ActionControl
                  action={action}
                  queued={queued}
                  enabled={
                    game.status === 'idle' &&
                    !!actor.position_ID &&
                    actor.action_cooldowns[action.ID] == undefined
                  }
                  context={context}
                  onContextChange={setContext}
                />
              </motion.div>
            )}
          </AnimatePresence>
        </div>

        <div className="pointer-events-none fixed inset-x-0 -bottom-32 z-0 flex justify-center">
          <motion.div
            initial={false}
            animate={{
              y: idle ? -56 : 18,
              scale: idle ? [1, pulseScale, 1] : 1,
            }}
            transition={{
              y: { type: 'spring', stiffness: 420, damping: 34, mass: 0.68 },
              scale: { duration: 0.2, ease: 'easeOut' },
            }}
            className="pointer-events-none"
            style={{ pointerEvents: idle ? 'auto' : 'none' }}
            aria-hidden={!idle}
          >
            <AnimatePresence mode="wait" initial={false}>
              <motion.div
                key={actor.ID}
                initial={{ y: 10, opacity: 0 }}
                animate={{ y: 0, opacity: 1 }}
                exit={{ y: -10, opacity: 0 }}
                transition={{
                  type: 'spring',
                  stiffness: 480,
                  damping: 34,
                  mass: 0.68,
                }}
                className="flex items-end -space-x-9 px-4"
              >
                {actor.actions.map((a, i) => {
                  const selected = context.action_ID === a.ID
                  const center = (actor.actions.length - 1) / 2
                  const fanRotate = (i - center) * 1.25

                  return (
                    <motion.div
                      key={`${actor.ID}:${a.ID}`}
                      layout
                      className={cn('relative pointer-events-auto', {
                        'pointer-events-none': !!queued && !selected,
                      })}
                      initial={{ y: 12, opacity: 0, scale: 0.985 }}
                      animate={{
                        y: selected ? -64 : 0,
                        scale: selected ? 1.015 : 1,
                        rotate: selected ? 0 : fanRotate,
                        opacity: idle ? 1 : 0,
                      }}
                      exit={{ y: 10, opacity: 0, scale: 0.985 }}
                      whileHover={{
                        y: selected ? -64 : -36,
                        scale: selected ? 1.025 : 1.02,
                        rotate: 0,
                        opacity: 1,
                        zIndex: 90,
                      }}
                      whileTap={{ scale: 1.01 }}
                      transition={{
                        type: 'spring',
                        stiffness: 760,
                        damping: 38,
                        mass: 0.5,
                      }}
                      style={{ zIndex: selected ? 80 : 20 + i }}
                    >
                      <ActionCard
                        action={a}
                        disabled={!!queued && !selected}
                        selected={selected}
                        onClick={() => setContextAction(a.ID)}
                      />
                    </motion.div>
                  )
                })}
              </motion.div>
            </AnimatePresence>
          </motion.div>
        </div>
      </div>
    </LayoutGroup>
  )
}

export { BattleActions }
