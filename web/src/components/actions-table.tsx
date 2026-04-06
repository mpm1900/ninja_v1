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
import { memo, useState, type ReactNode } from 'react'
import { StatBadge } from './stat-badge'
import { Checkbox } from './ui/checkbox'
import { Button } from './ui/button'

const helper = createColumnHelper<Action>()
const columns = [
  helper.display({
    id: 'select',
    cell: ({ row }) => (
      <Checkbox checked={row.getIsSelected()} disabled={row.original.locked || !row.getCanSelect()} />
    ),
  }),
  helper.accessor('config.name', {
    header: ({ column }) => <Button
      className="-ml-4"
      variant="ghost"
      onClick={() => column.toggleSorting()}
    >
      Name
    </Button>
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
]

function ActionsTable({
  data,
  enabled,
  rowSelection,
  onRowSelectionChange,
  subRow,
}: {
  data: Action[]
  enabled: boolean
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
    enableRowSelection: enabled,
    onRowSelectionChange: (updater) => {
      onRowSelectionChange(functionalUpdate(updater, rowSelection))
    },
    onSortingChange: updater => {
      setSorting(functionalUpdate(updater, sorting))
    },
    getRowId: (a) => a.ID,
    state: {
      rowSelection,
      sorting,
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
          <Fragment key={row.id}>
            <TableRow
              onClick={
                enabled
                  ? () => row.toggleSelected()
                  : undefined
              }
            >
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
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

const MemoizedActionsTable = memo(ActionsTable)

export { MemoizedActionsTable as ActionsTable }
