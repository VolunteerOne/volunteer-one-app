import React from "react";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import { Block, theme } from "galio-framework";
const { width } = Dimensions.get("screen");

class Settings extends React.Component {

  render() {
    return (
      <Block flex center style={styles.home}>
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,
  },
});

export default Settings;
