import {
    Table,
    Spinner,
    Box,
    Text,
  } from "@chakra-ui/react"
import { JSX } from "react"
import { Link } from "react-router-dom"
import {useClients} from "../../src/components/helperFunc/Queries"
  

export default function Demo(): JSX.Element {
  const { data: clients, isLoading, isError, error } = useClients()

  if (isLoading) {
    return (
      <Box p={4}>
        <Spinner />
        <Text ml={2}>Loading...</Text>
      </Box>
    )
  }

  if (isError) {
    return (
      <Box p={4}>
        <Text color="red.500">Failed to load data: {(error as Error).message}</Text>
      </Box>
    )
  }

  return (
    <Table.Root size="sm">
      <Table.Header>
        <Table.Row>
        <Table.ColumnHeader >Name</Table.ColumnHeader>
          <Table.ColumnHeader>Address</Table.ColumnHeader>
          <Table.ColumnHeader>Message</Table.ColumnHeader>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {clients?.map((client) => (
          <Table.Row key={client.phone}>
            <Table.Cell style={{ textDecoration: 'underline' }}>
              <Link to={`/customer/${client.phone}`}>
                {client.name}
              </Link>
            </Table.Cell>
            <Table.Cell>{client.address}</Table.Cell>
            <Table.Cell>{client.message}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  )
  
}
