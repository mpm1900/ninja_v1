import { ACTOR_FOCUS_DETAILS, actorFocuses, type ActorFocus } from '#/lib/game/actor'
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
  ...props
}: React.ComponentProps<typeof Field> & {
  value: ActorFocus
  onValueChange: (value: ActorFocus) => void
}) {
  return (
    <Field {...props}>
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
                  {focus} (+{ACTOR_FOCUS_DETAILS[focus].up}, -{ACTOR_FOCUS_DETAILS[focus].down})
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
