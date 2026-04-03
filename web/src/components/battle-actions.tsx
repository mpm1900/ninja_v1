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
  const has_queued_action = game.queued_actions[actor.ID]
  const idle = game.status === 'idle'
  const staged = game.actions.find(
    (tx) => tx.context.source_actor_ID === actor.ID
  )

  return (
    <LayoutGroup id="battle-actions">
      <div className="pointer-events-none relative flex w-full flex-col items-center gap-4 pb-8">
        <div
          className={cn(
            'pointer-events-auto w-xl transition-opacity duration-150',
            !idle && 'opacity-0'
          )}
        >
          <AnimatePresence mode="wait" initial={false}>
            {action && (
              <motion.div
                key={`${actor.ID}.${action.ID}`}
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
                {!staged && (
                  <div className="grid place-items-center text-muted-foreground mb-6">
                    <div>
                      {action ? action.config.name : 'Choose an Action'}
                    </div>
                  </div>
                )}
                <ActionControl
                  action={action}
                  staged={staged}
                  enabled={
                    game.status === 'idle' &&
                    !!actor.position_ID &&
                    action.cooldown == null
                  }
                  context={context}
                  onContextChange={setContext}
                />
              </motion.div>
            )}
          </AnimatePresence>
        </div>

        <div className="pointer-events-none fixed inset-x-0 -bottom-4 z-0 flex justify-center">
          <motion.div
            initial={false}
            animate={{
              y: idle && !has_queued_action ? -56 : 18,
            }}
            transition={{
              y: { type: 'spring', stiffness: 420, damping: 34, mass: 0.68 },
            }}
            className="pointer-events-none"
            aria-hidden={!idle}
          >
            <AnimatePresence mode="wait" initial={false}>
              <motion.div
                key={actor.ID}
                initial={{ y: 100, opacity: 0 }}
                animate={{ y: 0, opacity: idle && !has_queued_action ? 1 : 0 }}
                exit={{ y: 100, opacity: 0 }}
                transition={{
                  type: 'spring',
                  stiffness: 800,
                  damping: 50,
                  mass: 0.68,
                }}
                className="flex items-end -space-x-9 px-4 transform-gpu"
              >
                {actor.actions.map((a, i) => {
                  const selected = context.action_ID === a.ID
                  const center = (actor.actions.length - 1) / 2
                  const fanRotate = (i - center) * 1.25

                  return (
                    <motion.div
                      key={`${actor.ID}:${a.ID}`}
                      className={cn(
                        'relative pointer-events-auto transform-gpu will-change-transform',
                        {
                          'pointer-events-none': !!staged && !selected,
                        }
                      )}
                      style={{
                        zIndex: selected ? 80 : 20 + i,
                        willChange: 'transform',
                      }}
                      initial={{ y: 12, scale: 0.985 }}
                      animate={{
                        y: selected ? -64 : 0,
                        scale: selected ? 1.015 : 1,
                        rotate: selected ? 0 : fanRotate,
                      }}
                      exit={{ y: 10, scale: 0.985 }}
                      whileHover={{
                        y: selected ? -64 : -36,
                        scale: selected ? 1.05 : 1.02,
                        rotate: 0,
                        zIndex: 90,
                        paddingBottom: 28,
                        marginBottom: -28,
                      }}
                      whileTap={{ scale: 1.01 }}
                      transition={{
                        type: 'spring',
                        stiffness: 760,
                        damping: 38,
                        mass: 0.5,
                      }}
                    >
                      <ActionCard
                        action={a}
                        disabled={!!staged && !selected}
                        selected={selected}
                        onClick={() => setContextAction(a.ID)}
                        style={{
                          zIndex: selected ? 80 : 20 + i,
                        }}
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
