import { mutationOptions, useMutation, useQueryClient } from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import { setResponseHeader } from '@tanstack/react-start/server'
import z from 'zod'
import type { User } from '#/lib/queries/auth'

const requestSchema = z.object({
  email: z.string(),
  password: z.string(),
})

const signup = createServerFn()
  .inputValidator(requestSchema)
  .handler(async ({ data }) => {
    const response = await fetch(`${process.env.API_URL}/api/auth/signup`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    const setCookie = response.headers.get('set-cookie')
    if (setCookie) {
      setResponseHeader('set-cookie', setCookie)
    }

    if (!response.ok) {
      throw new Error(`Signup failed with status ${response.status}`)
    }

    return (await response.json()) as User
  })

function useSignup() {
  const queryClient = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['signup'],
      mutationFn: async (data: z.output<typeof requestSchema>) => {
        return await signup({ data })
      },
      onSuccess: (user) => {
        queryClient.setQueryData(['me'], user)
      },
    })
  )
}

export { useSignup }
