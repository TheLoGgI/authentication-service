import {
  Box,
  Button,
  Input as ChakraInput,
  FormControl,
  FormHelperText,
  FormLabel,
  InputProps,
  Text,
} from "@chakra-ui/react"
import { forwardRef, useRef } from "react"

import { NavLink } from "react-router-dom"
import { Props } from "../components"

const Input: React.FC<any> = forwardRef((props, ref) => {
  return <ChakraInput bgColor="white" ref={ref} {...props} />
})

const FormField: React.FC<Props> = ({ children, ...props }) => {
  return (
    <FormControl mt="4" {...props}>
      {children}
    </FormControl>
  )
}

export function Login() {
  const usernameEmail = useRef<HTMLInputElement | undefined>()
  const password = useRef<HTMLInputElement | undefined>()

  const handleSubmit: React.FormEventHandler<HTMLFormElement> | undefined = (
    event
  ) => {
    event.preventDefault()
    const options = {
      method: "POST",
      mode: "no-cors",
      body: new FormData(event.currentTarget),
      //   credentials: "include",
      //   body: JSON.stringify({
      //     email: usernameEmail.current?.value,
      //     password: password.current?.value,
      //   }),
      headers: {
        Accept: "application/json",
        Authorization: "Basic bGFzc2VBYWtqYWVyOnNlY3JldFBhc3N3b3Jk",
        "Content-Type": "application/json",
      },
    }

    console.log("options: ", options)
    fetch(`http://127.0.0.1:5000/authorize/login`, options).then(console.log)
    // .then((res) => res.())
  }

  return (
    <Box p="4" w="md" bg="gray.100" borderRadius="8">
      <form onSubmit={handleSubmit}>
        <Text fontSize="xl" fontWeight="medium">
          Login Form
        </Text>
        <FormField>
          <FormLabel>Username / Email</FormLabel>
          <Input
            ref={usernameEmail}
            type="text"
            placeholder="JohnDoe / john-doe@gmail.com.."
          />
        </FormField>
        <FormField>
          <FormLabel>Password</FormLabel>
          <Input ref={password} type="password" />
        </FormField>
        <Button colorScheme="blue" type="submit" mt="8" w="full">
          Login
        </Button>
        <NavLink to="/register">
          <Button
            colorScheme="blue"
            variant="outline"
            type="button"
            mt="4"
            w="full"
          >
            SignUp
          </Button>
        </NavLink>
      </form>
    </Box>
  )
}

export function SignUp() {
  const username = useRef<HTMLInputElement | undefined>()
  const password = useRef<HTMLInputElement | undefined>()
  const email = useRef<HTMLInputElement | undefined>()

  const handleSubmit: React.FormEventHandler<HTMLFormElement> | undefined = (
    event
  ) => {
    event.preventDefault()

    fetch(`http://127.0.0.1:5000/create/user`, {
      method: "POST",
      body: JSON.stringify({
        username: username.current?.value,
        password: password.current?.value,
        email: email.current?.value,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then(console.log)
  }

  return (
    <Box p="4" w="md" bg="gray.100" borderRadius="8">
      <form onSubmit={handleSubmit}>
        <Text fontSize="xl" fontWeight="medium">
          Sign Up Form
        </Text>
        <FormField>
          <FormLabel>Username</FormLabel>
          <Input type="text" ref={username} placeholder="John Doe" />
          <FormHelperText>What do you what to be called?</FormHelperText>
        </FormField>
        <FormField>
          <FormLabel>Email</FormLabel>
          <Input type="email" ref={email} placeholder="john-doe@gmail.com" />
          <FormHelperText>We'll never share your email.</FormHelperText>
        </FormField>
        <FormField>
          <FormLabel>Password</FormLabel>
          <Input type="password" ref={password} />
          <FormHelperText>
            Password has to be at least 8 characters long and have both numbers
            and letters
          </FormHelperText>
        </FormField>
        <FormField>
          <FormLabel>Re-Password</FormLabel>
          <Input type="password" />
          <FormHelperText>Confirm your password</FormHelperText>
        </FormField>
        <Button colorScheme="blue" type="submit" mt="8" w="full">
          Register Account
        </Button>
      </form>
    </Box>
  )
}
