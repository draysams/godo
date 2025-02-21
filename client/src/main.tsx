import { ChakraProvider, defaultSystem } from "@chakra-ui/react"
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ThemeProvider } from "next-themes"
import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"


const queryClient = new QueryClient();
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
    <ChakraProvider value={defaultSystem}>
      <ThemeProvider attribute="class" disableTransitionOnChange>
        <App />
      </ThemeProvider>
      </ChakraProvider>
      </QueryClientProvider>
  </React.StrictMode>,
)