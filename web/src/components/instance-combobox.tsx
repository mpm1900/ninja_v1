import { instancesQuery } from '#/lib/queries/instances'
import { useQuery } from '@tanstack/react-query'
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
  ComboboxTrigger,
  ComboboxValue,
} from './ui/combobox'
import { Button } from './ui/button'
import type { ReactNode } from 'react'
import { v4 } from 'uuid'
import { cn } from '#/lib/utils'

function InstanceCombobox({
  icon,
  value = null,
  onValueChange,
}: {
  icon: ReactNode
  value: string | undefined | null
  onValueChange: (value: string) => void
}) {
  const query = useQuery(instancesQuery)

  return (
    <Combobox
      items={[...(query.data ?? []), { ID: null }]}
      value={value}
      onValueChange={(v) => {
        if (v) {
          onValueChange(v)
        } else {
          onValueChange(v4())
        }
        query.refetch()
      }}
    >
      <ComboboxTrigger
        render={
          <Button
            variant="outline"
            className={cn("min-w-80 justify-start font-normal font-mono", {
              'text-destructive': !value
            })}
          >
            {icon}
            <ComboboxValue placeholder='NOT CONNECTED' />
          </Button>
        }
      />
      < ComboboxContent >
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No items found.</ComboboxEmpty>
        <ComboboxList>
          {(item) => (
            <ComboboxItem
              key={item.ID}
              value={item.ID}
              className={cn({
                "[&_[data-slot='combobox-item-indicator']]:hidden": !value,
              })}
            >
              {item.ID ?? 'Create Instance'}
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent >
    </Combobox >
  )
}

export { InstanceCombobox }
