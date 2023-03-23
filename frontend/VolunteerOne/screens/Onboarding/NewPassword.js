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
            color="#8898AA"
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
  container: {
      flex: 1,
      backgroundColor: "#F4F5F7",
      alignItems: "center",
      justifyContent: "center",
  },
  title: {
      marginBottom: 30,
  },
  inputContainer: {
      marginVertical: 10,
  },
  inputLabel: {
      marginBottom: 5,
      fontWeight: "bold",
  },
  input: {
      borderColor: argonTheme.COLORS.BORDER,
      borderWidth: 1,
      height: 44,
      backgroundColor: "#FFFFFF",
      shadowColor: argonTheme.COLORS.BLACK,
      shadowOffset: {
          width: 0,
          height: 1
      },
      shadowRadius: 2,
      shadowOpacity: 0.05,
      elevation: 2,
      paddingLeft: 10,
  },
  button: {
      width: width * 0.8,
      marginTop: 25,
      backgroundColor: argonTheme.COLORS.PRIMARY,
  },
  buttonText: {
      color: "#FFFFFF",
  },
  logo: {
      width: 265,
      height: 50,
      zIndex: 2,
      position: 'relative',
      marginTop: '20%',
  },
});

export default NewPassword;
