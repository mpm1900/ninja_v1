import type { Action } from '#/lib/game/action'
import {
  createColumnHelper,
  flexRender,
  functionalUpdate,
  getCoreRowModel,
  getSortedRowModel,
  useReactTable,
  type Row,
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
import { Fragment } from 'react/jsx-runtime'
import { NatureBadge } from './nature-badge'
import { useState, type ReactNode } from 'react'
import { StatBadge } from './stat-badge'
import { Checkbox } from './ui/checkbox'
import { Button } from './ui/button'

const helper = createColumnHelper<Action>()
const columns = [
  helper.display({
    id: 'select',
    header: ({ table }) =>
      `${table.getSelectedRowModel().rows.length}/${(table.options.meta as any).total}`,
    cell: ({ row, table }) => (
      <Checkbox
        checked={row.getIsSelected()}
        disabled={
          !row.getIsSelected() &&
          (row.original.locked ||
            !row.getCanSelect() ||
            (table.options.meta as any).total ==
            table.getSelectedRowModel().rows.length)
        }
      />
    ),
  }),
  helper.accessor('config.name', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        Name
      </Button>
    ),
  }),
  helper.accessor('config.nature', {
    header: 'Nature',
    cell: ({ row }) =>
      row.original.config.nature ? (
        <NatureBadge nature={row.original.config.nature} />
      ) : (
        '-'
      ),
  }),
  helper.accessor('config.stat', {
    header: 'Stat',
    cell: ({ row }) =>
      row.original.config.stat ? (
        <StatBadge
          stat={row.original.config.stat}
          contentProps={{ side: 'right' }}
        />
      ) : (
        '-'
      ),
  }),
  helper.accessor('config.power', {
    header: 'Power',
    cell: ({ row }) => row.original.config.power ?? '-',
  }),
  helper.accessor('config.accuracy', {
    header: 'Accuracy',
    cell: ({ row }) =>
      row.original.config.accuracy ? `${row.original.config.accuracy}%` : '-',
  }),
  helper.accessor('config.cooldown', {
    header: 'C/D',
    cell: ({ row }) => row.original.config.cooldown ?? '-',
  }),
  helper.accessor('config.description', {
    id: 'description',
    header: 'Description',
    cell: ({ row }) => (
      <span className="block truncate">
        {row.original.config.description}
      </span>
    ),
  }),
]

function ActionsTable({
  total,
  data,
  rowSelection,
  onRowSelectionChange,
  subRow,
}: {
  total: number
  data: Action[]
  rowSelection: RowSelectionState
  onRowSelectionChange: (rowSelection: RowSelectionState) => void
  subRow?: (props: { row: Row<Action> }) => ReactNode
}) {
  const [sorting, setSorting] = useState([{ id: 'config_name', desc: false }])

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    enableRowSelection: total > Object.keys(rowSelection).length,
    onRowSelectionChange: (updater) => {
      onRowSelectionChange(functionalUpdate(updater, rowSelection))
    },
    onSortingChange: (updater) => {
      setSorting(functionalUpdate(updater, sorting))
    },
    getRowId: (a) => a.ID,
    state: {
      rowSelection,
      sorting,
    },
    meta: {
      total,
    },
  })

  return (
    <Table>
      <TableHeader>
        {table.getHeaderGroups().map((hg) => (
          <tr key={hg.id}>
            {hg.headers.map((header) => (
              <TableHead
                key={header.id}
                colSpan={header.colSpan}
                className={header.column.id === 'description' ? 'w-full' : ''}
              >
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
            <TableRow onClick={() => row.toggleSelected()}>
              {row.getVisibleCells().map((cell) => (
                <TableCell
                  key={cell.id}
                  className={cell.column.id === 'description' ? 'w-full max-w-0' : ''}
                >
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
            {subRow && row.getCanSelect() && row.getIsSelected() && (
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
