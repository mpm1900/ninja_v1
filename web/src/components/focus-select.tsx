import { actorFocuses, type ActorFocus } from '#/lib/game/actor'
import { Field, FieldContent, FieldLabel } from './ui/field'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from './ui/select'

function FocusSelect({
  value,
  onValueChange,
}: {
  value: ActorFocus
  onValueChange: (value: ActorFocus) => void
}) {
  return (
    <Field className="w-40">
      <FieldLabel>Focus</FieldLabel>
      <FieldContent>
        <Select value={value} onValueChange={onValueChange}>
          <SelectTrigger className="w-full">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              {actorFocuses.map((focus) => (
                <SelectItem key={focus} value={focus}>
                  {focus}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </FieldContent>
    </Field>
  )
}

export { FocusSelect }
