import type { Context } from "#/lib/game/context"
import { clientsStore } from "#/lib/stores/clients"
import { sendContextMessage, subscribeSocketMessages } from "#/lib/stores/socket"
import { Store, useStore } from "@tanstack/react-store"
import { useEffect } from "react"

const validateContextStore = new Store<{ valid: boolean, loading: boolean }>({
  valid: false,
  loading: false,
})

function useValidateContext(context: Context, prompt_ID?: string) {
  const client = useStore(clientsStore, (c) => c.me)
  const store = useStore(validateContextStore, s => s)

  useEffect(() => {
    validateContextStore.setState(() => ({ valid: false, loading: false }))
    return subscribeSocketMessages((_, message) => {
      if (message?.type !== 'validate-context') return
      validateContextStore.setState(() => ({
        valid: message.valid ?? false,
        loading: false,
      }))
    })
  }, [])

  useEffect(() => {
    if (!client || validateContextStore.get().loading) return

    validateContextStore.setState(s => ({
      ...s,
      loading: true
    }))
    sendContextMessage({
      type: 'validate-context',
      context: context,
      client_ID: client.ID,
      prompt_ID: prompt_ID
    })
  }, [context])


  return store
}

export { useValidateContext }
