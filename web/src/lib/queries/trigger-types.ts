import { createServerFn } from '@tanstack/react-start'
import { queryOptions } from '@tanstack/react-query'

const getTriggers = createServerFn().handler(async () => {
  const response = await fetch(`${process.env.API_URL}/api/triggers`)
  const data = await response.json()
  return data as Array<string>
})

const triggersQuery = queryOptions({
  queryKey: ['triggers'],
  queryFn: () => getTriggers(),
})

export { triggersQuery, getTriggers }
