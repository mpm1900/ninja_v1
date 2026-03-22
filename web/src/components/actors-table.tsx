import { getTotalBaseStats, type Actor } from '#/lib/game/actor'
import {
  createColumnHelper,
  flexRender,
  functionalUpdate,
  getCoreRowModel,
  getExpandedRowModel,
  getSortedRowModel,
  useReactTable,
  type Row,
  type RowSelectionState,
  type SortingState,
  type Table as TableDef,
} from '@tanstack/react-table'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from './ui/table'
import { Checkbox } from './ui/checkbox'
import { ActorStatBase } from './actor-stat'
import { Button } from './ui/button'
import { ChevronDown, ChevronLeft } from 'lucide-react'
import { Fragment, useState, type ReactNode } from 'react'

type ActorsTableMeta = {
  onRowCheckedChange?: (actor: Actor, selected: boolean) => void
}

function getActorsTableMeta(
  table: TableDef<Actor>
): ActorsTableMeta | undefined {
  return table.options.meta as ActorsTableMeta | undefined
}

const helper = createColumnHelper<Actor>()
const columns = [
  helper.display({
    id: 'select',
    cell: ({ row, table }) => (
      <Checkbox
        checked={row.getIsSelected()}
        disabled={!row.getCanSelect()}
        onCheckedChange={(checked) => {
          row.toggleSelected(!!checked)
          getActorsTableMeta(table)?.onRowCheckedChange?.(
            row.original,
            !!checked
          )
        }}
      />
    ),
  }),
  helper.accessor('name', {}),
  helper.accessor('level', {}),
  helper.accessor('base_stats.hp', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        hp
      </Button>
    ),
    cell: (props) => <ActorStatBase actor={props.row.original} stat="hp" />,
  }),
  helper.accessor('base_stats.stamina', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        stamina
      </Button>
    ),
    cell: (props) => (
      <ActorStatBase actor={props.row.original} stat="stamina" />
    ),
  }),
  helper.accessor('base_stats.speed', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        speed
      </Button>
    ),
    cell: (props) => <ActorStatBase actor={props.row.original} stat="speed" />,
  }),
  helper.accessor('base_stats.ninjutsu', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        ninjutsu
      </Button>
    ),
    cell: (props) => (
      <ActorStatBase actor={props.row.original} stat="ninjutsu" />
    ),
  }),
  helper.accessor('base_stats.genjutsu', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        genjutsu
      </Button>
    ),
    cell: (props) => (
      <ActorStatBase actor={props.row.original} stat="genjutsu" />
    ),
  }),
  helper.accessor('base_stats.taijutsu', {
    header: ({ column }) => (
      <Button
        className="-ml-4"
        variant="ghost"
        onClick={() => column.toggleSorting()}
      >
        taijutsu
      </Button>
    ),
    cell: (props) => (
      <ActorStatBase actor={props.row.original} stat="taijutsu" />
    ),
  }),
  helper.display({
    header: 'total',
    cell: ({ row }) => getTotalBaseStats(row.original),
  }),
  helper.display({
    id: 'actions',
    cell: ({ row }) => (
      <Button
        size="icon"
        variant="ghost"
        disabled={!row.getCanExpand()}
        onClick={row.getToggleExpandedHandler()}
      >
        {row.getIsExpanded() ? <ChevronDown /> : <ChevronLeft />}
      </Button>
    ),
  }),
]

function ActorsTable({
  data,
  enabled,
  rowSelection,
  onRowSelectionChange,
  onRowCheckedChange,
  subRow,
}: {
  data: Array<Actor>
  enabled: boolean
  rowSelection: RowSelectionState
  onRowSelectionChange?: (rowSelection: RowSelectionState) => void
  onRowCheckedChange?: (actor: Actor, selected: boolean) => void
  subRow?: (props: { row: Row<Actor> }) => ReactNode
}) {
  // const [expanded, setExpanded] = useState<ExpandedState>({})
  const [sorting, setSorting] = useState<SortingState>([])
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
    getExpandedRowModel: getExpandedRowModel(),
    getRowCanExpand: (row) => row.getIsSelected(),
    getRowId: (a) => a.actor_ID,
    getSortedRowModel: getSortedRowModel(),
    enableRowSelection: enabled,
    // onExpandedChange: setExpanded,
    onRowSelectionChange: (updater) => {
      onRowSelectionChange?.(functionalUpdate(updater, rowSelection))
    },
    onSortingChange: setSorting,
    state: {
      expanded: rowSelection,
      rowSelection,
      sorting,
    },
    meta: {
      onRowCheckedChange,
    } as const,
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
            <TableRow>
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
            {row.getIsExpanded() && subRow && (
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

export { ActorsTable }
