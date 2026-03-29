import { mutationOptions, useMutation, useQueryClient } from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import { getRequest, setResponseHeader } from '@tanstack/react-start/server'
import { useRouter } from '@tanstack/react-router'

const logout = createServerFn().handler(async () => {
  const request = getRequest()
  const cookies = request?.headers.get('cookie') || ''

  const response = await fetch(`${process.env.API_URL}/api/auth/logout`, {
    method: 'POST',
    headers: {
      Cookie: cookies,
    },
  })

  const setCookie = response.headers.get('set-cookie')
  if (setCookie) {
    setResponseHeader('set-cookie', setCookie)
  }
})

function useLogout() {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation(
    mutationOptions({
      mutationKey: ['logout'],
      mutationFn: async () => {
        await logout()
      },
      onSuccess: () => {
        queryClient.setQueryData(['me'], null)
        router.navigate({ to: '/login' })
      },
    })
  )
}

export { useLogout }
