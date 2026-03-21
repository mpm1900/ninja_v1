import { Store } from '@tanstack/store'

type Client = {
  ID: string
}

const clientsStore = new Store<Array<Client>>([])

export type { Client }
export { clientsStore }
