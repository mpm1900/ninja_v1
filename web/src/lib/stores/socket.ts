import { Store } from '@tanstack/store'
import { clientsStore, type Client } from './clients'
import { gameStore } from './game'
import type { Game } from '../game/game'
import type { Context } from '../game/context'
import { setContextPlayer } from './battle-context'
import type { ActorFocus } from '../game/actor'

type SocketStatus =
  | 'idle'
  | 'connecting'
  | 'open'
  | 'closing'
  | 'closed'
  | 'error'

type SocketState = {
  instanceID: string | null
  socket: WebSocket | null
  status: SocketStatus
  url: string | null
}

type SocketRequestType =
  | 'add-actor'
  | 'remove-actor'
  | 'update-actor'
  | 'push-action'
  | 'remove-action'
  | 'run-game-actions'
  | 'resolve-prompt'
  | 'validate-state'
  | 'validate-context'
  | 'get-targets'

type SocketRequest = {
  type: SocketRequestType
  prompt_ID?: string
  client_ID: string
  context: Context
  actor_config?: {
    ability_ID?: string
    action_IDs?: Array<string>
    focus?: ActorFocus
  }
}

type SocketResponse = {
  type: 'game' | 'clients' | 'join-success' | 'validate-context' | 'target-IDs'
  state: Game | null
  clients: Array<Client> | null
  context: Context | null
  valid: boolean | null
}

type SocketMessageSubscriber = (
  event: MessageEvent,
  message: SocketResponse | null
) => void

function getSocketUrl(instanceID: string): string {
  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
  const hostname = window.location.hostname
  const port = '3005'

  return `${protocol}://${hostname}:${port}/socket/${instanceID}/connect`
}

const socketStore = new Store<SocketState>({
  instanceID: null,
  socket: null,
  status: 'idle',
  url: null,
})

const messageSubscribers = new Set<SocketMessageSubscriber>()

function isCurrentSocket(socket: WebSocket): boolean {
  return socketStore.state.socket === socket
}

function subscribeSocketMessages(subscriber: SocketMessageSubscriber) {
  messageSubscribers.add(subscriber)
  return () => {
    messageSubscribers.delete(subscriber)
  }
}

function clearSocketEventHandlers(socket: WebSocket) {
  socket.onopen = null
  socket.onclose = null
  socket.onerror = null
  socket.onmessage = null
}

function connectSocket(instanceID: string) {
  if (!instanceID) return

  const previous = socketStore.state.socket
  if (previous) {
    clearSocketEventHandlers(previous)
    if (
      previous.readyState === WebSocket.CONNECTING ||
      previous.readyState === WebSocket.OPEN
    ) {
      previous.close(1000, 'Switching instance')
    }
  }

  const url = getSocketUrl(instanceID)
  const socket = new WebSocket(url)

  socketStore.setState((s) => ({
    ...s,
    instanceID,
    socket,
    status: 'connecting',
    url,
  }))

  socket.onopen = () => {
    if (!isCurrentSocket(socket)) return
    socketStore.setState((s) => ({
      ...s,
      status: 'open',
    }))
  }

  socket.onmessage = (event) => {
    if (!isCurrentSocket(socket)) return

    let message: SocketResponse | null = null

    if (typeof event.data === 'string') {
      try {
        message = JSON.parse(event.data) as SocketResponse
      } catch {
        console.error('non-JSON payload')
      }
    }

    if (message?.type === 'game') {
      gameStore.setState(() => message.state!)
    }
    if (message?.type === 'clients') {
      clientsStore.setState((c) => ({
        ...c,
        clients: message.clients!,
      }))
    }
    if (message?.type === 'join-success') {
      gameStore.setState(() => message.state!)
      clientsStore.setState((c) => ({
        ...c,
        me: message.clients![0],
      }))
      setContextPlayer(message.clients![0].ID)
    }

    for (const subscriber of messageSubscribers) {
      try {
        subscriber(event, message)
      } catch (error) {
        console.error('socket message subscriber error', error)
      }
    }
  }

  socket.onerror = () => {
    if (!isCurrentSocket(socket)) return
    socketStore.setState((s) => ({
      ...s,
      status: 'error',
    }))
  }

  socket.onclose = () => {
    if (!isCurrentSocket(socket)) return
    socketStore.setState((s) => ({
      ...s,
      socket: null,
      status: 'closed',
    }))
  }
}

function disconnectSocket(code = 1000, reason = 'Manual disconnect') {
  const socket = socketStore.state.socket
  if (!socket) return

  socketStore.setState((s) => ({
    ...s,
    status: 'closing',
  }))

  if (
    socket.readyState === WebSocket.CONNECTING ||
    socket.readyState === WebSocket.OPEN
  ) {
    socket.close(code, reason)
  } else {
    socketStore.setState((s) => ({
      ...s,
      socket: null,
      status: 'closed',
    }))
  }
}

function sendSocketMessage(
  payload: string | ArrayBufferLike | Blob | ArrayBufferView
): boolean {
  const socket = socketStore.state.socket
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    return false
  }

  socket.send(payload)
  return true
}

function sendContextMessage(request: SocketRequest) {
  return sendSocketMessage(JSON.stringify(request))
}

export type { SocketResponse, SocketMessageSubscriber }
export {
  socketStore,
  connectSocket,
  disconnectSocket,
  sendSocketMessage,
  sendContextMessage,
  subscribeSocketMessages,
}
