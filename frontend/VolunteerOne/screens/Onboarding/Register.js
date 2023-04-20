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
  TouchableOpacity,
} from "react-native";
import { Block, Checkbox, Text, theme } from "galio-framework";

import { Button, Icon, Input } from "../../components";
import { Images, argonTheme } from "../../constants";

const { width, height } = Dimensions.get("screen");

import logo from "../../assets/logo/logo2.png";

import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

/** ==================================== Register Screen ==================================== **/

const Register = ({ route, navigation }) => {

  const { userType } = route.params;
  console.log('userType passed: ', userType);

  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [seePassword, setSeePassword] = useState(false);

  const [passwordStrength, setPasswordStrength] = useState("weak");
  const [showPasswordStrength, setShowPasswordStrength] = useState(false);

  function handleNameInput(input) {
    setName(input);
  }

  // add email validation
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
        source={Images.Onboarding}
        style={{ width, height, zIndex: 1 }}
      >
        <Block safe flex middle>
          <Block style={styles.registerContainer}>
            <Block flex>
              <Block flex={0.5} middle>
                <Block center>
                  <Image source={Images.VolunteerOneIcon} style={styles.logo} />
                </Block>
              </Block>
              <Block flex={0.17} middle>
                <Text color="#8898AA" size={14}>
                  Sign up with credentials
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
                  <Block width={width * 0.8} style={styles.inputPassword}>
                    <Block row>
                      <TextInput
                        flex={0.9}
                        secureTextEntry={seePassword}
                        // style={styles.input}
                        placeholder="Password"
                        onChangeText={handlePasswordInput}
                      />
                      <Block middle center style={styles.visibleIcon}>
                        <TouchableOpacity
                          flex={0.2}
                          onPress={() =>
                            setSeePassword((seePassword) => !seePassword)
                          }
                        >
                          {seePassword ? (
                            <MaterialCommunityIcons
                              size={24}
                              name="eye"
                              color={theme.COLORS.ICON}
                            />
                          ) : (
                            <MaterialCommunityIcons
                              size={24}
                              name="eye-off"
                              color={theme.COLORS.ICON}
                            />
                          )}
                        </TouchableOpacity>
                      </Block>
                    </Block>
                  </Block>
                  <Block middle>
                    <Button
                      color="primary"
                      style={styles.createButton}
                      onPress={handleRegisterBtnClick}
                    >
                      <Text bold size={14} color={argonTheme.COLORS.WHITE}>
                        CREATE ACCOUNT
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
  inputPassword: {
    borderColor: argonTheme.COLORS.BORDER,
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
    paddingTop:10,
    marginBottom:15
  },
  registerContainer: {
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
  passwordCheck: {
    paddingLeft: 10,
    paddingTop: 13,
    paddingBottom: 30,
  },
  createButton: {
    width: width * 0.5,
    marginTop: 25,
  },
  visibleIcon: {
    paddingLeft: 90,
  },
  logo: {
    width: 265,
    height: 50,
    zIndex: 2,
    position: 'relative',
    marginTop: '20%'
  },
});

export default Register;
