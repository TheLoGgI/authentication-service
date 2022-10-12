import "./App.css"

import { Box, Heading } from "@chakra-ui/react"
import { Header, Layout } from "./components"

import { Outlet } from "react-router-dom"

function App() {
  return (
    <>
      <Header />
      <Layout>
        <Outlet />
      </Layout>
    </>
  )
}

export default App
