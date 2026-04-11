import { ActionControl } from './action-control'
import { useStore } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import type { Actor } from '#/lib/game/actor'
import { battleContext, setContext } from '#/lib/stores/battle-context'
import { AnimatePresence, LayoutGroup, motion } from 'motion/react'
import { BattleCards } from './battle-cards'

function BattleActions({ actor }: { actor: Actor }) {
  const status = useStore(gameStore, (g) => g.status)
  const actions = useStore(gameStore, (g) => g.actions)
  const context = useStore(battleContext, (c) => c)
  const action = actor.actions.find((a) => a.ID === context.action_ID)
  const idle = status === 'idle'
  const staged = actions.find((tx) => tx.context.source_actor_ID === actor.ID)
  const action_locked = actor.action_locked && actor.last_used_action_ID != null
  const action_enabled = action?.cooldown == null && !action?.disabled
  const is_switch = action?.config.name == 'Switch'
  const switch_locked = actor.switch_locked && is_switch
  const choice_locked =
    action_locked && actor.last_used_action_ID !== action?.ID

  return (
    <LayoutGroup id="battle-actions">
      <div className="pointer-events-none relative flex w-full flex-col items-center gap-4 pb-8">
        <div className="pointer-events-auto">
          <AnimatePresence mode="wait" initial={false}>
            {action && idle && (
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
                  <div className="grid place-items-center mb-6 nanum-brush-script-regular text-5xl text-shadow-lg">
                    <div>
                      {action ? action.config.name : 'Choose an Action'}
                    </div>
                  </div>
                )}
                <ActionControl
                  action={action}
                  staged={staged}
                  enabled={
                    idle &&
                    !!actor.position_ID &&
                    action_enabled &&
                    !((choice_locked && !is_switch) || switch_locked)
                  }
                  context={context}
                  onContextChange={setContext}
                />
              </motion.div>
            )}
          </AnimatePresence>
        </div>
        <BattleCards actor={actor} />
      </div>
    </LayoutGroup>
  )
}

export { BattleActions }
