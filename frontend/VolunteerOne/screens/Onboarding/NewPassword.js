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

const NewPassword = () => {
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const [confirmPasswordError, setConfirmPasswordError] = useState("");

  const handlePasswordChange = (text) => {
    setPassword(text);
  };

  const handleConfirmPasswordChange = (text) => {
    setConfirmPassword(text);
  };

  const handleSavePassword = () => {
    if (password === "") {
      setPasswordError("Please enter a password");
      return;
    }

    if (confirmPassword === "") {
      setConfirmPasswordError("Please confirm your password");
      return;
    }

    if (password !== confirmPassword) {
      setConfirmPasswordError("Passwords do not match");
      return;
    }

    // Save password logic goes here
    console.log("Password saved!");
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Create New Password</Text>
      <TextInput
        style={styles.input}
        placeholder="New Password"
        onChangeText={handlePasswordChange}
        value={password}
        secureTextEntry={true}
      />
      <Text style={styles.error}>{passwordError}</Text>
      <TextInput
        style={styles.input}
        placeholder="Confirm Password"
        onChangeText={handleConfirmPasswordChange}
        value={confirmPassword}
        secureTextEntry={true}
      />
      <Text style={styles.error}>{confirmPasswordError}</Text>
      <TouchableOpacity style={styles.button} onPress={handleSavePassword}>
        <Text style={styles.buttonText}>Save Password</Text>
      </TouchableOpacity>
    </View>
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
});

export default NewPassword;
