import { NULL_CONTEXT } from '#/lib/game/context'
import { sendContextMessage } from '#/lib/stores/socket'
import { useStore } from '@tanstack/react-store'
import { Field, FieldContent, FieldLabel } from './ui/field'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select'
import type { Modifier } from '#/lib/game/modifier'

function AbilitySelect({
  value,
  onValueChange,
  options,
}: {
  value: string | null
  onValueChange: (value: string) => void
  options: Array<Modifier>
}) {
  return (
    <Field>
      <FieldLabel>Ability</FieldLabel>
      <FieldContent>
        <Select value={value ?? ''} onValueChange={onValueChange}>
          <SelectTrigger className="w-full">
            <SelectValue placeholder="Select an Ability " />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              {options.map((ability) => (
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

export { AbilitySelect }
