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
import { Block, Text } from "galio-framework";

import { Button } from "../../components";
import { Images, argonTheme } from "../../constants";


const { width, height } = Dimensions.get("screen");

/** ==================================== Create Account Screen ==================================== **/

const CreateAccount = ({ navigation }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [accountType, setAccountType] = useState("")
  console.log(accountType);

  return (
    <Block flex middle>
      <StatusBar hidden />
      <ImageBackground
        source={Images.RegisterBackground}
        style={{ width, height, zIndex: 1 }}
      >
        <Block safe flex middle>
          <Block style={styles.createAccountContainer}>
            <Block flex={1} middle style={styles.instructionText}>
              <Text
                color="#fff"
                size={38}
                style={{ fontWeight: "bold", padding: 6 }}
              >
                Are you a volunteer or organization?
              </Text>
            </Block>
            <Block flex={1} middle>
              <Button
                color="secondary"
                style={styles.optionButton}
                onPress={() => setAccountType('volunteer')}
              >
                <Text bold size={14} color={argonTheme.COLORS.BLACK}>
                  I'd like to volunteer
                </Text>
              </Button>
              <Button
                color="secondary"
                style={styles.optionButton}
                onPress={() => setAccountType('organization')}
              >
                <Text bold size={14} color={argonTheme.COLORS.BLACK}>
                  I'd like to recruit volunteers
                </Text>
              </Button>
            </Block>

          </Block>
        </Block>
      </ImageBackground>
    </Block>
  );
};
// }

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
