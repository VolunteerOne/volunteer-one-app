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
} from "react-native";
import { Block, Checkbox, Text, theme } from "galio-framework";

import { Button, Icon, Input } from "../../components";
import { Images, argonTheme } from "../../constants";

const { width, height } = Dimensions.get("screen");

import logo from "../../assets/imgs/argon-logo.png";

/** ==================================== Register Screen ==================================== **/

const Register = ({ navigation }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [passwordStrength, setPasswordStrength] = useState('weak');
  const [showPasswordStrength, setShowPasswordStrength] = useState(false);

  function handleNameInput(input) {
    setName(input);
  }

  function handleEmailInput(input) {
    setEmail(input);
  }

  function handlePasswordInput(input) {
    setPassword(input);
    setShowPasswordStrength(true);
  }

  function handleRegisterBtnClick() {
    setShowPasswordStrength(false);
    console.log(name, email, password);
  }

  return (
    <Block flex middle>
      <StatusBar hidden />
      <ImageBackground
        source={Images.RegisterBackground}
        style={{ width, height, zIndex: 1 }}
      >
        <Block safe flex middle>
          <Block style={styles.registerContainer}>
            <Block flex>
              <Block flex={0.25} middle style={styles.instructionText}>
                <Image source={logo} />
              </Block>
              <Block flex={0.17} middle>
                <Text color="#8898AA" size={12}>
                  Create a new account
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
                      placeholder="Name"
                      onChangeText={handleNameInput}
                    />
                  </Block>
                  <Block width={width * 0.8} style={{ marginBottom: 15 }}>
                    <TextInput
                      style={styles.input}
                      placeholder="Email"
                      onChangeText={handleEmailInput}
                    />
                  </Block>
                  <Block width={width * 0.8}>
                    <TextInput
                      style={styles.input}
                      placeholder="Password"
                      onChangeText={handlePasswordInput}
                    />
                    <Block row style={styles.passwordCheck}>
                      <Text size={12} color={argonTheme.COLORS.MUTED}>
                        password strength:
                      </Text>
                      {showPasswordStrength && (
                        <Text bold size={12} color={argonTheme.COLORS.SUCCESS}>
                          {passwordStrength}
                        </Text>
                      )}
                    </Block>
                  </Block>
                  <Block row width={width * 0.75}>
                    <Checkbox
                      checkboxStyle={{
                        borderWidth: 3,
                      }}
                      color={argonTheme.COLORS.PRIMARY}
                      label="I agree with the"
                    />
                    <Button
                      style={{ width: 100 }}
                      color="transparent"
                      textStyle={{
                        color: argonTheme.COLORS.PRIMARY,
                        fontSize: 14,
                      }}
                    >
                      Privacy Policy
                    </Button>
                  </Block>
                  <Block middle>
                    <Button color="primary" style={styles.createButton} onPress={handleRegisterBtnClick}>
                      <Text bold size={14} color={argonTheme.COLORS.WHITE}>
                        CREATE ACCOUNT
                      </Text>
                    </Button>
                    <Button
                      color="secondary"
                      style={styles.createButton}
                      onPress={() => navigation.navigate("App")}
                    >
                      <Text color={argonTheme.COLORS.BLACK}>Go to home</Text>
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
  registerContainer: {
    width: width * 0.9,
    height: height * 0.875,
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
  socialConnect: {
    backgroundColor: argonTheme.COLORS.WHITE,
    borderBottomWidth: StyleSheet.hairlineWidth,
    borderColor: "#8898AA",
  },
  socialButtons: {
    width: 120,
    height: 40,
    backgroundColor: "#fff",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    shadowRadius: 8,
    shadowOpacity: 0.1,
    elevation: 1,
  },
  socialTextButtons: {
    color: argonTheme.COLORS.PRIMARY,
    fontWeight: "800",
    fontSize: 14,
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

export default Register;
