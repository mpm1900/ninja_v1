import { instancesQuery } from "#/lib/queries/instances"
import { useQuery } from "@tanstack/react-query"
import { Combobox, ComboboxContent, ComboboxEmpty, ComboboxInput, ComboboxItem, ComboboxList, ComboboxTrigger, ComboboxValue } from "./ui/combobox"
import { Button } from "./ui/button"
import type { ReactNode } from "react"

function InstanceCombobox({ icon, value = null, onValueChange }: {
  icon: ReactNode
  value: string | undefined | null,
  onValueChange: (value: string) => void
}) {
  const query = useQuery(instancesQuery)

  return (
    <Combobox items={query.data} value={value} onValueChange={v => v && onValueChange(v)}>
      <ComboboxTrigger render={<Button variant="outline" className="min-w-80 justify-start font-normal">
        {icon}<ComboboxValue placeholder='NOT_CONNECTED' /></Button>} />
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No items found.</ComboboxEmpty>
        <ComboboxList>
          {(item) => (
            <ComboboxItem key={item.ID} value={item.ID}>
              {item.ID}
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { InstanceCombobox }
