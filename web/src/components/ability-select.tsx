import { Field, FieldContent, FieldLabel } from './ui/field'
import { Select, SelectTrigger, SelectValue } from './ui/select'

function AbilitySelect() {
  return (
    <Field>
      <FieldLabel>Ability</FieldLabel>
      <FieldContent>
        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Select an Ability " />
          </SelectTrigger>
        </Select>
      </FieldContent>
    </Field>
  )
}

export { AbilitySelect }
