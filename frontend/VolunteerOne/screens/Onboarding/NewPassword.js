import React from "react";
import { useState } from "react";
import {
  StyleSheet,
  ImageBackground,
  Dimensions,
  StatusBar,
  KeyboardAvoidingView,
  Image,
  TextInput,
  TouchableOpacity
} from "react-native";
import { Block, Text } from "galio-framework";

import { Button } from "../../components";
import { Images, argonTheme } from "../../constants";

import logo from "../../assets/logo/logo2.png";

const { width, height } = Dimensions.get("screen");

/** ==================================== New Password Screen ==================================== **/

const NewPassword = ({ navigation }) => {
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const handleNewPasswordInput = (input) => {
    setNewPassword(input);
  };

  const handleConfirmPasswordInput = (input) => {
    setConfirmPassword(input);
  };

  const handleSaveBtnClick = () => {
    console.log(newPassword, confirmPassword);
    navigation.navigate('Login');
  };

  return (
    <Block flex middle>
      <Block style={styles.container}>
        <Text h4 style={styles.title}>
          Reset Password
        </Text>
        <Block width={width * 0.8} style={styles.inputContainer}>
          <Text style={styles.inputLabel}>New Password</Text>
          <Block width={width * 0.8}>
            <TextInput
              style={styles.input}
              secureTextEntry={true}
              placeholder="New Password"
              onChangeText={handleNewPasswordInput}
            />
          </Block>
        </Block>
        <Block width={width * 0.8} style={styles.inputContainer}>
          <Text style={styles.inputLabel}>Confirm Password</Text>
          <Block width={width * 0.8}>
            <TextInput
              style={styles.input}
              secureTextEntry={true}
              placeholder="Confirm Password"
              onChangeText={handleConfirmPasswordInput}
            />
          </Block>
        </Block>
        <Block middle>
          <Button
            style={styles.button}
            onPress={handleSaveBtnClick}
            color="#009688"
          >
            <Text style={styles.buttonText}>RESET PASSWORD</Text>
          </Button>
        </Block>
      </Block>
    </Block>
  );
};
// }

const styles = StyleSheet.create({
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
  },
  instructionText: {
    flexDirection: "row",
  },
  loginContainer: {
    width: width * 0.9,
    height: height * 0.75,
    backgroundColor: "#F4F5F7",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    shadowRadius: 8,
    shadowOpacity: 0.1,
    elevation: 1,
    overflow: "hidden",
  },
  inputIcons: {
    marginRight: 12,
  },
  passwordCheck: {
    paddingLeft: 15,
    paddingTop: 13,
    paddingBottom: 30,
  },
  createButton: {
    width: width * 0.5,
    marginTop: 25,
  },
  logo: {
    width: 265,
    height: 50,
    zIndex: 2,
    position: 'relative',
    marginTop: '20%'
  },
});

export default NewPassword;
