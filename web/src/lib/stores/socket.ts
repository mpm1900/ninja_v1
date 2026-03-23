import { Store } from '@tanstack/store'
import { clientsStore, type Client } from './clients'
import { gameStore } from './game'
import type { Game } from '../game/game'
import type { Context } from '../game/context'

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

type SocketRequest = {
  type: string
  context: Context
  clientID: string
  modifierID?: string
}

const DEFAULT_HOST = 'localhost:3005'

function getSocketUrl(instanceID: string): string {
  return `ws://${DEFAULT_HOST}/socket/${instanceID}/connect`
}

const socketStore = new Store<SocketState>({
  instanceID: null,
  socket: null,
  status: 'idle',
  url: null,
})

function isCurrentSocket(socket: WebSocket): boolean {
  return socketStore.state.socket === socket
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
    if (typeof event.data !== 'string') return

    try {
      const message = JSON.parse(event.data) as {
        type: 'state' | 'clients' | 'join-success'
        state: Game | null
        clients: Array<Client> | null
      }
      if (message.type === 'state') {
        gameStore.setState(() => message.state!)
      }
      if (message.type === 'clients') {
        clientsStore.setState(() => message.clients!)
      }
      if (message.type === 'join-success') {
        gameStore.setState(() => message.state!)
      }
    } catch {
      // Ignore non-JSON websocket payloads
      console.error('non-JSON payload')
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

export type {}
export {
  socketStore,
  connectSocket,
  disconnectSocket,
  sendSocketMessage,
  sendContextMessage,
}
