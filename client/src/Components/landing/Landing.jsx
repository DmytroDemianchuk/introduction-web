import React from 'react';
import { Link } from 'react-router-dom';
import { Box, Button, Container, Stack, Text } from '@chakra-ui/react';
import { EditIcon, ArrowForwardIcon } from '@chakra-ui/icons';
import './Landing.css'; // Import the CSS file

function Landing ()  {
  return (
    <Container className="container" maxW="2xl">
      <Text className="text">
        This is a simple chat application written in Reactjs powered by Chakra
        component. It is using websocket for communication.
      </Text>
      <Box>
        <Stack className="buttons" direction="row" spacing={7}>
          <Link to="register">
            <Button
              className="button register"
              size="lg"
              leftIcon={<EditIcon />}
              colorScheme="green"
              variant="solid"
            >
              Register
            </Button>
          </Link>
          <Link to="login">
            <Button
              className="button login"
              size="lg"
              rightIcon={<ArrowForwardIcon />}
              colorScheme="green"
              variant="outline"
            >
              Login
            </Button>
          </Link>
        </Stack>
      </Box>
    </Container>
  );
}

export default Landing;
