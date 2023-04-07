import React from "react";
import { StyleSheet, Dimensions, ScrollView, Text } from "react-native";
import { Block, theme } from "galio-framework";
const { width } = Dimensions.get("screen");
import NewAnnouncementModal from "../../components/Modals/NewAnnouncementModal";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import { Button } from "../../components";
import argonTheme from "../../constants/Theme";

class Feed extends React.Component {
  state = {
    modalVisible: false,
  };

  render() {
    const { modalVisible } = this.state;

    const handleModalVisible = () => {
      this.setState({ modalVisible: !modalVisible });
    };

    return (
      <Block flex center style={styles.home}>

      <Block middle>
          <Button
            color="primary"
            style={styles.button}
            onPress={() => handleModalVisible()}
          >
            <Block row middle>
              <MaterialCommunityIcons
                size={24}
                name="plus-box-outline"
                color={theme.COLORS.WHITE}
              />
              <Text bold size={14} style={styles.buttonTitle}>
                New Event
              </Text>
            </Block>
          </Button>
        </Block>
        {modalVisible && (
          <NewAnnouncementModal
            visible={this.state.modalVisible}
            setState={handleModalVisible}
          />
        )}

      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  button: {
    marginTop: theme.SIZES.BASE,
    marginBottom: 0,
    width: width * 0.9,
  },
  buttonTitle: {
    paddingLeft: 5,
    lineHeight: 19,
    fontWeight: "600",
    color: argonTheme.COLORS.WHITE,
  },
});

export default Feed;
