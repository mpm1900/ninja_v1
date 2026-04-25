import { Store } from '@tanstack/store'
import { clientsStore, type Client } from './clients'
import { gameStore } from './game'
import type { Game } from '../game/game'
import type { Context } from '../game/context'
import { setContextPlayer } from './battle-context'
import type { ActorConfig, TeamConfig } from './config'

type SocketStatus =
  | 'idle'
  | 'connecting'
  | 'open'
  | 'closing'
  | 'closed'
  | 'error'
  | 'reconnecting'

type SocketState = {
  instanceID: string | null
  socket: WebSocket | null
  status: SocketStatus
  url: string | null
  reconnectCount: number
  isManualDisconnect: boolean
}

type SocketRequestType =
  | 'set-team'
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
  actor_config?: Partial<ActorConfig>
  team_config?: TeamConfig
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

const INSTANCE_ID_KEY = 'ninja_instance_id'
const MAX_RECONNECT_DELAY = 30000
const INITIAL_RECONNECT_DELAY = 1000

function getSocketUrl(instanceID: string): string {
  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
  const hostname = window.location.hostname
  const port = '3005'

  return `${protocol}://${hostname}:${port}/socket/${instanceID}/connect`
}

const savedInstanceID = localStorage.getItem(INSTANCE_ID_KEY)

const socketStore = new Store<SocketState>({
  instanceID: savedInstanceID,
  socket: null,
  status: 'idle',
  url: null,
  reconnectCount: 0,
  isManualDisconnect: false,
})

const messageSubscribers = new Set<SocketMessageSubscriber>()
let reconnectTimer: number | null = null

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

function connectSocket(instanceID: string, onOpen?: () => void) {
  if (!instanceID) return

  if (reconnectTimer) {
    clearTimeout(reconnectTimer)
    reconnectTimer = null
  }

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

  localStorage.setItem(INSTANCE_ID_KEY, instanceID)
  const url = getSocketUrl(instanceID)
  const socket = new WebSocket(url)

  socketStore.setState((s) => ({
    ...s,
    instanceID,
    socket,
    status: s.reconnectCount > 0 ? 'reconnecting' : 'connecting',
    url,
    isManualDisconnect: false,
  }))

  socket.onopen = () => {
    if (!isCurrentSocket(socket)) return
    socketStore.setState((s) => ({
      ...s,
      status: 'open',
      reconnectCount: 0,
    }))
    onOpen?.()
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

  socket.onerror = (error) => {
    if (!isCurrentSocket(socket)) return
    console.error('WebSocket error:', error)
    socketStore.setState((s) => ({
      ...s,
      status: 'error',
    }))
  }

  socket.onclose = (event) => {
    if (!isCurrentSocket(socket)) return

    const { isManualDisconnect } = socketStore.state

    socketStore.setState((s) => ({
      ...s,
      socket: null,
      status: isManualDisconnect ? 'closed' : 'error',
    }))

    if (!isManualDisconnect) {
      attemptReconnect()
    }
  }
}

function attemptReconnect() {
  const { instanceID, reconnectCount, isManualDisconnect } = socketStore.state
  if (!instanceID || isManualDisconnect) return

  const delay = Math.min(
    INITIAL_RECONNECT_DELAY * Math.pow(2, reconnectCount),
    MAX_RECONNECT_DELAY
  )

  console.log(`Attempting to reconnect in ${delay}ms... (attempt ${reconnectCount + 1})`)

  socketStore.setState((s) => ({
    ...s,
    reconnectCount: s.reconnectCount + 1,
    status: 'reconnecting',
  }))

  reconnectTimer = window.setTimeout(() => {
    reconnectTimer = null
    connectSocket(instanceID)
  }, delay)
}

function disconnectSocket(code = 1000, reason = 'Manual disconnect') {
  if (reconnectTimer) {
    clearTimeout(reconnectTimer)
    reconnectTimer = null
  }

  const socket = socketStore.state.socket
  if (!socket) {
    socketStore.setState((s) => ({
      ...s,
      status: 'closed',
      instanceID: null,
      reconnectCount: 0,
      isManualDisconnect: true,
    }))
    localStorage.removeItem(INSTANCE_ID_KEY)
    return
  }

  socketStore.setState((s) => ({
    ...s,
    status: 'closing',
    isManualDisconnect: true,
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
      instanceID: null,
      reconnectCount: 0,
    }))
  }

  localStorage.removeItem(INSTANCE_ID_KEY)
}

function sendSocketMessage(
  payload: string | ArrayBufferLike | Blob | ArrayBufferView
): boolean {
  const socket = socketStore.state.socket
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    console.warn('Attempted to send message while socket is not open')
    return false
  }

  socket.send(payload)
  return true
}

function sendContextMessage(request: SocketRequest) {
  return sendSocketMessage(JSON.stringify(request))
}

// Auto-connect on load if we have a saved instance ID
if (savedInstanceID) {
  connectSocket(savedInstanceID)
}

export type { SocketResponse, SocketMessageSubscriber, ActorConfig }
export {
  socketStore,
  connectSocket,
  disconnectSocket,
  sendSocketMessage,
  sendContextMessage,
  subscribeSocketMessages,
}
