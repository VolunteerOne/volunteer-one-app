import React from "react";
import { useState } from "react";
import { StyleSheet, Dimensions, ScrollView, TouchableOpacity } from "react-native";
import { Block, theme, Text } from "galio-framework";
const { width,height } = Dimensions.get("screen");

import { Images, argonTheme } from "../../constants";
import { TextInput } from "react-native-gesture-handler";


class Settings extends React.Component {
  render() {

    function handleNameSetting(input) {
      console.log(input);      
    }


    return (
      <ScrollView
      showsVerticalScrollIndicator={false}
      >
      <Block flex center style={styles.home} paddingTop={15} gap={15}>
        <Block style={styles.settingContainer}>
          <Block>
                <Text 
                  color="#000000" 
                  size={20}
                  style={{fontWeight: "bold"}}
                >
                  Name 
                </Text>
          </Block>
            <TextInput
              paddingLeft={50}
              color="#000000"
              placeholder="Name"
              onChangeText={handleNameSetting}
            />
        </Block>
        <Block style={styles.settingContainer}>
          <Block>
                <Text 
                  color="#000000" 
                  size={20}
                  style={{fontWeight: "bold"}}
                >
                  Email 
                </Text>
          </Block>
        </Block>
        <Block style={styles.settingContainer}>
          <Block>
                <Text 
                  color="#000000" 
                  size={20}
                  style={{fontWeight: "bold"}}
                >
                  Location 
                </Text>
          </Block>
        </Block>
        <Block style={styles.bioContainer}>
          <Block>
                <Text 
                  color="#000000" 
                  size={20}
                  style={{fontWeight: "bold"}}
                >
                  Bio 
                </Text>
          </Block>
        </Block>
      </Block>
    </ScrollView>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  settingContainer: {
    width: width * 0.8,
    height: height * 0.05,
    backgroundColor: "#FFFFFF",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    flex: 1,
    justifyContent: "center",
    paddingLeft: 5,
    shadowRadius: 8,
    shadowOpacity: 0.1,
    elevation: 1,
  },
  bioContainer: {
    width: width * 0.8,
    height: height * 0.3,
    backgroundColor: "#FFFFFF",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    flex: 1,
    paddingTop: 10,
    paddingLeft: 5,
    shadowRadius: 8,
    shadowOpacity: 0.1,
  },
});

export default Settings;
