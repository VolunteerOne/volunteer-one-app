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

/** ==================================== Login Screen ==================================== **/

const Login = ({ navigation }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  function handleEmailInput(input) {
    setEmail(input);
  }

  function handlePasswordInput(input) {
    setPassword(input);
  }

  function handleLoginBtnClick() {
    console.log(email, password);
    navigation.navigate("App");
  }

  return (
    <Block flex middle>
      <StatusBar hidden />
      <ImageBackground
        source={Images.RegisterBackground}
        style={{ width, height, zIndex: 1 }}
      >
        <Block safe flex middle>
          <Block style={styles.loginContainer}>
            <Block flex>
              <Block flex={0.5} middle style={styles.instructionText}>
                <Image source={logo} />
              </Block>
              <Block flex={0.17} middle style={styles.instructionText}>
                <TouchableOpacity
                  onPress={() => navigation.navigate('CreateAccount')}
                >
                  <Text
                    color="#8898AA"
                    size={12}
                    style={{
                      fontWeight: "bold",
                      textDecorationLine: "underline",
                      paddingRight: 5,
                    }}
                  >
                    Create Account
                  </Text>
                </TouchableOpacity>
                <Text color="#8898AA" size={12}>
                  or Login with credentials
                </Text>
              </Block>
              <Block flex center>
                <KeyboardAvoidingView
                  style={{ flex: 1 }}
                  behavior="padding"
                  enabled
                >
                  <Block width={width * 0.8} style={{ marginBottom: 15 }}>
                    <TextInput
                      style={styles.input}
                      placeholder="Email"
                      onChangeText={handleEmailInput}
                    />
                  </Block>
                  <Block width={width * 0.8}>
                    <TextInput
                      secureTextEntry={true}
                      style={styles.input}
                      placeholder="Password"
                      onChangeText={handlePasswordInput}
                    />
                    <Block row style={styles.passwordCheck}>
                      <TouchableOpacity
                        onPress={() => navigation.navigate('ForgotPassword')}
                      >
                        <Text
                          color="#8898AA"
                          size={12}
                          style={{
                            textDecorationLine: "underline",
                          }}
                        >
                          Forgot Password?
                        </Text>
                      </TouchableOpacity>
                    </Block>
                  </Block>
                  <Block middle>
                    <Button
                      color="primary"
                      style={styles.createButton}
                      onPress={handleLoginBtnClick}
                    >
                      <Text bold size={14} color={argonTheme.COLORS.WHITE}>
                        LOGIN
                      </Text>
                    </Button>
                  </Block>
                </KeyboardAvoidingView>
              </Block>
            </Block>
          </Block>
        </Block>
      </ImageBackground>
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
});

export default Login;
