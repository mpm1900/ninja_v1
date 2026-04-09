import { itemsQuery } from '#/lib/queries/items'
import { useQuery } from '@tanstack/react-query'
import { Field, FieldContent, FieldLabel } from './ui/field'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select'

function ItemSelect({
  value,
  onValueChange,
}: {
  value: string | null
  onValueChange: (value: string | null) => void
}) {
  const query = useQuery(itemsQuery)
  return (
    <Field>
      <FieldLabel>Item</FieldLabel>
      <FieldContent>
        <Select value={value ?? ''} onValueChange={onValueChange}>
          <SelectTrigger className="w-full">
            <SelectValue placeholder="Select an Item " />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              {query.data?.map((ability) => (
                <SelectItem key={ability.ID} value={ability.ID}>
                  {ability.name}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </FieldContent>
    </Field>
  )
}

export { ItemSelect }
