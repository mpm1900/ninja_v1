import {
  checkActorStat,
  type Actor,
  type ActorBaseStat,
} from '#/lib/game/actor'
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from '@tanstack/react-table'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from './ui/table'
import { cn } from '#/lib/utils'

function Stat({ stat, actor }: { actor: Actor; stat: ActorBaseStat }) {
  return (
    <span>
      <span
        className={cn({
          'text-green-400': checkActorStat(actor, stat) === 1,
          'text-red-300': checkActorStat(actor, stat) === -1,
        })}
      >
        {actor.stats[stat]}
      </span>{' '}
      <span className="text-muted-foreground">({actor.base_stats[stat]})</span>
    </span>
  )
}

const helper = createColumnHelper<Actor>()
const columns = [
  helper.accessor('name', {}),
  helper.accessor('stats.hp', {
    cell: (props) => <Stat actor={props.row.original} stat="hp" />,
  }),
  helper.accessor('stats.stamina', {
    cell: (props) => <Stat actor={props.row.original} stat="stamina" />,
  }),
  helper.accessor('stats.speed', {
    cell: (props) => <Stat actor={props.row.original} stat="speed" />,
  }),
  helper.accessor('stats.ninjutsu', {
    cell: (props) => <Stat actor={props.row.original} stat="ninjutsu" />,
  }),
  helper.accessor('stats.genjutsu', {
    cell: (props) => <Stat actor={props.row.original} stat="genjutsu" />,
  }),
  helper.accessor('stats.taijutsu', {
    cell: (props) => <Stat actor={props.row.original} stat="taijutsu" />,
  }),
]

function ActorsTable({ data }: { data: Array<Actor> }) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  })

  return (
    <Table>
      <TableHeader>
        {table.getHeaderGroups().map((hg) => (
          <tr key={hg.id}>
            {hg.headers.map((header) => (
              <TableHead key={header.id} colSpan={header.colSpan}>
                {flexRender(
                  header.column.columnDef.header,
                  header.getContext()
                )}
              </TableHead>
            ))}
          </tr>
        ))}
      </TableHeader>
      <TableBody>
        {table.getRowModel().rows.map((row) => (
          <TableRow key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <TableCell key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

export { ActorsTable }
