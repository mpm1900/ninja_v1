import { queryOptions } from "@tanstack/react-query";
import { createServerFn } from "@tanstack/react-start";

const getInstances = createServerFn().handler(async () => {
  const response = await fetch(`${process.env.API_URL}/api/instances`)
  console.log(response)
  const data = await response.json()
  return data
})

const instancesQuery = queryOptions({
  queryKey: ['instances'],
  queryFn: () => getInstances()
})

export { instancesQuery, getInstances }
