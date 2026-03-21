import { createServerFn } from '@tanstack/react-start'
import { queryOptions } from '@tanstack/react-query'
import type { Modifier } from '../game/modifier'

const getModifiers = createServerFn().handler(async () => {
  const response = await fetch(`${process.env.API_URL}/api/modifiers`)
  const data = await response.json()
  return data as Array<Modifier>
})

const modifiersQuery = queryOptions({
  queryKey: ['modifiers'],
  queryFn: () => getModifiers(),
})

export { modifiersQuery, getModifiers }
