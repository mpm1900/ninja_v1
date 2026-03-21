import {
  checkActorStat,
  type Actor,
  type ActorBaseStat,
} from '#/lib/game/actor'
import {
  createColumnHelper,
  flexRender,
  functionalUpdate,
  getCoreRowModel,
  useReactTable,
  type RowSelectionState,
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
import { Checkbox } from './ui/checkbox'

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
  helper.accessor('ID', {
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllRowsSelected()
            ? true
            : table.getIsSomeRowsSelected()
              ? 'indeterminate'
              : false
        }
        onCheckedChange={(checked) => {
          table.toggleAllRowsSelected(!!checked)
        }}
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        disabled={!row.getCanSelect()}
        onCheckedChange={(checked) => {
          row.toggleSelected(!!checked)
        }}
      />
    ),
  }),
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

function ActorsTable({
  data,
  enabled,
  rowSelection,
  onRowSelectionChange,
}: {
  data: Array<Actor>
  enabled: boolean
  rowSelection: RowSelectionState
  onRowSelectionChange: (rowSelection: RowSelectionState) => void
}) {
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
    getRowId: (a) => a.actor_ID,
    enableRowSelection: enabled,
    onRowSelectionChange: (updater) => {
      onRowSelectionChange(functionalUpdate(updater, rowSelection))
    },
    state: {
      rowSelection,
    },
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
