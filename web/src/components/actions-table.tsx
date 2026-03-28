import type { Action } from '#/lib/game/action'
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
import { Fragment } from 'react/jsx-runtime'
import { NatureBadge } from './nature-badge'
import { Circle, CircleCheck, CircleX } from 'lucide-react'
import type { ReactNode } from 'react'

const helper = createColumnHelper<Action>()
const columns = (selected: string | undefined) => [
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
]

function ActionsTable({
  data,
  enabled,
  selected,
  onSelectedChange,
  subRow,
}: {
  data: Action[]
  enabled: boolean
  selected: string | undefined
  onSelectedChange: (selected: string) => void
  subRow?: ReactNode
}) {
  const table = useReactTable({
    data,
    columns: columns(selected),
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
                <td colSpan={row.getAllCells().length}>{subRow}</td>
              </tr>
            )}
          </Fragment>
        ))}
      </TableBody>
    </Table>
  )
}

export { ActionsTable }
