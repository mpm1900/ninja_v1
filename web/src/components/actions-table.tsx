import type { Action } from '#/lib/game/action'
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
  type Row,
} from '@tanstack/react-table'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from './ui/table'
import { Fragment } from 'react/jsx-runtime'
import { NatureBadge } from './nature-badge'
import { Circle, CircleCheck, CircleX } from 'lucide-react'
import type { ReactNode } from 'react'
import type { Actor } from '#/lib/game/actor'

const helper = createColumnHelper<Action>()
const columns = (selected: string | undefined, cooldowns: Actor['action_cooldowns']) => [
  helper.display({
    id: 'select',
    cell: ({ row }) =>
      !row.getCanSelect() ? (
        <CircleX className="size-4 opacity-0" />
      ) : row.original.ID === selected ? (
        <CircleCheck className="size-4 text-muted-foreground" />
      ) : (
        <Circle className="size-4 text-muted-foreground/40" />
      ),
  }),
  helper.accessor('config.name', {
    header: 'name',
  }),
  helper.accessor('config.nature', {
    header: 'nature',
    cell: ({ row }) =>
      row.original.config.nature ? (
        <NatureBadge nature={row.original.config.nature} />
      ) : (
        '-'
      ),
  }),
  helper.accessor('config.stat', {
    header: 'stat',
    cell: ({ row }) => row.original.config.stat ?? '-',
  }),
  helper.accessor('config.power', {
    header: 'power',
    cell: ({ row }) => row.original.config.power ?? '-',
  }),
  helper.accessor('config.accuracy', {
    header: 'accuracy',
    cell: ({ row }) =>
      row.original.config.accuracy ? `${row.original.config.accuracy}%` : '-',
  }),
  helper.accessor('config.cooldown', {
    header: 'cooldown',
    cell: ({ row }) => row.original.config.cooldown ?? '-',
  }),
  helper.display({
    header: 'active cooldown',
    cell: ({ row }) => cooldowns[row.original.ID] ?? '-'
  })
]

function ActionsTable({
  cooldowns,
  data,
  enabled,
  selected,
  onSelectedChange,
  subRow,
}: {
  cooldowns: Actor['action_cooldowns']
  data: Action[]
  enabled: boolean
  selected: string | undefined
  onSelectedChange: (selected: string) => void
  subRow?: (props: { row: Row<Action> }) => ReactNode
}) {
  const table = useReactTable({
    data,
    columns: columns(selected, cooldowns),
    getCoreRowModel: getCoreRowModel(),
    enableRowSelection: enabled,
    getRowId: (a) => a.ID,
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
          <Fragment key={row.id}>
            <TableRow
              onClick={() => {
                if (!enabled) return
                onSelectedChange(row.original.ID)
              }}
            >
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
            {subRow && row.getCanSelect() && row.original.ID === selected && (
              <tr>
                <td colSpan={row.getAllCells().length}>{subRow({ row })}</td>
              </tr>
            )}
          </Fragment>
        ))}
      </TableBody>
    </Table>
  )
}

export { ActionsTable }
