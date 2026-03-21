import type { ModifierTransaction } from '#/lib/game/modifier'
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
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
import { Button } from './ui/button'
import { Trash } from 'lucide-react'

type ModifiersTableMeta = {
  onRowRemove?: (modifier: ModifierTransaction) => void
}

function getModifiersTableMeta(
  table: TableDef<ModifierTransaction>
): ModifiersTableMeta | undefined {
  return table.options.meta as ModifiersTableMeta | undefined
}

const helper = createColumnHelper<ModifierTransaction>()
const columns = [
  helper.accessor('mutation.name', {}),
  helper.accessor('ID', {}),
  helper.display({
    id: 'remove',
    cell: ({ row, table }) => {
      const meta = getModifiersTableMeta(table)
      return (
        <Button
          size="icon"
          variant="ghost"
          onClick={() => meta?.onRowRemove?.(row.original)}
        >
          <Trash />
        </Button>
      )
    },
  }),
]

function ModifiersTable({
  data,
  onRowRemove,
}: {
  data: Array<ModifierTransaction>
} & ModifiersTableMeta) {
  const table = useReactTable({
    columns,
    data,
    getCoreRowModel: getCoreRowModel(),
    meta: {
      onRowRemove,
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

export { ModifiersTable }
