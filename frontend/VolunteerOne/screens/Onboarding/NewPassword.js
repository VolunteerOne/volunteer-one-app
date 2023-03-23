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
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 20,
  },
  input: {
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 4,
    padding: 10,
    marginVertical: 10,
    minWidth: 300,
  },
  error: {
    color: "red",
    marginBottom: 10,
  },
  button: {
    backgroundColor: "blue",
    padding: 10,
    borderRadius: 4,
  },
  buttonText: {
    color: "white",
    fontWeight: "bold",
  },
});

export default NewPassword;
