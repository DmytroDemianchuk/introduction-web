import React, { useState } from 'react';


import { Box, Heading, Center } from '@chakra-ui/react';
import { ChatIcon } from '@chakra-ui/icons';
import { Link } from 'react-router-dom';
import "./header.css";

function Header() {
    return (
      <Box paddingBottom={5}>
        <Center>
          <Link to="/">
            <Heading size="2xl">
               Introduction To the Web
            </Heading>
          </Link>
        </Center>
      </Box>
    );
  }

export default Header;