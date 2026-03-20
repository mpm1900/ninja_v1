import { actorsQuery } from '#/lib/queries/actors'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: App,
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(actorsQuery)
  },
})

function App() {
  const query = useSuspenseQuery(actorsQuery)
  console.log(query.data)
  return (
    <main className="px-4 pb-8 pt-14">
      <pre>{JSON.stringify(query.data, null, 4)}</pre>
    </main>
  )
}
