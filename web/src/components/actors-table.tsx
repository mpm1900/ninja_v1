import { type Actor } from '#/lib/game/actor'
import {
  createColumnHelper,
  flexRender,
  functionalUpdate,
  getCoreRowModel,
  getExpandedRowModel,
  useReactTable,
  type ExpandedState,
  type Row,
  type RowSelectionState,
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
import { ActorStat } from './actor-stat'
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
  helper.accessor('stats.hp', {
    cell: (props) => <ActorStat actor={props.row.original} stat="hp" />,
  }),
  helper.accessor('stats.stamina', {
    cell: (props) => <ActorStat actor={props.row.original} stat="stamina" />,
  }),
  helper.accessor('stats.speed', {
    cell: (props) => <ActorStat actor={props.row.original} stat="speed" />,
  }),
  helper.accessor('stats.ninjutsu', {
    cell: (props) => <ActorStat actor={props.row.original} stat="ninjutsu" />,
  }),
  helper.accessor('stats.genjutsu', {
    cell: (props) => <ActorStat actor={props.row.original} stat="genjutsu" />,
  }),
  helper.accessor('stats.taijutsu', {
    cell: (props) => <ActorStat actor={props.row.original} stat="taijutsu" />,
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
  onRowSelectionChange: (rowSelection: RowSelectionState) => void
  onRowCheckedChange?: (actor: Actor, selected: boolean) => void
  subRow?: (props: { row: Row<Actor> }) => ReactNode
}) {
  //const [expanded, setExpanded] = useState<ExpandedState>({})
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
    getExpandedRowModel: getExpandedRowModel(),
    getRowCanExpand: (row) => row.getIsSelected(),
    getRowId: (a) => a.actor_ID,
    enableRowSelection: enabled,
    // onExpandedChange: setExpanded,
    onRowSelectionChange: (updater) => {
      onRowSelectionChange(functionalUpdate(updater, rowSelection))
    },
    state: {
      expanded: rowSelection,
      rowSelection,
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
