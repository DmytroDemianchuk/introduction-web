import React, { Component } from 'react';
import axios from 'axios';

import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
  Input,
} from '@chakra-ui/react';


class Register extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
      message: '',
      isInvalid: '',
      endpoint: 'http://localhost:8000/auth/sign-up',
      redirect: false,
      redirectTo: '/chat?u=',
    };
  }

  // on change of input, set the value to the message state
  onChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  };

  onSubmit = async e => {
    e.preventDefault();

    try {
      const res = await axios.post(this.state.endpoint, {
        username: this.state.username,
        password: this.state.password,
      });

      console.log('register', res);
      if (res.data.status) {
        const redirectTo = this.state.redirectTo + this.state.username;
        this.setState({ redirect: true, redirectTo });
      } else {
        // on failed
        this.setState({ message: res.data.message, isInvalid: true });
      }
    } catch (error) {
      console.log(error);
      this.setState({ message: 'something went wrong', isInvalid: true });
    }
  };

  render() {
    return (
      <div>

        <div marginBlockStart={10} textAlign={'left'} maxW="2xl">
          <div borderRadius="lg" padding={10} borderWidth="2px">
            <div spacing={5}>
              <FormControl isInvalid={this.state.isInvalid}>
                <FormLabel>Username</FormLabel>
                <Input
                  type="text"
                  placeholder="Username"
                  name="username"
                  value={this.state.username}
                  onChange={this.onChange}
                />

                {!this.state.isInvalid ? (
                  <></>
                ) : (
                  <FormErrorMessage>{this.state.message}</FormErrorMessage>
                )}
                {/* <FormHelperText>use a unique username</FormHelperText> */}
              </FormControl>
              <FormControl>
                <FormLabel>Password</FormLabel>
                <Input
                  type="password"
                  placeholder="Password"
                  name="password"
                  value={this.state.password}
                  onChange={this.onChange}
                />
                <FormHelperText>use a dummy password</FormHelperText>
              </FormControl>
              <div
                size="lg"
                colorScheme="green"
                variant="solid"
                type="submit"
                onClick={this.onSubmit}
              >
                Register
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Register;
