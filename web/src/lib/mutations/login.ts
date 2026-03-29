import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import { setResponseHeader } from '@tanstack/react-start/server'
import z from 'zod'
import type { User } from '../queries/auth'

const requestSchema = z.object({
  email: z.string(),
  password: z.string(),
})

const login = createServerFn()
  .inputValidator(requestSchema)
  .handler(async ({ data }) => {
    const response = await fetch(`${process.env.API_URL}/api/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      throw new Error(`Login failed with status ${response.status}`)
    }

    const setCookie = response.headers.get('set-cookie')
    if (setCookie) {
      setResponseHeader('set-cookie', setCookie)
    }

    return (await response.json()) as User
  })

function useLogin() {
  const queryClient = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['login'],
      mutationFn: async (data: z.output<typeof requestSchema>) => {
        return await login({ data })
      },
      onSuccess: (user) => {
        queryClient.setQueryData(['me'], user)
      },
    })
  )
}

export { useLogin }
