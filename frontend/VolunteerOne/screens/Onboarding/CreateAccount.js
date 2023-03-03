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
  TouchableOpacity,
} from "react-native";
import { Block, Text } from "galio-framework";

import { Button } from "../../components";
import { Images, argonTheme } from "../../constants";


const { width, height } = Dimensions.get("screen");

/** ==================================== Create Account Screen ==================================== **/

const CreateAccount = ({ navigation }) => {
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
          {/* <Block style={styles.loginContainer}>
            <Block flex>
              <Block flex={0.5} middle style={styles.instructionText}>
                <Image source={logo} />
              </Block>
              <Block flex={0.17} middle style={styles.instructionText}>
                <TouchableOpacity
                  onPress={() => console.log("create account btn")}
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
                        onPress={() => console.log("forgot password btn clicked")}
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
          </Block> */}
        </Block>
      </ImageBackground>
    </Block>
  );
};


const styles = StyleSheet.create({
  instructionText: {
    flexDirection: "row",
  },
  createAccountContainer: {
    width: width * 0.9,
    height: height * 0.4,
    elevation: 1,
    overflow: "wrap",
  },
  optionButton: {
    width: width * 0.75,
    height: 60,
    marginTop: 25,
  },
});

export default CreateAccount;
